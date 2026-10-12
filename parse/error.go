package parse

import (
	"fmt"
	"regexp"
	"strings"
)

type ParseError struct {
	Message string
	Line    int
	Column  int
	Current string
	Before  []string
	After   []string
}

var tabRegexp = regexp.MustCompile(`^(\t)+`)

func (p ParseError) Error() string {
	maxLine := p.Line + len(p.After)
	prefixWidth := len(fmt.Sprintf("%d", maxLine))
	lineFmtString := fmt.Sprintf("%%%dd | %%s", prefixWidth)

	strs := []string{
		fmt.Sprintf("parsing error at line %d, column %d", p.Line, p.Column),
		"",
	}

	for i, line := range p.Before {
		strs = append(strs, fmt.Sprintf(lineFmtString, p.Line-len(p.Before)+i, line))
	}

	strs = append(strs, fmt.Sprintf(lineFmtString, p.Line, p.Current))
	tabCount := len(tabRegexp.FindString(p.Current))
	strs = append(
		strs,
		strings.Repeat("!", prefixWidth)+" | "+strings.Repeat("\t", tabCount)+strings.Repeat(" ", p.Column-tabCount-1)+"`-- "+replaceTokensInErrorMessage(p.Message),
	)

	for i, line := range p.After {
		strs = append(strs, fmt.Sprintf(lineFmtString, p.Line+i+1, line))
	}

	return strings.Join(strs, "\n")
}

func replaceTokensInErrorMessage(msg string) string {
	mapping := map[string]string{
		"TOK_ARRAY":       `"[]"`,
		"TOK_BLOCK_OPEN":  `"{"`,
		"TOK_BLOCK_CLOSE": `"}"`,
		"TOK_SEMICOLON":   `";"`,
		"TOK_ASSIGN":      `"="`,
		"TOK_QUOTE":       `'"'`,
		"TOK_COMMA":       `","`,
		"TOK_COLON":       `":"`,
		"TOK_DEFINE":      `"#define"`,
	}

	for from, to := range mapping {
		msg = strings.ReplaceAll(msg, from, to)
	}

	return msg
}
