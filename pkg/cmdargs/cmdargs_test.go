package cmdargs

import (
	"flag"
	"fmt"
	"testing"

	mockfile "github.com/jeanlin88/ascii-tree-generator/mocks/file"
	"github.com/jeanlin88/ascii-tree-generator/tests"
)

var (
	existFile  = "exist.txt"
	outputFile = "output.txt"
)

func TestCommandLineParser(t *testing.T) {
	NewTestFlagSet := func() *flag.FlagSet {
		return flag.NewFlagSet("test", flag.ContinueOnError)
	}
	fsMock := mockfile.NewMockFileSystem(nil, nil, []string{existFile})
	parser := NewCommandLineParser(fsMock)
	t.Run("default options", func(t *testing.T) {
		args := []string{}
		expect := CommandLineOptions{
			IncludeHidden:     false,
			OutputFile:        ArgOutputFileUnset,
			ReplaceOutputFile: false,
		}
		get, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.NoError(t, err)
		EqualOptions(t, expect, get)
	})
	t.Run("include hidden file", func(t *testing.T) {
		args := []string{fmt.Sprintf("--%s", ArgIncludeHidden)}
		expect := CommandLineOptions{
			IncludeHidden:     true,
			OutputFile:        ArgOutputFileUnset,
			ReplaceOutputFile: false,
		}
		get, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.NoError(t, err)
		EqualOptions(t, expect, get)
	})
	t.Run("setting output file", func(t *testing.T) {
		args := []string{fmt.Sprintf("--%s=%s", ArgOutputFile, outputFile)}
		expect := CommandLineOptions{
			IncludeHidden:     false,
			OutputFile:        outputFile,
			ReplaceOutputFile: false,
		}
		get, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.NoError(t, err)
		EqualOptions(t, expect, get)
	})
	t.Run("replace existing output file", func(t *testing.T) {
		outputFileName := "output.txt"
		args := []string{
			fmt.Sprintf("--%s=%s", ArgOutputFile, outputFileName),
			fmt.Sprintf("--%s", ArgReplaceOutputFile),
		}
		expect := CommandLineOptions{
			IncludeHidden:     false,
			OutputFile:        outputFileName,
			ReplaceOutputFile: true,
		}
		get, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.NoError(t, err)
		EqualOptions(t, expect, get)
	})
	t.Run("replace set but no output file provided", func(t *testing.T) {
		args := []string{
			fmt.Sprintf("--%s", ArgReplaceOutputFile),
		}
		_, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.EqualError(t, ErrOutputFileNotProvided, err)
	})
	t.Run("output file already exist without replace flag", func(t *testing.T) {
		args := []string{
			fmt.Sprintf("--%s=%s", ArgOutputFile, existFile),
		}
		_, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.EqualError(t, ErrOutputFileExist, err)
	})
	t.Run("invalid output file: empty string", func(t *testing.T) {
		args := []string{fmt.Sprintf("--%s=", ArgOutputFile)}
		_, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.EqualError(t, ErrInvalidOutputFile, err)
	})
	t.Run("unknown parameter", func(t *testing.T) {
		args := []string{"--unknown"}
		_, err := parser.ParseOptions(args, NewTestFlagSet())
		tests.Error(t, err)
	})
}

func EqualOptions(t *testing.T, expect, get CommandLineOptions) {
	t.Helper()
	if expect != get {
		t.Errorf("expect %v get %v", expect, get)
		t.FailNow()
	}
}
