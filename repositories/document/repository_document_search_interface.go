package document

import "github.com/elastic/go-elasticsearch/v8"

type RepositoryDocumentSearchInterface interface {
	SearchByProductName(client *elasticsearch.Client, name string)
}
