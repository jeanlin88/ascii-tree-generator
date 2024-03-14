package mockfile

import (
	"io/fs"
	"time"
)

type MockFileInfo struct {
	name  string
	isDir bool
}

func NewMockFileInfo(name string, isDir bool) *MockFileInfo {
	return &MockFileInfo{name: name, isDir: isDir}
}

func (m MockFileInfo) Name() string {
	return m.name
}

func (m MockFileInfo) Size() int64 {
	return 0
}

func (m MockFileInfo) Mode() fs.FileMode {
	return 0
}

func (m MockFileInfo) ModTime() time.Time {
	return time.Now()
}

func (m MockFileInfo) IsDir() bool {
	return m.isDir
}

func (m MockFileInfo) Sys() any {
	return nil
}
