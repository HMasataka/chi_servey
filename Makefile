chi:
	oapi-codegen -generate "types" -package "domain" openapi.yml > types.gen.go
	oapi-codegen -generate "chi-server" -package "domain" openapi.yml > server.gen.go

clean:
	go clean

test:
	go test -v ./...

lint:
	golangci-lint -v run

.PHONY: chi
