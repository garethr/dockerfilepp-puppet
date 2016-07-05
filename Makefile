all: build

bindata:
	go-bindata -prefix "processors/" -o processors.go  processors

build: bindata
	go build

test: build
	cat Dockerfile | ./dockerfilepp

diff: build
	cat Dockerfile | ./dockerfilepp > Dockerfile.result
	-colordiff -y Dockerfile Dockerfile.result
	rm Dockerfile.result
