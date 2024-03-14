package tree

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/jeanlin88/ascii-tree-generator/pkg/cmdargs"
	"github.com/jeanlin88/ascii-tree-generator/pkg/fileutils"
)

type TreeBuilder struct {
	fileSystem fileutils.FileSystem
	options    cmdargs.CommandLineOptions
}

func NewTreeBuilder(fs fileutils.FileSystem, options cmdargs.CommandLineOptions) *TreeBuilder {
	return &TreeBuilder{
		fileSystem: fs,
		options:    options,
	}
}

func (tb *TreeBuilder) Execute(path string) (TreeNode, error) {
	root := TreeNode{}

	rootDir, err := tb.fileSystem.Getwd()
	if err != nil {
		log.Println("Getwd failed")
		return root, err
	}

	children, err := tb.getSubTrees(path)
	if err != nil {
		log.Println("getSubTrees failed")
		return root, err
	}

	root = TreeNode{
		Type:     DIRECTORY,
		Name:     filepath.Base(rootDir),
		Children: &children,
	}
	return root, nil
}

func (tb *TreeBuilder) getSubTrees(path string) ([]TreeNode, error) {
	entries, err := tb.fileSystem.ReadDir(path)
	if err != nil {
		log.Println("ReadDir failed")
		return nil, err
	}

	subTrees := []TreeNode{}
	for _, entry := range entries {
		name := entry.Name()
		if !tb.options.IncludeHidden && strings.HasPrefix(name, ".") {
			log.Printf("ignore hidden file/directory: %s\n", name)
			continue
		}

		node := TreeNode{
			Type: FILE,
			Name: entry.Name(),
		}
		if entry.IsDir() {
			node.Type = DIRECTORY
			children, err := tb.getSubTrees(filepath.Join(path, node.Name))
			if err != nil {
				log.Println("getSubTrees failed")
				return nil, err
			}
			node.Children = &children
		}
		subTrees = append(subTrees, node)
	}
	return subTrees, nil
}
