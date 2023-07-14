package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/fatih/color"

	"flango/evaluator"
	"flango/lexer"
	"flango/object"
	"flango/parser"
)

const FLANGO = `
 _____ _     ____  _      _____ ____ 
/    // \   /  _ \/ \  /|/  __//  _ \
|  __\| |   | / \|| |\ ||| |  _| / \|
| |   | |_/\| |-||| | \||| |_//| \_/|
\_/   \____/\_/ \|\_/  \|\____\\____/

`

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	red := color.New(color.FgWhite)
	boldRed := red.Add(color.Bold)
	boldRed.Println("Hello in FLANGO")

	color.Cyan(FLANGO)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(errors []string) {
	color.Cyan(FLANGO)
	color.Red("parser error :\n")
	for _, msg := range errors {
		color.Red("\t" + msg + "\n")
	}
}
