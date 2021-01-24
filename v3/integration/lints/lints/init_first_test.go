package lints

import (
	"os"
	"testing"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

func TestInitFirst_Lint(t *testing.T) {
	t.Log(os.Getwd())
	data := map[string]bool{
		"testdata/initializeFirst.go":            true,
		"testdata/initializeNotFirst.go":         false,
		"testdata/initializeFirstNoFunctions.go": false,
	}
	l := &InitFirst{}
	for file, want := range data {
		t.Run(file, func(t *testing.T) {
			r, err := lint.RunLint(file, l)
			if err != nil {
				t.Error(err)
				return
			}
			if want && r != nil {
				t.Errorf("got unexepcted error result, %s", r)
			} else if !want && r == nil {
				t.Errorf("expected failure but got nothing")
			}
		})

	}
}
