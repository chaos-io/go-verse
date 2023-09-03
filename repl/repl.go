// Package repl REPL: Read-Evaluate-Print Loop(读取-求值-打印循环)
// RLPL: Read-Lex-Print Loop(读取-词法分析-打印循环)
// RPPL: Read-Parse-Print Loop(读取-语法分析-打印循环) now
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/chaos-io/go-verse/lexer"
	"github.com/chaos-io/go-verse/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok)
		// }
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "\nparser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
