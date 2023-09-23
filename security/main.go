package main

import (
	"flag"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	log.SetFlags(0)

	var (
		err error

		// --> Configure the path to the certificate authority and the password
		//
		cacert   = flag.String("cacert", "certificates/ca/ca.crt", "Path to the file with certificate authority")
		password = flag.String("password", "elastic", "Elasticsearch password")
	)
	flag.Parse()

	// --> Read the certificate from file
	//
	cert, err := os.ReadFile(*cacert)
	if err != nil {
		log.Fatalf("ERROR: Unable to read CA from %q: %s", *cacert, err)
	}

	es, err := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{"https://es01:9200"},
			Username:  "elastic",
			Password:  *password,

			// --> Pass the certificate to the client
			//
			CACert: cert,
		})
	if err != nil {
		log.Fatalf("ERROR: Unable to create client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("ERROR: Unable to get response: %s", err)
	}

	log.Println(res)
}
