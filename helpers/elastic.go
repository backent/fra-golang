package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func DecodeRequestElastic(r *esapi.Response, requestVar interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(requestVar)
	PanicIfError(err)
}

func CreateIndex(client *elasticsearch.Client, indexName string, configuration string) error {
	// Create a strings reader for the index settings JSON
	settingsReader := strings.NewReader(configuration)
	res, err := client.Indices.Create(indexName, client.Indices.Create.WithBody(settingsReader))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusCreated {
		var resErr interface{}
		DecodeRequestElastic(res, &resErr)
		log.Fatal(resErr)
		return errors.New("error while create index with status code :" + strconv.Itoa(res.StatusCode))
	}

	return nil
}

func DeleteIndex(client *elasticsearch.Client, indexName string) error {
	res, err := client.Indices.Delete([]string{indexName})
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusCreated {
		return errors.New("error while delete index with status code :" + strconv.Itoa(res.StatusCode))
	}

	return nil
}

func CheckIndexExists(client *elasticsearch.Client, indexName string) bool {
	res, err := client.Indices.Exists([]string{indexName})
	PanicIfError(err)
	defer res.Body.Close()

	return res.StatusCode == 200
}
