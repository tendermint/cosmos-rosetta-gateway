gen-all: gen-clients gen-mocks

gen-clients:
	scripts/gen-clients

gen-mocks:
	scripts/gen-mocks
test:
	go test ./...