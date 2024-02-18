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
		log.Fatalf("Error loading CA certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	address := []string{"https://localhost:9200"}
	username := "elastic"
	password := "adminlocal123"
	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: address, Username: username, Password: password, Transport: &http.Transport{TLSClientConfig: tlsConfig}})
	helpers.PanicIfError(err)

	return client
}
