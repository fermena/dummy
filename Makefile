.PHONY: test
test: build
	docker compose run app go test ./... -cover -coverprofile=coverage.out

.PHONY: build
build:
	 docker compose build --no-cache
