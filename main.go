// dockerfilepp-puppet is a tool which adds new instructions to Dockerfile for
// using Puppet as part of an image build process.
//
// It does that by post-processing a standard Dockerfile and replacing new
// instructions with valid snippets of Dockerfile.
//
// Under the hood dockerfilepp-puppet uses the dockerfilepp library.
//
// Usage:
//
//    dockerfilepp-puppet < Dockerfile
//
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
	dockerfilepp.Process(replacements, `dockerfilepp-puppet is a tool for adding new instructions to Dockerfile

Usage:

    dockerfilepp-puppet < Dockerfile
    Dockerfile | dockerfilepp-puppet

dockerfilepp-puppet takes a Dockerfile on stdin and outputs to stdout a modified version
of the same Dockerfile which has been through the registered pre-processors.`)
}
