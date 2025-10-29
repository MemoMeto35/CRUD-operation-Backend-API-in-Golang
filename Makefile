build:
	@go build -o bin/GOPROJECT cmd/main.go
test:
	@go test -v ./...
run: build
	@bin/GOPROJECT