.PHONY: test
test:
	go test ./... -cover -coverprofile=coverage.out
