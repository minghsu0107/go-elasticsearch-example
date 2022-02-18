export ELASTICSEARCH_URL=http://elastic:elastic@localhost:9200

test: test-default test-indexer

test-default:
	go run default/default.go

test-indexer:
	go run indexer/indexer.go

.PHONY: test test-default test-indexer
