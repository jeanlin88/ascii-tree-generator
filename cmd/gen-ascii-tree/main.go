package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jeanlin88/ascii-tree-generator/pkg/cmdargs"
	"github.com/jeanlin88/ascii-tree-generator/pkg/fileutils"
	"github.com/jeanlin88/ascii-tree-generator/pkg/tree"
)

func main() {
	fs := fileutils.OSFileSystem{}
	parser := cmdargs.NewCommandLineParser(&fs)
	options, err := parser.ParseOptions(
		os.Args[1:],
		flag.NewFlagSet("main", flag.ContinueOnError),
	)
	if err != nil {
		os.Exit(1)
	}

	builder := tree.NewTreeBuilder(&fs, options)
	treeRoot, err := builder.Execute(".")
	if err != nil {
		os.Exit(1)
	}

	asciiTree, err := treeRoot.ToAsciiTree("")
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(asciiTree)
}
