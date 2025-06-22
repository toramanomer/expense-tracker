DATA_DIR 	= 	./data
GOOS 		?= 	$(shell go env GOOS)
GOARCH 		?= 	$(shell go env GOARCH)

.PHONY: clean
clean:
	@rm -rf $(DATA_DIR)

.PHONY: test
test:
	@go test ./...

.PHONY download
download:
	@go mod download

.PHONY: build
build: download
	@go build -o expense-tracker main.go
