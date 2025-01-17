.phony: test check lint

check: lint
	@go vet ./...
	@gofmt -l .
	@test -z "$$( gofmt -l . )"

lint:
ifndef SKIP_LINT
	@golangci-lint run ./...
endif

test:
	go test -race -v ./...

cov:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...
