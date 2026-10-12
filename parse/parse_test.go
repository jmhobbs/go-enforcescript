package parse_test

import (
	"testing"

	"github.com/jmhobbs/go-enforcescript/ast"
	"github.com/jmhobbs/go-enforcescript/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ParseSmokeTest(t *testing.T) {
	input := []byte(`void Hello()
{
}

void Hello2()
{
}`)

	output, err := parse.Parse(input, false)
	require.NoError(t, err)
	assert.Equal(
		t,
		ast.File{
			Entries: []ast.ClassOrFunction{
				ast.Function{
					ReturnType: ast.TypeVoid,
					Identifier: "Hello",
					Arguments:  []ast.Argument{},
					Body: ast.Block{
						Statements: []ast.Statement{},
					},
				},
				ast.Function{
					ReturnType: ast.TypeVoid,
					Identifier: "Hello2",
					Arguments:  []ast.Argument{},
					Body: ast.Block{
						Statements: []ast.Statement{},
					},
				},
			},
		},
		output,
	)
}
