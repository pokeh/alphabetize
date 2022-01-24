package main

import (
	"github.com/pokeh/alphabetize"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	// using unitchecker allows the analyzer to be integrated into existing linting tools e.g. go vet and golangci-lint
	unitchecker.Main(alphabetize.Analyzer)
}
