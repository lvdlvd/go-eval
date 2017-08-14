//go:generate goyacc -o expr.go expr.y

// Expr is a simple expression evaluator that serves as a working example of
// how to use Go's yacc implementation.
package eval

import (
	"strings"
	"text/scanner"
)

func Parse(line string) (Evaluator, error) {
	l := &exprLex{}
	l.scan.Init(strings.NewReader(line))
	l.scan.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanComments | scanner.SkipComments
	p := yyParserImpl{}
	p.Parse(l)
	return l.val, l.err
}
