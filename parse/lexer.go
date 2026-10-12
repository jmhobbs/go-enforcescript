package parse

import "strings"

//go:generate ragel -Z -G2 -o lex.go lex.rl
//go:generate gofmt -w lex.go
type lexer struct {
	data        []byte
	p, pe, cs   int
	ts, te, act int
	err         error
}

func (lex *lexer) Error(e string) {
	line, col := lex.position(lex.ts)
	before, current, after := lex.context(line, 5)
	lex.err = ParseError{
		Message: e,
		Line:    line,
		Column:  col,
		Current: current,
		Before:  before,
		After:   after,
	}
}

func (lex *lexer) position(offset int) (line, col int) {
	line, col = 1, 1
	for _, b := range lex.data[:offset] {
		if b == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return line, col
}

func (lex *lexer) context(target, contextLines int) (before []string, line string, after []string) {
	lines := strings.Split(string(lex.data), "\n")

	idx := target - 1
	if idx < 0 || idx >= len(lines) {
		return nil, "", nil
	}

	start := max(0, idx-contextLines)
	before = lines[start:idx]

	line = lines[idx]

	end := min(len(lines), idx+1+contextLines)
	after = lines[idx+1 : end]

	return before, line, after
}

func (lex *lexer) commentIsInline() bool {
	for i := lex.ts - 1; i >= 0; i-- {
		switch lex.data[i] {
		// whitespace we can skip
		case ' ', '\t', '\r':
			continue
			// newline means we are the only thing on this line
		case '\n':
			return false
		default:
			// anything else should be part of another token
			return true
		}
	}
	// if we make it to the start of data, we are the first token
	return false
}
