package main

import (
	"flag"
	"log"
	"os"

	"github.com/jeanlin88/ascii-tree-generator/pkg/cmdargs"
	"github.com/jeanlin88/ascii-tree-generator/pkg/fileutils"
)

func main() {
	parser := cmdargs.NewCommandLineParser(&fileutils.OSFileSystem{})
	options, err := parser.ParseOptions(
		os.Args[1:],
		flag.NewFlagSet("main", flag.ContinueOnError),
	)
	if err != nil {
		os.Exit(1)
	}
	log.Printf("options: %v", options)
}
