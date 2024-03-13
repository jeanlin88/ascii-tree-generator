package tree

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/jeanlin88/ascii-tree-generator/mocks/file"
	"github.com/jeanlin88/ascii-tree-generator/tests"
)

func TestTreeBuilder(t *testing.T) {
	t.Run("mock file system", func(t *testing.T) {
		file1Mock := mockfile.NewMockFileInfo("file1", false)
		dir1Mock := mockfile.NewMockFileInfo("directory1", true)
		fsMap := map[string][]fs.DirEntry{
			".": {
				fs.FileInfoToDirEntry(file1Mock),
				fs.FileInfoToDirEntry(dir1Mock),
			},
			"directory1": {},
		}
		fsMock := mockfile.NewMockFileSystem(fsMap, nil, nil)
		builder := NewTreeBuilder(fsMock)

		expect := TreeNode{
			Type: DIRECTORY,
			Name: "project-root",
			Children: &[]TreeNode{
				{Type: FILE, Name: "file1"},
				{Type: DIRECTORY, Name: "directory1", Children: &[]TreeNode{}},
			},
		}
		get, err := builder.execute(".")
		tests.NoError(t, err)
		EqualTreeNode(t, expect, get)
	})

	t.Run("ignore hidden file/directory", func(t *testing.T) {
		hiddenFileMock := mockfile.NewMockFileInfo(".hidden", false)
		fsMap := map[string][]fs.DirEntry{
			".": {
				fs.FileInfoToDirEntry(hiddenFileMock),
			},
		}
		fsMock := mockfile.NewMockFileSystem(fsMap, nil, nil)
		builder := NewTreeBuilder(fsMock)

		expect := TreeNode{
			Type:     DIRECTORY,
			Name:     "project-root",
			Children: &[]TreeNode{},
		}
		get, err := builder.execute(".")
		tests.NoError(t, err)
		EqualTreeNode(t, expect, get)
	})

	t.Run("Getwd error", func(t *testing.T) {
		getwdErr := errors.New("Getwd failed")
		fsMock := mockfile.NewMockFileSystem(nil, getwdErr, nil)
		builder := NewTreeBuilder(fsMock)

		_, err := builder.execute(".")
		tests.EqualError(t, getwdErr, err)
	})

	t.Run("ReadDir error", func(t *testing.T) {
		dir1Mock := mockfile.NewMockFileInfo("directory1", true)
		fsMap := map[string][]fs.DirEntry{
			".": {
				fs.FileInfoToDirEntry(dir1Mock),
			},
		}
		fsMock := mockfile.NewMockFileSystem(fsMap, nil, []string{"directory1"})
		builder := NewTreeBuilder(fsMock)

		_, err := builder.execute(".")
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
