package tree

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	//"unicode/utf8"
)

type TreeNodeType string

const (
	DIRECTORY TreeNodeType = "directory"
	FILE      TreeNodeType = "file"

	INDENT_FORMAT   = `^((│ {3}| {4}))*((└── |├── ))?$`
	LAST_INDENT     = "└── "
	NON_LAST_INDENT = "├── "
)

type TreeNode struct {
	Type     TreeNodeType
	Name     string
	Children *[]TreeNode
}

func (t *TreeNode) generateName() string {
	if t.Type == DIRECTORY {
		return t.Name + "/"
	}
	return t.Name
}

func (t *TreeNode) indentValid(indent string) bool {
	re := regexp.MustCompile(INDENT_FORMAT)
	return indent == "" || re.Match([]byte(indent))
}

func (t *TreeNode) ToAsciiTree(indent string) (string, error) {
	if !t.indentValid(indent) {
		log.Printf("indent invalid: %q", indent)
		return "", fmt.Errorf("invalid indent %q", indent)
	}

	name := indent + t.generateName()
	if t.Children == nil {
		return name, nil
	}

	result := []string{name}
	lastIdx := len(*t.Children) - 1
	for idx, child := range *t.Children {
		indent = strings.ReplaceAll(indent, NON_LAST_INDENT, "│   ")
		indent = strings.ReplaceAll(indent, LAST_INDENT, "    ")
		subIndent := indent
		if idx == lastIdx {
			subIndent += LAST_INDENT
		} else {
			subIndent += NON_LAST_INDENT
		}
		subAsciiTree, _ := child.ToAsciiTree(subIndent)
		result = append(result, subAsciiTree)
	}
	return strings.Join(result, "\n"), nil
}
