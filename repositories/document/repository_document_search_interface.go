package document

import (
	"github.com/backent/fra-golang/models"
	"github.com/backent/fra-golang/models/elastic"
	"github.com/elastic/go-elasticsearch/v8"
)

type RepositoryDocumentSearchInterface interface {
	SearchByProductName(client *elasticsearch.Client, name string) ([]elastic.DocumentSearchGlobal, error)
	IndexProduct(client *elasticsearch.Client, document models.Document) error
}
