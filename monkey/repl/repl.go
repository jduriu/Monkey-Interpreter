package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
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

		// while the token does not equal EOF, obtain and print next token
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
