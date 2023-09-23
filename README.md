# Go Elasticsearch Example
Examples to show how to use [go-elasticsearch](https://github.com/elastic/go-elasticsearch), a Golang Elasticsearch client SDK.

Reference: https://github.com/elastic/go-elasticsearch
## Client Configuration
```go
tr := http.DefaultTransport.(*http.Transport).Clone()
tr.MaxIdleConns = 100
tr.MaxConnsPerHost = 100
tr.MaxIdleConnsPerHost = 100

cfg := elasticsearch.Config{
	Addresses: []string{"http://localhost:9200"},
	Transport: &tr,
}

es, err := elasticsearch.NewClient(cfg)
```
