package utils

import (
	"strings"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

func IsALint(file *lint.File) bool {
	return strings.HasPrefix(file.Name, "lint_") && IsAGoFile(file) && !IsATest(file)
}

func IsAGoFile(file *lint.File) bool {
	return strings.HasSuffix(file.Name, ".go")
}

func IsATest(file *lint.File) bool {
	return strings.HasSuffix(file.Name, "test.go")
}
