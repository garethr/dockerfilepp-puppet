package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Args struct {
	Value string
}

func replace(args string, temp string) string {
	argument := Args{args}
	tmpl, err := template.New("replacer").Parse(temp)
	if err != nil {
		fmt.Println("A problem occured parsing one of the processors:")
		fmt.Println(err)
		fmt.Print(temp)
		os.Exit(2)
	}
	buff := bytes.NewBufferString("")
	err = tmpl.Execute(buff, argument)
	if err != nil {
		fmt.Println("A problem occured executing one of the processors:")
		fmt.Println(err)
		fmt.Print(temp)
		os.Exit(2)
	}
	return buff.String()
}

func main() {
	stat, _ := os.Stdin.Stat()
	// We detect whether we have anything on stdin to process
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var buffer bytes.Buffer
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			buffer.WriteString(scanner.Text() + "\n")
		}
		value := buffer.String()

		replacements := make(map[string]string)
		files := AssetNames()
		for _, file := range files {
			content, _ := Asset(file)
			replacements[file] = string(content)
		}

		for regex, tmpl := range replacements {
			re := regexp.MustCompile(regex + "(.*)")
			matches := re.FindStringSubmatch(value)
			if len(matches) == 2 {
				args := matches[1]
				args = strings.TrimLeft(args, " ")
				value = re.ReplaceAllString(value, replace(args, tmpl))
			}
		}
		fmt.Print(value)
	} else {
		fmt.Println(`dockerfilepp is a tool for adding new instructions to Dockerfile

Usage:

    dockerfilepp-puppet < Dockerfile
    Dockerfile | dockerfilepp-puppet

dockerfilepp takes a Dockerfile on stdin and outputs to stdout a modified version
of the same Dockerfile which has been through the registered pre-processors.`)
	}
}
