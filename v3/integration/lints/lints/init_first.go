package lints

import (
	"go/ast"
	"go/token"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

type InitFirst struct{}

func (i *InitFirst) Lint(tree *ast.File, file *lint.File) *lint.Result {
	functions := filter(tree.Decls, func(decl ast.Decl) bool {
		_, ok := decl.(*ast.FuncDecl)
		return ok
	})
	if len(functions) == 0 {
		return lint.NewResult("Lint does not contain any functions or methods").AddCodeCitation(token.NoPos, token.NoPos, file)
	}
	function := functions[0].(*ast.FuncDecl)
	if isInitialize(function) {
		return nil
	}
	return lint.NewResult("Got the wrong method signature as the first function declaration within the linter.\n"+
		"ZLint lints must have func (l *LinterStruct) Initialize() error { ... } as their first function declaration").
		AddCodeCitation(function.Pos(), function.End(), file).
		SetCitations("https://github.com/zmap/zlint/issues/371")

}

func isInitialize(function *ast.FuncDecl) bool {
	notInit := function.Name.Name != "Initialize"
	isNotAMethod := function.Recv == nil
	hasParameters := len(function.Type.Params.List) != 0
	wrongNumberOfReturns := function.Type.Results == nil || len(function.Type.Results.List) != 1
	if notInit || isNotAMethod || hasParameters || wrongNumberOfReturns {
		return false
	}
	ret := function.Type.Results.List[0]
	identifier, ok := ret.Type.(*ast.Ident)
	if !ok {
		return false
	}
	return identifier.Name == "error"
}

func filter(decls []ast.Decl, predicate func(decl ast.Decl) bool) (filtered []ast.Decl) {
	for _, decl := range decls {
		if predicate(decl) {
			filtered = append(filtered, decl)
		}
	}
	return
}
