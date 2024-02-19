package main

import (
	"context"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/models"
	"github.com/backent/fra-golang/models/elastic"
	"github.com/backent/fra-golang/repositories/document"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

const INDEX_NAME = elastic.IndexNameDocumentSearchGlobal

func main() {

	err := godotenv.Load(".env")
	helpers.PanicIfError(err)

	sql := libs.NewDatabase()
	tx, err := sql.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	ctx := context.Background()

	repositoryDocument := document.NewRepositoryDocumentImpl()
	documents, err := repositoryDocument.FindAllWithDetail(ctx, tx)
	helpers.PanicIfError(err)

	client := libs.NewElastic()

	if helpers.CheckIndexExists(client, INDEX_NAME) {
		err := helpers.DeleteIndex(client, INDEX_NAME)
		helpers.PanicIfError(err)
	}

	configuration := elastic.IndexNameDocumentSearchGlobalSettings

	err = helpers.CreateIndex(client, INDEX_NAME, configuration)
	helpers.PanicIfError(err)

	reIndexing(client, documents)
}

func reIndexing(client *elasticsearch.Client, documents []models.Document) {
	repositoryDocumentSearch := document.NewRepositoryDocumentSearchEsImpl()

	for _, document := range documents {
		err := repositoryDocumentSearch.IndexProduct(client, document)
		helpers.PanicIfError(err)
	}
}
