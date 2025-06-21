DATA_DIR = ./data


.PHONY: clean
clean:
	@rm -rf $(DATA_DIR)

.PHONY: test
test:
	@go test ./...
