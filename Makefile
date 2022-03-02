.PHONY: build clean deploy

build:
	go build -o bin/main cmd/main.go

clean:
	rm -rf ./bin .DS_Store

run:
	go run cmd/main.go

test:
	go test -race ./...

docker-build:
	- docker build -t ${APP_NAME} .

docker-run: docker-build
	- docker run -it ${APP_NAME} go run .

docker-test: docker-build
	- docker run -it ${APP_NAME} go test -race -v ./...

coverage: ## Creates unit tests coverage files
	@./.github/scripts/gocoverage.sh

unit-test: ## Runs unit tests
	@test -z "$$(golangci-lint run ./...)"

lint-test: ## Runs lint
	@golangci-lint -v run ./...
	@test -z "$$(golangci-lint run ./...)"