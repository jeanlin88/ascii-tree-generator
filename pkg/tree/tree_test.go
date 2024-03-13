package tree

import (
	"fmt"
	"testing"

	"github.com/jeanlin88/ascii-tree-generator/tests"
)

func TestTree(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		structure := TreeNode{}
		expect := ``
		get, err := structure.toAsciiTree("")
		tests.NoError(t, err)
		tests.EqualString(t, expect, get)
	})
	t.Run("two-layer", func(t *testing.T) {
		structure := TreeNode{
			Type: DIRECTORY,
			Name: "project-root",
			Children: &[]TreeNode{
				{
					Type: FILE,
					Name: "file1",
				},
			},
		}
		expect := `project-root/
└── file1`
		get, err := structure.toAsciiTree("")
		tests.NoError(t, err)
		tests.EqualString(t, expect, get)
	})
	t.Run("three-layer", func(t *testing.T) {
		structure := TreeNode{
			Type: DIRECTORY,
			Name: "project-root",
			Children: &[]TreeNode{
				{
					Type: FILE,
					Name: "file1",
				},
				{
					Type: DIRECTORY,
					Name: "dir2",
					Children: &[]TreeNode{
						{
							Type: DIRECTORY,
							Name: "dir3",
						},
					},
				},
				{
					Type: DIRECTORY,
					Name: "dir4",
					Children: &[]TreeNode{
						{
							Type: FILE,
							Name: "file2",
						},
					},
				},
			},
		}
		expect := `project-root/
├── file1
├── dir2/
│   └── dir3/
└── dir4/
    └── file2`
		get, err := structure.toAsciiTree("")
		tests.NoError(t, err)
		tests.EqualString(t, expect, get)
	})
	t.Run("invalid-indent", func(t *testing.T) {
		invalidIndent := "+-- "
		structure := TreeNode{}
		_, err := structure.toAsciiTree(invalidIndent)
		tests.Error(t, err)
		tests.EqualString(t, fmt.Sprintf("invalid indent %q", invalidIndent), err.Error())
	})
}
