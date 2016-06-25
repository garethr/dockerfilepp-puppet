package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

type Args struct {
	Value string
}

func replace(args string, temp string) string {
	sweaters := Args{args}
	tmpl, _ := template.New("replacer").Parse(temp)
	buff := bytes.NewBufferString("")
	_ = tmpl.Execute(buff, sweaters)
	return buff.String()
}

func main() {
	var buffer bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer.WriteString(scanner.Text() + "\n")
	}
	value := buffer.String()

	replacements := make(map[string]string)
	files, _ := ioutil.ReadDir("pp")
	for _, file := range files {
		content, _ := ioutil.ReadFile("pp/" + file.Name())
		replacements[file.Name()] = string(content)
	}

	for regex, tmpl := range replacements {
		re := regexp.MustCompile(regex + "(.*)")
		args := re.FindStringSubmatch(value)[1]
		args = strings.TrimLeft(args, " ")
		value = re.ReplaceAllString(value, replace(args, tmpl))
	}

	fmt.Println(value)
}
