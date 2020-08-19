gen-all: gen-clients gen-mocks

gen-clients:
	scripts/gen-clients

gen-mocks:
	scripts/gen-mocks

dev:
	scripts/dev

test:
	go test ./...