package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rezbow/lang/ast"
)

func Test_parse_expression_statement(t *testing.T) {
	got := Parse("8 + 3 * 3 * (2 - 1)")
	want := &ast.AST{
		Root: []ast.Node{
			&ast.NodeBinaryOperator{
				Left:  &ast.NodeNumber{N: 8},
				Right: &ast.NodeNumber{N: 3},
				Op:    "+",
			},
		},
	}
	assertAstEqual(t, got, want)
}

func assertAstEqual(t testing.TB, got, want *ast.AST) {
	t.Helper()
	if len(got.Root) != len(want.Root) {
		t.Fatalf("different length of root elements, got %d, wanted %d", len(got.Root), len(want.Root))
	}
	for idx, n := range got.Root {
		equal := cmp.Equal(n, want.Root[idx], cmp.Comparer(func(a, b ast.Node) bool {
			return a.Equal(b)
		}))
		if !equal {
			t.Errorf("got %q, wanted %q", n.String(), want.Root[idx].String())
		}
	}
}
