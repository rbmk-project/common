package mocks

import (
	"io/fs"
	"net"
	"time"

	"github.com/rbmk-project/common/internal/fsmodel"
)

// FsmodelFS implements [fsmodel.FS] for testing
type FsmodelFS struct {
	// MockChmod implements Chmod
	MockChmod func(name string, mode fs.FileMode) error

	// MockChown implements Chown
	MockChown func(name string, uid, gid int) error

	// MockChtimes implements Chtimes
	MockChtimes func(name string, atime time.Time, mtime time.Time) error

	// MockCreate implements Create
	MockCreate func(name string) (fsmodel.File, error)

	// MockDialUnix implements DialUnix
	MockDialUnix func(name string) (net.Conn, error)

	// MockListenUnix implements ListenUnix
	MockListenUnix func(name string) (net.Listener, error)

	// MockLstat implements Lstat
	MockLstat func(name string) (fs.FileInfo, error)

	// MockMkdir implements Mkdir
	MockMkdir func(name string, perm fs.FileMode) error

	// MockMkdirAll implements MkdirAll
	MockMkdirAll func(path string, perm fs.FileMode) error

	// MockOpen implements Open
	MockOpen func(name string) (fsmodel.File, error)

	// MockOpenFile implements OpenFile
	MockOpenFile func(name string, flag int, perm fs.FileMode) (fsmodel.File, error)

	// MockReadDir implements ReadDir
	MockReadDir func(dirname string) ([]fs.DirEntry, error)

	// MockRemove implements Remove
	MockRemove func(name string) error

	// MockRemoveAll implements RemoveAll
	MockRemoveAll func(path string) error

	// MockRename implements Rename
	MockRename func(oldname, newname string) error

	// MockStat implements Stat
	MockStat func(name string) (fs.FileInfo, error)
}

// Ensure FS implements fsmodel.FS
var _ fsmodel.FS = &FsmodelFS{}

// Chmod calls MockChmod
func (m *FsmodelFS) Chmod(name string, mode fs.FileMode) error {
	return m.MockChmod(name, mode)
}

// Chown calls MockChown
func (m *FsmodelFS) Chown(name string, uid, gid int) error {
	return m.MockChown(name, uid, gid)
}

// Chtimes calls MockChtimes
func (m *FsmodelFS) Chtimes(name string, atime, mtime time.Time) error {
	return m.MockChtimes(name, atime, mtime)
}

// Create calls MockCreate
func (m *FsmodelFS) Create(name string) (fsmodel.File, error) {
	return m.MockCreate(name)
}

// DialUnix calls MockDialUnix
func (m *FsmodelFS) DialUnix(name string) (net.Conn, error) {
	return m.MockDialUnix(name)
}

// ListenUnix calls MockListenUnix
func (m *FsmodelFS) ListenUnix(name string) (net.Listener, error) {
	return m.MockListenUnix(name)
}

// Lstat calls MockLstat
func (m *FsmodelFS) Lstat(name string) (fs.FileInfo, error) {
	return m.MockLstat(name)
}

// Mkdir calls MockMkdir
func (m *FsmodelFS) Mkdir(name string, perm fs.FileMode) error {
	return m.MockMkdir(name, perm)
}

// MkdirAll calls MockMkdirAll
func (m *FsmodelFS) MkdirAll(path string, perm fs.FileMode) error {
	return m.MockMkdirAll(path, perm)
}

// Open calls MockOpen
func (m *FsmodelFS) Open(name string) (fsmodel.File, error) {
	return m.MockOpen(name)
}

// OpenFile calls MockOpenFile
func (m *FsmodelFS) OpenFile(name string, flag int, perm fs.FileMode) (fsmodel.File, error) {
	return m.MockOpenFile(name, flag, perm)
}

// ReadDir calls MockReadDir
func (m *FsmodelFS) ReadDir(dirname string) ([]fs.DirEntry, error) {
	return m.MockReadDir(dirname)
}

// Remove calls MockRemove
func (m *FsmodelFS) Remove(name string) error {
	return m.MockRemove(name)
}

// RemoveAll calls MockRemoveAll
func (m *FsmodelFS) RemoveAll(path string) error {
	return m.MockRemoveAll(path)
}

// Rename calls MockRename
func (m *FsmodelFS) Rename(oldname, newname string) error {
	return m.MockRename(oldname, newname)
}

// Stat calls MockStat
func (m *FsmodelFS) Stat(name string) (fs.FileInfo, error) {
	return m.MockStat(name)
}

// FsmodelFile implements [fsmodel.File] for testing
type FsmodelFile struct {
	// MockRead implements Read
	MockRead func(b []byte) (int, error)

	// MockWrite implements Write
	MockWrite func(b []byte) (int, error)

	// MockClose implements Close
	MockClose func() error
}

// Ensure FsmodelFile implements [fsmodel.File].
var _ fsmodel.File = &FsmodelFile{}

// Read calls MockRead
func (m *FsmodelFile) Read(b []byte) (int, error) {
	return m.MockRead(b)
}

// Write calls MockWrite
func (m *FsmodelFile) Write(b []byte) (int, error) {
	return m.MockWrite(b)
}

// Close calls MockClose
func (m *FsmodelFile) Close() error {
	return m.MockClose()
}
