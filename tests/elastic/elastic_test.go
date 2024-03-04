package elastic

import (
	"fmt"
	"testing"
	"time"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/libs"
	"github.com/backent/fra-golang/repositories/document"
	"github.com/joho/godotenv"
)

func TestElastic(t *testing.T) {
	start := time.Now()
	err := godotenv.Load("../../.env")
	helpers.PanicIfError(err)
	client := libs.NewElastic()
	repository := document.NewRepositoryDocumentSearchEsImpl()
	repository.SearchByProductName(client, "test", 10, 0)

	elapsed := time.Since(start)
	fmt.Println("Time take :", elapsed)
}
