package document

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/models"
	"github.com/backent/fra-golang/models/elastic"
	webElastic "github.com/backent/fra-golang/web/elastic"
	"github.com/elastic/go-elasticsearch/v8"
)

type RepositoryDocumentSearchEsImpl struct{}

func NewRepositoryDocumentSearchEsImpl() RepositoryDocumentSearchInterface {
	return &RepositoryDocumentSearchEsImpl{}
}

func (implementation *RepositoryDocumentSearchEsImpl) SearchByProductName(client *elasticsearch.Client, name string) ([]elastic.DocumentSearchGlobal, error) {
	query := elastic.GenerateQuery("risk")
	res, err := client.Search(
		client.Search.WithIndex(elastic.IndexNameDocumentSearchGlobal),
		client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		return nil, err
	}

	var responseObj webElastic.Response

	helpers.DecodeRequestElastic(res, &responseObj)
	var actualDoc []elastic.DocumentSearchGlobal
	for _, val := range responseObj.HitsData.Hits {
		doc, err := toReal(val.Source)
		if err != nil {
			return nil, err
		}
		actualDoc = append(actualDoc, doc)
	}

	return actualDoc, nil

}

func toReal(data map[string]interface{}) (elastic.DocumentSearchGlobal, error) {
	var result elastic.DocumentSearchGlobal

	// Marshal the interface{} to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	// Unmarshal the JSON into a struct
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (implementation *RepositoryDocumentSearchEsImpl) IndexProduct(client *elasticsearch.Client, document models.Document) error {

	documentIndex := elastic.ModelDocumentToIndexDocumentSearchGlobal(document)
	data, _ := json.Marshal(documentIndex)
	res, err := client.Index(elastic.IndexNameDocumentSearchGlobal, bytes.NewReader(data), client.Index.WithDocumentID(documentIndex.Uuid))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusCreated {
		return errors.New("error while indexing document with status code :" + strconv.Itoa(res.StatusCode))
	}

	return nil
}
