package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in) // creates a new scanner with the input

	// Continue searching until a token is not scanned
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan() // produces a boolean for whether a line was scanned or not
		if !scanned {
			return
		}

		line := scanner.Text() // the string produced from the scan
		l := lexer.New(line)   // create a lexer for the line
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
