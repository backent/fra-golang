package document

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/models/elastic"
	webElastic "github.com/backent/fra-golang/web/elastic"
	"github.com/elastic/go-elasticsearch/v8"
)

type RepositoryDocumentSearchEsImpl struct{}

func NewRepositoryDocumentSearchEsImpl() RepositoryDocumentSearchInterface {
	return &RepositoryDocumentSearchEsImpl{}
}

func (implementation *RepositoryDocumentSearchEsImpl) SearchByProductName(client *elasticsearch.Client, name string) {
	query := elastic.GenerateQuery("risk")
	res, err := client.Search(
		client.Search.WithIndex(elastic.IndexNameDocumentSearchGlobal),
		client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		panic(err)
	}

	var responseObj webElastic.Response

	helpers.DecodeRequestElastic(res, &responseObj)
	var actualDoc []elastic.DocumentSearchGlobal
	for _, val := range responseObj.HitsData.Hits {
		doc, err := toReal(val.Source)
		if err != nil {
			panic(err)
		}
		actualDoc = append(actualDoc, doc)
	}

	for _, val := range actualDoc {
		fmt.Println(val.ProductName)
	}

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
