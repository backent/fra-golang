package libs

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"github.com/backent/fra-golang/helpers"
	"github.com/elastic/go-elasticsearch/v8"
)

func NewElastic() *elasticsearch.Client {

	// Load CA certificate
	caCert, err := os.ReadFile(os.Getenv("ELASTIC_CRT_FILE"))
	if err != nil {
		log.Printf("Error loading CA certificate: %s", err)
		client, err := elasticsearch.NewDefaultClient()
		helpers.PanicIfError(err)
		return client
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	address := []string{"https://localhost:9200"}
	username := os.Getenv("ELASTIC_USERNAME")
	password := os.Getenv("ELASTIC_PASSWORD")
	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: address, Username: username, Password: password, Transport: &http.Transport{TLSClientConfig: tlsConfig}})
	helpers.PanicIfError(err)

	return client
}
