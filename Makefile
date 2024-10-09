VERSION ?= "latest"

MIN_COVERAGE = 95

.PHONY: lint
lint:
	go fmt ./...
	golangci-lint run
	go vet ./...
	go mod download && go mod tidy && go mod verify

@PHONY: tag-next
tag-next:
	@./cmd/tag-next.sh
