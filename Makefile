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

bindata: tools
	go-bindata -prefix "processors/" -o processors.go  processors

dirs:
	mkdir -p releases
	mkdir -p bin/linux/amd64
	mkdir -p bin/windows/amd64
	mkdir -p bin/darwin/amd64

build_deps: bindata deps dirs

build: darwin linux windows

darwin: build_deps
	pwd
	go build -v -o ${OUTPUT}
	ls -alh
	ls bin
	ls bin/linux
	ls bin/linux/amd64
	tar -cvzf releases/dockerfilepp-puppet-darwin-amd64.tar.gz bin/linux/amd64/dockerfilepp-puppet

linux: build_deps
	env GOOS=linux GOAARCH=amd64 go build -v -o bin/linux/amd64/dockerfilepp-puppet
	tar -cvzf releases/dockerfilepp-puppet-linux-amd64.tar.gz bin/linux/amd64/dockerfilepp-puppet

windows: build_deps
	env GOOS=windows GOAARCH=amd64 go build -v -o bin/windows/amd64/dockerfilepp-puppet
	tar -cvzf releases/dockerfilepp-puppet-windows-amd64.tar.gz bin/windows/amd64/dockerfilepp-puppet

example: darwin
	cat example/Dockerfile | ./${OUTPUT}

diff: darwin
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
