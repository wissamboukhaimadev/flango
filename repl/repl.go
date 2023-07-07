package repl

import (
	"bufio"
	"fmt"
	"io"
	"flango/evaluator"
	"flango/lexer"
	"flango/object"
	"flango/parser"
  "github.com/fatih/color"
)

const MONKEY_FACE = `
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

  color.Cyan(MONKEY_FACE)

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
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
  color.Cyan(MONKEY_FACE)
  color.Red("parser error :\n")
	for _, msg := range errors {
    color.Red("\t"+msg+"\n")
	}
}
