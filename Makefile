OUTPUT ?= dockerfilepp-puppet

all: build

bindata:
	go-bindata -prefix "processors/" -o processors.go  processors

build: bindata
	go build -o ${OUTPUT}

test: build
	cat Dockerfile | ./${OUTPUT}

diff: build
	cat Dockerfile | ./${OUTPUT} > Dockerfile.result
	-colordiff -y Dockerfile Dockerfile.result
	rm Dockerfile.result
