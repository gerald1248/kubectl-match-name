#!/bin/sh

GO111MODULE="on"

go mod download
go get
go vet
go test -v -cover
go build -o kubectl-match_name .
