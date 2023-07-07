package main

import (
	 "github.com/fatih/color"
	"flango/repl"
	"os"
)

func main() {
  red := color.New(color.FgWhite)

  boldRed := red.Add(color.Bold)
  boldRed.Println("Hello in FLANGO")

	repl.Start(os.Stdin, os.Stdout)
}
