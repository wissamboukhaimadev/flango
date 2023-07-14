package main

import (
	"os"

	"flango/fileparser"
	"flango/repl"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		filePath := args[1]
		fileparser.ParseFile(filePath)
	}
}
