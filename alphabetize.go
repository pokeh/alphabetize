package alphabetize

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "alphabetize is a linter to sort struct fields in alphabetical order"

var Analyzer = &analysis.Analyzer{
	Name: "alphabetize",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	/*
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

		nodeFilter := []ast.Node{
			(*ast.Ident)(nil),
		}

		inspect.Preorder(nodeFilter, func(n ast.Node) {
			switch n := n.(type) {
			case *ast.Ident:
				if n.Name == "gopher" {
					pass.Reportf(n.Pos(), "identifier is gopher")
				}
			}
		})
	*/

	// TODO: find struct declarations from the nodes

	// TODO: iterate through its fields

	// TODO: if field name is alphabetically before the previous field
	// pass.Reportf(n.Pos(), "not alphabetized")

	return nil, nil
}
