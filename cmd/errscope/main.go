package main

import (
	"github.com/kimuson13/errscope"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(errscope.Analyzer) }
