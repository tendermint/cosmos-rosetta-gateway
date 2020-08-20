gen-all: gen-clients gen-mocks

gen-clients:
	scripts/gen-clients

gen-mocks:
	scripts/gen-mocks
format:
	go fmt ./...
test:
	go test ./...
dev:
	scripts/dev
