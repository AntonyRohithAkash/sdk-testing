#!/usr/bin/sh

go test -v -coverpkg=./sdk -coverprofile=coverage.out ./sdk

go tool cover -func coverage.out