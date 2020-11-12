package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const mandoc = `
NAME
  %[1]s - replaces TOKEN in TEMPLATE with STDIN

SYNOPSIS
  %[1]s -token TOKEN -template ""

DESCRIPTION
  It is sometimes desirable to combine two files by inserting one into the middle of another.

OPTIONS
%[2]s
EXAMPLE
  $ cat giant.txt | %[1]s -template small.tmpl > combined.txt

VERSION
  version:  %[3]s
  compiled: %[4]s
  built:    %[5]s

`

func init() {
	flag.Usage = func() {
		var def bytes.Buffer
		flag.CommandLine.SetOutput(&def)
		flag.PrintDefaults()

		fmt.Printf(
			mandoc,
			filepath.Base(os.Args[0]),
			def.String(),
			buildVersion,
			compiledBy,
			buildTimestamp,
		)
	}

	flag.StringVar(&fname, "template", "", "the name of the template file")
	flag.StringVar(&token, "token", "TOKEN", "the string being replaced")
	flag.Parse()

	if len(fname) < 1 {
		stderr.Fatalln("-template is required; -help for more info")
	}

	if len(token) < 1 {
		stderr.Fatalln("-token is required; -help for more info")
	}
}

func underline(s string) string {
	return "\u001b[;4m" + s + "\u001b[;0m"
}
