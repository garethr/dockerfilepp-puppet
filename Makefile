OUTPUT ?= dockerfilepp-puppet

all: build

bindata:
	go-bindata -prefix "processors/" -o processors.go  processors

build: bindata
	go build -o ${OUTPUT}

test: build
	cat example/Dockerfile | ./${OUTPUT}

diff: build
	cat example/Dockerfile | ./${OUTPUT} > Dockerfile.result
	-colordiff -y example/Dockerfile Dockerfile.result
	rm Dockerfile.result
