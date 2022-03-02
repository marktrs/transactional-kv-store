#!/bin/bash

unit_coverage_test() {
  go mod download
  go test $(go list ./...| grep -v /internal/) -race -covermode atomic -coverprofile=coverage.out ./...
}

unit_coverage_test
