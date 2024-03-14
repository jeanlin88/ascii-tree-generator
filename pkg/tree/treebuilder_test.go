package tree

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/jeanlin88/ascii-tree-generator/mocks/file"
	"github.com/jeanlin88/ascii-tree-generator/pkg/cmdargs"
	"github.com/jeanlin88/ascii-tree-generator/tests"
)

const (
	rootName       = "project-root"
	dir1Name       = "dir1"
	file1Name      = "file1"
	hiddenDirName  = ".dir"
	hiddenFileName = ".file"
)

func TestTreeBuilder(t *testing.T) {
	dir1Mock := mockfile.NewMockFileInfo(dir1Name, true)
	file1Mock := mockfile.NewMockFileInfo(file1Name, false)
	hiddenDirMock := mockfile.NewMockFileInfo(hiddenDirName, true)
	hiddenFileMock := mockfile.NewMockFileInfo(hiddenFileName, false)
	fsMap := map[string][]fs.DirEntry{
		".": {
			fs.FileInfoToDirEntry(hiddenDirMock),
			fs.FileInfoToDirEntry(dir1Mock),
			fs.FileInfoToDirEntry(file1Mock),
		},
		dir1Name: {},
		hiddenDirName: {
			fs.FileInfoToDirEntry(hiddenFileMock),
		},
	}

	t.Run("default options", func(t *testing.T) {
		fsMock := mockfile.NewMockFileSystem(fsMap, nil, nil)
		options := cmdargs.CommandLineOptions{
			IncludeHidden:     false,
			OutputFile:        cmdargs.ArgOutputFileUnset,
			ReplaceOutputFile: false,
		}
		builder := NewTreeBuilder(fsMock, options)

		expect := TreeNode{
			Type: DIRECTORY,
			Name: rootName,
			Children: &[]TreeNode{
				{Type: DIRECTORY, Name: dir1Name, Children: &[]TreeNode{}},
				{Type: FILE, Name: file1Name},
			},
		}
		get, err := builder.Execute(".")
		tests.NoError(t, err)
		EqualTreeNode(t, expect, get)
	})

	t.Run("include hidden", func(t *testing.T) {
		fsMock := mockfile.NewMockFileSystem(fsMap, nil, nil)
		options := cmdargs.CommandLineOptions{
			IncludeHidden:     true,
			OutputFile:        cmdargs.ArgOutputFileUnset,
			ReplaceOutputFile: false,
		}
		builder := NewTreeBuilder(fsMock, options)

		expect := TreeNode{
			Type: DIRECTORY,
			Name: rootName,
			Children: &[]TreeNode{
				{Type: DIRECTORY, Name: hiddenDirName, Children: &[]TreeNode{
					{Type: FILE, Name: hiddenFileName},
				}},
				{Type: DIRECTORY, Name: dir1Name, Children: &[]TreeNode{}},
				{Type: FILE, Name: file1Name},
			},
		}
		get, err := builder.Execute(".")
		tests.NoError(t, err)
		EqualTreeNode(t, expect, get)
	})

	t.Run("Getwd error", func(t *testing.T) {
		getwdErr := errors.New("Getwd failed")
		fsMock := mockfile.NewMockFileSystem(nil, getwdErr, nil)
		builder := NewTreeBuilder(fsMock, cmdargs.CommandLineOptions{})

		_, err := builder.Execute(".")
		tests.EqualError(t, getwdErr, err)
	})

	t.Run("ReadDir error", func(t *testing.T) {
		fsMock := mockfile.NewMockFileSystem(fsMap, nil, []string{dir1Name})
		builder := NewTreeBuilder(fsMock, cmdargs.CommandLineOptions{})

		_, err := builder.Execute(".")
		tests.EqualError(t, mockfile.ErrReadDirFailed, err)
	})
}

func EqualTreeNode(t *testing.T, expect, get TreeNode) {
	if expect.Type != get.Type {
		t.Errorf("expect type %s get type %s", expect.Type, get.Type)
		t.FailNow()
	}
	if expect.Name != get.Name {
		t.Errorf("expect name %s get name %s", expect.Name, get.Name)
		t.FailNow()
	}
	if expect.Children == nil && get.Children == nil {
		return
	} else if expect.Children == nil || get.Children == nil || len(*expect.Children) != len(*get.Children) {
		t.Errorf("expect children %v get children %v", expect.Children, get.Children)
		t.FailNow()
	} else {
		for idx, node := range *expect.Children {
			EqualTreeNode(t, node, (*get.Children)[idx])
		}
	}
}
