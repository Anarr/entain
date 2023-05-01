#!/bin/sh
echo "Run tests..."
CGO_ENABLED=0  go test -mod=vendor -v ./...
echo "Finish tests"
