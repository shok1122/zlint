package main // import "github.com/zmap/zlint/v3/integration/lints"

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zmap/zlint/v3/integration/lints/lint"

	"github.com/zmap/zlint/v3/integration/lints/lints"
)

var Linters = []lint.Lint{
	&lints.InitFirst{},
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE ./main <path to lint directory>\n")
		os.Exit(1)
	}
	results, err := run(os.Args[1])
	if err != nil {
		fmt.Printf("A fatal error has occurred: %v\n", err)
		os.Exit(2)
	}
	exitCode := 0
	if len(results) != 0 {
		exitCode = 1
		fmt.Printf("Found %d linting errors\n", len(results))
	}
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
	os.Exit(exitCode)
}

func run(dir string) ([]*lint.Result, error) {
	var results []*lint.Result
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !isAGoFile(info) {
			return nil
		}
		r, err := lint.RunLints(path, Linters)
		if err != nil {
			return err
		}
		results = append(results, r...)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func isAGoFile(info os.FileInfo) bool {
	return !info.IsDir() && strings.HasSuffix(info.Name(), ".go")
}
