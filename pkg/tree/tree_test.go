package tree

import (
	"fmt"
	"testing"

	"github.com/jeanlin88/ascii-tree-generator/tests"
)

const (
	dir2Name  = "dir2"
	dir3Name  = "dir3"
	file2Name = "file2"
	file3Name = "file3"
	file4Name = "file4"
)

func TestTree(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		structure := TreeNode{}
		expect := ``
		get, err := structure.ToAsciiTree("")
		tests.NoError(t, err)
		tests.EqualString(t, expect, get)
	})
	t.Run("two-layer", func(t *testing.T) {
		structure := TreeNode{
			Type: DIRECTORY,
			Name: rootName,
			Children: &[]TreeNode{
				{
					Type: FILE,
					Name: file1Name,
				},
			},
		}
		expect := `project-root/
└── file1`
		get, err := structure.ToAsciiTree("")
		tests.NoError(t, err)
		tests.EqualString(t, expect, get)
	})
	t.Run("three-layer", func(t *testing.T) {
		structure := TreeNode{
			Type: DIRECTORY,
			Name: rootName,
			Children: &[]TreeNode{
				{
					Type: FILE,
					Name: file1Name,
				},
				{
					Type: DIRECTORY,
					Name: dir1Name,
					Children: &[]TreeNode{
						{
							Type: DIRECTORY,
							Name: dir2Name,
						},
					},
				},
				{
					Type: DIRECTORY,
					Name: dir3Name,
					Children: &[]TreeNode{
						{
							Type: FILE,
							Name: file2Name,
						},
					},
				},
			},
		}
		expect := `project-root/
├── file1
├── dir1/
│   └── dir2/
└── dir3/
    └── file2`
		get, err := structure.ToAsciiTree("")
		tests.NoError(t, err)
		tests.EqualString(t, expect, get)
	})
	t.Run("four-layer", func(t *testing.T) {
		structure := TreeNode{
			Type: DIRECTORY,
			Name: rootName,
			Children: &[]TreeNode{
				{
					Type: FILE,
					Name: file1Name,
				},
				{
					Type: DIRECTORY,
					Name: dir1Name,
					Children: &[]TreeNode{
						{
							Type: DIRECTORY,
							Name: dir2Name,
							Children: &[]TreeNode{
								{
									Type: FILE,
									Name: file4Name,
								},
							},
						},
						{
							Type: FILE,
							Name: file3Name,
						},
					},
				},
				{
					Type: FILE,
					Name: file2Name,
				},
			},
		}
		expect := `project-root/
├── file1
├── dir1/
│   ├── dir2/
│   │   └── file4
│   └── file3
└── file2`
		get, err := structure.ToAsciiTree("")
		tests.NoError(t, err)
		tests.EqualString(t, expect, get)
	})
	t.Run("invalid-indent", func(t *testing.T) {
		invalidIndent := "+-- "
		structure := TreeNode{}
		_, err := structure.ToAsciiTree(invalidIndent)
		tests.Error(t, err)
		tests.EqualString(t, fmt.Sprintf("invalid indent %q", invalidIndent), err.Error())
	})
}
