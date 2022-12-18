lint:
	golangci-lint run

test:
	go test -race ./...

cover:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
