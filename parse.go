// Package enforcescript provides an entry point for parsing Bohemia Interactive .c files.
//
// The parser here is based on the docs at https://community.bistudio.com/wiki/DayZ:Enforce_Script_Syntax
// as well as real .c files seen "in the wild" of DayZ.
package enforcescript

import (
	"github.com/jmhobbs/go-enforcescript/ast"
	"github.com/jmhobbs/go-enforcescript/parse"
)

func Parse(input []byte) (ast.File, error) {
	return parse.Parse(input, false)
}

func ParseWithDebug(input []byte) (ast.File, error) {
	return parse.Parse(input, false)
}
