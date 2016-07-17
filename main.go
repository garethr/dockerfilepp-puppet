package main

import (
	"github.com/garethr/dockerfilepp"
)

func main() {
	replacements := make(map[string]string)
	files := AssetNames()
	for _, file := range files {
		content, _ := Asset(file)
		replacements[file] = string(content)
	}
	dockerfilepp.Process(replacements, `dockerfilepp is a tool for adding new instructions to Dockerfile

Usage:

    dockerfilepp-puppet < Dockerfile
    Dockerfile | dockerfilepp-puppet

dockerfilepp takes a Dockerfile on stdin and outputs to stdout a modified version
of the same Dockerfile which has been through the registered pre-processors.`)
}
