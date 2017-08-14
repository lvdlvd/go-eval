package eval

import (
	"fmt"
	"strconv"
	"text/scanner"
)

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

// The parser uses the type <prefix>Lex as a lexer. It must provide
// the methods Lex(*<prefix>SymType) int and Error(string).
type exprLex struct {
	scan scanner.Scanner
	err  error
	val  Evaluator
}

// The parser calls this method to get each new token.
func (x *exprLex) Lex(yylval *yySymType) int {
	// 	v := x.lex(yylval)
	// 	log.Printf("lex: %d %v", v, yylval)
	// 	return v
	// }
	// func (x *exprLex) lex(yylval *yySymType) int {
	r := x.scan.Scan()
	switch r {
	case scanner.Ident:
		yylval.val = varval(x.scan.TokenText())
		return VAR
	case scanner.Int:
		v, _ := strconv.ParseInt(x.scan.TokenText(), 0, 64)
		yylval.val = constval(float64(v))
		return CONST
	case scanner.Float:
		v, _ := strconv.ParseFloat(x.scan.TokenText(), 64)
		yylval.val = constval(v)
		return CONST
	}
	return int(r)
}

// The parser calls this method on a parse error.
func (x *exprLex) Error(s string) {
	x.err = fmt.Errorf("%v %q :parse error: %s", x.scan.Pos(), x.scan.TokenText(), s)
}
