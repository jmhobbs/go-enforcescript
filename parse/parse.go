package parse

import "github.com/jmhobbs/go-enforcescript/ast"

var file ast.File

//go:generate go tool goyacc -o cpp.go cpp.y
//go:generate gofmt -w cpp.go
func Parse(input []byte, debug bool) (ast.File, error) {
	file = ast.File{}

	yyErrorVerbose = true
	if debug {
		yyDebug = 5
	}
	lex := newLexer(input)
	if yyParse(lex) != 0 {
		return file, lex.err
	}
	return file, nil
}
