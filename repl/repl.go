package repl

import (
	"bufio"
	"fmt"
	"io"

	"bangu/lexer"
	"bangu/parser"

	"bangu/evaluator"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return // EOF or error
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const BANGU_FACE = `
    .-""""""-.
   /          \
  |  .    .    |
  |            |
  |     O      |
   \           /
    '-.......-'
       |  |
   ____||||____
  (____________)
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, BANGU_FACE)
	io.WriteString(out, "Woops! We ran into some bangu business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
