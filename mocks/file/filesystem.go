package mockfile

import (
	"errors"
	"fmt"
	"io/fs"
)

var (
	ErrReadDirFailed = errors.New("ReadDir failed")
)

type DirEntriesMap map[string][]fs.DirEntry

type MockFileSystem struct {
	fileSystemMap DirEntriesMap
	getwdError    error
	readDirError  error
	failMap       map[string]struct{}
}

func NewMockFileSystem(fsMap DirEntriesMap, getwdErr error, fails []string) *MockFileSystem {
	failMap := map[string]struct{}{}
	for _, fail := range fails {
		failMap[fail] = struct{}{}
	}
	return &MockFileSystem{
		fileSystemMap: fsMap,
		getwdError:    getwdErr,
		readDirError:  ErrReadDirFailed,
		failMap:       failMap,
	}
}

func (m *MockFileSystem) Getwd() (string, error) {
	if m.getwdError != nil {
		return "", m.getwdError
	}
	return "project-root", nil
}

func (m *MockFileSystem) ReadDir(name string) ([]fs.DirEntry, error) {
	if _, ok := m.failMap[name]; ok {
		return nil, m.readDirError
	}

	entries, ok := m.fileSystemMap[name]
	if !ok {
		return nil, fmt.Errorf("entries of %s not found", name)
	}
	return entries, nil
}

func (m *MockFileSystem) FileExist(name string) bool {
	_, ok := m.failMap[name]
	return ok
}
