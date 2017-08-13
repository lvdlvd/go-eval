//go:generate goyacc -o expr.go expr.y

// Expr is a simple expression evaluator that serves as a working example of
// how to use Go's yacc implementation.
package eval

import (
	"log"
	"strings"
	"text/scanner"
)

func Parse(line string) (Evaluator, error) {
	l := &exprLex{}
	l.scan.Init(strings.NewReader(line))
	l.scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanComments | scanner.SkipComments
	p := yyParserImpl{}
	p.Parse(l)
	log.Println("final lval:", p.lval)
	for i := range p.stack {
		log.Println(p.stack[i])
	}
	//not sure why but for some reason the top of the parse tree is in stack[1], not [0].
	return p.stack[1].val, l.err
}
