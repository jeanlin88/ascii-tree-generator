package tree

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type TreeNodeType string

const (
	DIRECTORY TreeNodeType = "directory"
	FILE      TreeNodeType = "file"

	INDENT_FORMAT   = `^(((│ {3}| {4})( {4})*)?(└── |├── ))?$`
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

func (t *TreeNode) toAsciiTree(indent string) (string, error) {
	if !t.indentValid(indent) {
		return "", fmt.Errorf("invalid indent %q", indent)
	}

	name := indent + t.generateName()
	if t.Children == nil {
		return name, nil
	}

	result := []string{name}
	lastIdx := len(*t.Children) - 1
	for idx, child := range *t.Children {
		if indent != "" {
			indentSuffix := strings.Repeat(" ", (utf8.RuneCountInString(indent) - 1))
			if strings.HasPrefix(indent, "├") {
				indent = "│" + indentSuffix
			} else {
				indent = " " + indentSuffix
			}
		}
		subIndent := indent
		if idx == lastIdx {
			subIndent += LAST_INDENT
		} else {
			subIndent += NON_LAST_INDENT
		}
		subAsciiTree, _ := child.toAsciiTree(subIndent)
		result = append(result, subAsciiTree)
	}
	return strings.Join(result, "\n"), nil
}
