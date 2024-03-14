package fileutils

import (
	"io/fs"
	"os"
)

type FileSystem interface {
	Getwd() (string, error)
	ReadDir(name string) ([]fs.DirEntry, error)
	FileExist(name string) bool
}

type OSFileSystem struct{}

func (fs *OSFileSystem) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(name)
}

func (fs *OSFileSystem) Getwd() (string, error) {
	return os.Getwd()
}

func (fs *OSFileSystem) FileExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
