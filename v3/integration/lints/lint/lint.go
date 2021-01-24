package lint

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

type Result struct {
	message       string
	codeCitations []string
	citations     []string
}

func NewResult(message string) *Result {
	return &Result{message: message}

}

func (r *Result) AddCodeCitation(start, end token.Pos, file *File) *Result {
	srcCode := make([]byte, end-start)
	reader := strings.NewReader(file.Src)
	reader.ReadAt(srcCode, int64(start))
	lineno := file.LineOf(start)
	citation := fmt.Sprintf("File %s, line %d\n\n%s\n\n", file.Name, lineno, string(srcCode))
	r.codeCitations = append(r.codeCitations, citation)
	return r
}

func (r *Result) SetCitations(citations ...string) *Result {
	r.citations = citations
	return r
}

func (r *Result) String() string {
	b := strings.Builder{}
	b.WriteString("--------------------\n")
	b.WriteString("Linting Error\n\n")
	b.WriteString(r.message)
	b.WriteString("\n\n")
	for _, code := range r.codeCitations {
		b.WriteString(code)
	}
	if len(r.citations) > 0 {
		b.WriteString("For more information, please see the following citations.\n")
	}
	for _, citation := range r.citations {
		b.WriteByte('\t')
		b.WriteString(citation)
		b.WriteByte('\n')
	}
	return b.String()
}

type File struct {
	Src   string
	Name  string
	Lines []string
}

type Lint interface {
	Lint(tree *ast.File, file *File) *Result
}

func (f *File) LineOf(pos token.Pos) int {
	start := 0
	end := 0
	for lineno, line := range f.Lines {
		start = end
		end = start + len(line)
		if int(pos) >= start && int(pos) <= end {
			return lineno + 1
		}
	}
	return int(token.NoPos)
}

func NewFile(name, src string) *File {
	return &File{src, name, strings.Split(src, "\n")}
}

func Parse(path string) (*ast.File, *File, error) {
	fset := new(token.FileSet)
	tree, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}
	file := NewFile(path, string(b))
	return tree, file, nil
}

func RunLint(path string, lint Lint) (*Result, error) {
	tree, file, err := Parse(path)
	if err != nil {
		return nil, err
	}
	return lint.Lint(tree, file), nil
}

func RunLints(path string, lints []Lint) ([]*Result, error) {
	tree, file, err := Parse(path)
	if err != nil {
		return nil, err
	}
	var results []*Result
	for _, lint := range lints {
		if result := lint.Lint(tree, file); result != nil {
			results = append(results, result)
		}
	}
	return results, nil
}
