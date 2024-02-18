package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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
	for _, document := range documents {
		err := esIndex(client, document)
		helpers.PanicIfError(err)
	}
}

func esIndex(client *elasticsearch.Client, document models.Document) error {
	documentIndex := elastic.ModelDocumentToIndexDocumentSearchGlobal(document)
	data, _ := json.Marshal(documentIndex)
	res, err := client.Index(INDEX_NAME, bytes.NewReader(data), client.Index.WithDocumentID(documentIndex.Uuid))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted && res.StatusCode != http.StatusCreated {
		return errors.New("error while indexing document with status code :" + strconv.Itoa(res.StatusCode))
	}
	return nil
}
