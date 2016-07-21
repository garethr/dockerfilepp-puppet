OUTPUT ?= dockerfilepp-puppet
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

all: build

tools:
	go get -u github.com/golang/lint/golint
	go get -u github.com/jteeuwen/go-bindata/...

deps:
	glide up

bindata:
	go-bindata -prefix "processors/" -o processors.go  processors

build: bindata
	go build -o ${OUTPUT}

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
