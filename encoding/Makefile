export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test: clean setup ## Run tests
	go run gjson/gjson.go
	go generate ./model && go run easyjson/easyjson.go
	go run jsonreader/jsonreader.go

setup: ## Install & Setup
	go get -u github.com/mailru/easyjson/...

clean: ## Remove artifacts
	@rm -f model/*_easyjson.go

.PHONY: clean setup test
