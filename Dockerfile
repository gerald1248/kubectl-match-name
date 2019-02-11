FROM golang:1.11.1
WORKDIR /go/src/github.com/gerald1248/kubectl-match-name/
COPY * ./
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on
RUN \
  go mod download && \
  go get && \
  go vet && \
  go build -o kubectl-match_name .
