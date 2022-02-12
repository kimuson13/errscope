package main

import (
	"github.com/kimuson13/errscope"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(errscope.Analyzer) }
