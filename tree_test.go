package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		structure := TreeNode{}
		expect := ``
		get, err := structure.toAsciiTree("")
		NoError(t, err)
		EqualString(t, expect, get)
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
		NoError(t, err)
		EqualString(t, expect, get)
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
					Type: FILE,
					Name: "file2",
				},
			},
		}
		expect := `project-root/
├── file1
├── dir2/
│   └── dir3/
└── file2`
		get, err := structure.toAsciiTree("")
		NoError(t, err)
		EqualString(t, expect, get)
	})
	t.Run("invalid-indent", func(t *testing.T) {
		invalidIndent := "+-- "
		structure := TreeNode{}
		_, err := structure.toAsciiTree(invalidIndent)
		Error(t, err)
		EqualString(t, fmt.Sprintf("invalid indent %q", invalidIndent), err.Error())
	})
}

func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expect error get nil")
		t.FailNow()
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got unexpected error %q", err.Error())
		t.FailNow()
	}
}

func EqualError(t *testing.T, expect, get error) {
	t.Helper()
	if expect != get {
		t.Errorf("expect %q get %q", expect, get)
		t.FailNow()
	}
}

func EqualString(t *testing.T, expect, get string) {
	t.Helper()
	if expect != get {
		t.Errorf("expect %q get %q", expect, get)
		t.FailNow()
	}
}
