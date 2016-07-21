GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
OUTPUT ?= bin/darwin/amd64/dockerfilepp-puppet

all: build

tools:
	go get -u github.com/golang/lint/golint
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/kisielk/errcheck

deps:
	glide up

check:
	errcheck

bindata:
	go-bindata -prefix "processors/" -o processors.go  processors

build: bindata
	go build -v -o ${OUTPUT}

linux: bindata
	env GOOS=linux GOAARCH=amd64 go build -v -o bin/linux/amd64/dockerfilepp-puppet

example: build
	cat example/Dockerfile | ./${OUTPUT}

diff: build
	cat example/Dockerfile | ./${OUTPUT} > Dockerfile.result
	-colordiff -y example/Dockerfile Dockerfile.result
	rm Dockerfile.result

lint:
	golint main.go

test:
	go test

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

fmt:
	gofmt -w $(GOFMT_FILES)
