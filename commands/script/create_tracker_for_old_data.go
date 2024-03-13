package main

import (
	"context"
	"fmt"
	"log"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/models"
	"github.com/backent/fra-golang/repositories/document_tracker"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	helpers.PanicIfError(err)

	documentTrackerRepository := document_tracker.NewRepositoryDocumentTrackerImpl()
	ctx := context.Background()
	sql := libs.NewDatabase()
	tx, err := sql.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	query := fmt.Sprintf(`
		WITH main_table AS (
			SELECT * FROM %s 
		), group_by_uuid AS (
			SELECT d1.*
			FROM main_table d1
			JOIN (
					SELECT uuid, MAX(id) AS max_id
					FROM main_table
					GROUP BY uuid
			) d2 ON d1.uuid = d2.uuid AND d1.id = d2.max_id
		), main_table_after_grouped AS (
			SELECT * FROM group_by_uuid
		)
		SELECT 
		uuid,
		created_at
	FROM main_table_after_grouped`, models.DocumentTable)

	rows, err := tx.QueryContext(ctx, query)
	helpers.PanicIfError(err)
	defer rows.Close()

	var documents []models.Document

	for rows.Next() {
		var document models.Document
		err := rows.Scan(&document.Uuid, &document.CreatedAt)
		helpers.PanicIfError(err)

		documents = append(documents, document)
	}

	fmt.Println("document to process: ", len(documents))
	var documentSkip int
	var documentSuccess int
	var documentFail int

	for _, doc := range documents {
		_, err = documentTrackerRepository.GetByUUId(ctx, tx, doc.Uuid)

		if err != nil && err.Error() == "document tracker not found" {
			errCreate := documentTrackerRepository.Create(ctx, tx, doc.Uuid, doc.CreatedAt)
			if errCreate != nil {
				documentFail++
				log.Println("errCreate : ", errCreate)
			}
			documentSuccess++
		} else if err != nil {
			documentFail++
			log.Println("err : ", err)
		} else {
			documentSkip++
		}
	}

	fmt.Println("document skip :", documentSkip)
	fmt.Println("document success :", documentSuccess)
	fmt.Println("document fail :", documentFail)

}
