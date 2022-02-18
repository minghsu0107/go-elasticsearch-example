package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	wg        sync.WaitGroup
	testIndex = "testindex"
)

func main() {
	config := elasticsearch.Config{
		Addresses:           []string{"http://localhost:9200"},
		MaxRetries:          3,
		CompressRequestBody: false,
		EnableDebugLogger:   true,
		// Enable node discovery only when the client is connected to the cluster directly
		// not when the cluster is behind a proxy, which is also the case when using Elasticsearch Service
		DiscoverNodesOnStart: false,
	}
	es, _ := elasticsearch.NewClient(config)

	// Get cluster info
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])

	// Index documents concurrently
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      testIndex,
				DocumentID: strconv.Itoa(i + 1),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true", // refresh index
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	wg.Wait()

	// Update
	var b strings.Builder
	b.WriteString(`{"doc" : {"title" : "test: updated"}}`)
	req1 := esapi.UpdateRequest{
		Index:           testIndex,
		DocumentID:      "1",
		Body:            strings.NewReader(b.String()),
		Refresh:         "true",
		RetryOnConflict: esapi.IntPtr(3),
		Source:          []string{"title"},
	}
	res1, err := req1.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res1.Body.Close()
	if res1.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res1.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res1.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r1 map[string]interface{}
	if err := json.NewDecoder(res1.Body).Decode(&r1); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Println(r1)

	// Search for the indexed documents
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res2, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(testIndex),
		es.Search.WithBody(&buf),
		// indicate if the number of documents that match the query should be tracked.
		// a number can also be specified, to accurately track the total hit count up to the number
		es.Search.WithTrackTotalHits(true),
		// starting offset
		es.Search.WithFrom(0),
		// number of hits to return (default: 10)
		es.Search.WithSize(50),
		// a list of <field>:<direction> pairs
		// just for demo; sort by _id is not recommended
		es.Search.WithSort("_id:asc"), // direction could be asc or dsc
		es.Search.WithSource("title"),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res2.Body.Close()

	if res2.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res2.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res2.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r2 map[string]interface{}
	if err := json.NewDecoder(res2.Body).Decode(&r2); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res2.Status(),
		int(r2["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r2["took"].(float64)),
	)

	// Print the ID and document source for each hit.
	for _, hit := range r2["hits"].(map[string]interface{})["hits"].([]interface{}) {
		h := hit.(map[string]interface{})
		log.Printf(" * ID=%s, title=%s", h["_id"], h["_source"].(map[string]interface{})["title"].(string))
	}

	// Delete
	req3 := esapi.DeleteRequest{
		Index:      testIndex,
		DocumentID: "2",
		Refresh:    "true",
	}
	res3, err := req3.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res3.Body.Close()
	if res3.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res3.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res3.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	var r3 map[string]interface{}
	if err := json.NewDecoder(res3.Body).Decode(&r3); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("result: %s", r3["result"].(string))
}
