package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	fname  string
	token  string
	stderr = log.New(os.Stderr, "", 0)
)

func main() {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if !errors.Is(err, nil) {
		stderr.Fatal(err)
	}

	replacer := strings.NewReplacer(token, string(stdin))
	file, err := ioutil.ReadFile(fname)
	if !errors.Is(err, nil) {
		stderr.Fatal(err)
	}

	replacer.WriteString(os.Stdout, string(file))
}
