package cmdargs

import (
	"errors"
	"flag"
	"log"

	"github.com/jeanlin88/ascii-tree-generator/pkg/fileutils"
)

var (
	ErrOutputFileNotProvided = errors.New("output file not provided while replace set")
	ErrOutputFileExist       = errors.New("output file already existed")
	ErrInvalidOutputFile     = errors.New("invalid output file")
	ErrUnknownParameter      = errors.New("unknown parameter")

	ArgIncludeHidden          = "include-hidden"
	ArgIncludeHiddenUsage     = "include hidden file/directory"
	ArgOutputFile             = "output-file"
	ArgOutputFileUsage        = "output file name"
	ArgOutputFileUnset        = "OUTPUT_FILE_UNSET"
	ArgReplaceOutputFile      = "replace"
	ArgReplaceOutputFileUsage = "replace existing output file"
)

type CommandLineParser struct {
	fileSystem fileutils.FileSystem
}

func NewCommandLineParser(fs fileutils.FileSystem) *CommandLineParser {
	return &CommandLineParser{fileSystem: fs}
}

type CommandLineOptions struct {
	IncludeHidden     bool
	OutputFile        string
	ReplaceOutputFile bool
}

func (v *CommandLineParser) ParseOptions(args []string, flagSet *flag.FlagSet) (CommandLineOptions, error) {
	option := CommandLineOptions{}

	includeHiddenFlag := flagSet.Bool(ArgIncludeHidden, false, ArgIncludeHiddenUsage)
	outputFileFlag := flagSet.String(ArgOutputFile, ArgOutputFileUnset, ArgOutputFileUsage)
	replaceOutputFileFlag := flagSet.Bool(ArgReplaceOutputFile, false, ArgOutputFileUsage)

	if err := flagSet.Parse(args); err != nil {
		log.Println("parse failed")
		return option, err
	}

	if *outputFileFlag == "" {
		log.Println("empty output file")
		return option, ErrInvalidOutputFile
	}
	if *outputFileFlag == ArgOutputFileUnset && *replaceOutputFileFlag {
		return option, ErrOutputFileNotProvided
	}
	if *outputFileFlag != ArgOutputFileUnset &&
		v.fileSystem.FileExist(*outputFileFlag) &&
		!*replaceOutputFileFlag {
		log.Printf("output file %q already exist", *outputFileFlag)
		return option, ErrOutputFileExist
	}

	option.IncludeHidden = *includeHiddenFlag
	option.OutputFile = *outputFileFlag
	option.ReplaceOutputFile = *replaceOutputFileFlag
	return option, nil
}
