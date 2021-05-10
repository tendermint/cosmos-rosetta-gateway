#!/usr/bin/make -f

build:
	go build -mod=readonly ./...

test:
	go test -mod=readonly -race github.com/tendermint/cosmos-rosetta-gateway/...

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*/generated/*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*/generated/*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*/generated/*" | xargs goimports -w -local github.com/tendermint/cosmos-rosetta-gateway

clean:
	rm -f crg coverage.txt

.PHONY: format test clean dev gen-all gen-mocks gen-clients
