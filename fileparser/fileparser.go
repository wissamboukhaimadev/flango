package fileparser

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"flango/evaluator"
	"flango/lexer"
	"flango/object"
	"flango/parser"
)

const MONKEY_FACE = `
 _____ _     ____  _      _____ ____ 
/    // \   /  _ \/ \  /|/  __//  _ \
|  __\| |   | / \|| |\ ||| |  _| / \|
| |   | |_/\| |-||| | \||| |_//| \_/|
\_/   \____/\_/ \|\_/  \|\____\\____/

`

func ParseFile(filePath string) {
	fileext := filepath.Ext(filePath)
	if fileext == ".fl" {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("error reading the file")
			return
		}
		// fmt.Println(string(content))
		evaluateFile(string(content), os.Stdout)

	} else {
		fmt.Println("file extension should be .fl")
	}
}

func evaluateFile(fileContent string, out io.Writer) {
	env := object.NewEnvironment()
	l := lexer.New(fileContent)
	p := parser.New(l)
	programe := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
	}
	evaluated := evaluator.Eval(programe, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")

	}
}

func printParserErrors(errors []string) {
	color.Cyan(MONKEY_FACE)
	color.Red("parser error :\n")
	for _, msg := range errors {
		color.Red("\t" + msg + "\n")
	}
	os.Exit(1)
}
