package repl

import (
	"bufio"
	"fmt"
	"intergo/pkg/lexer"
	"intergo/pkg/token"
	"io"
)

const INPUT = ">> "

func Begin(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(INPUT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
