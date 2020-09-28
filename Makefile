#!/usr/bin/make -f

crg: go.sum
	go build -mod=readonly ./cmd/crg/

gen-all: gen-mocks

gen-mocks:
	scripts/gen-mocks

install:
	go install ./cmd/crg

test:
	go test -mod=readonly -race github.com/tendermint/cosmos-rosetta-gateway/...

dev:
	go run cmd/crg/*.go --blockchain "Test" --network "Test"

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*/generated/*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*/generated/*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "*/generated/*" | xargs goimports -w -local github.com/tendermint/cosmos-rosetta-gateway

clean:
	rm -f crg coverage.txt

.PHONY: format test clean dev gen-all gen-mocks gen-clients
