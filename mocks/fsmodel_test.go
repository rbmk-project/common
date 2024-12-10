package mocks_test

import (
	"errors"
	"io/fs"
	"net"
	"testing"
	"time"

	"github.com/rbmk-project/common/internal/fsmodel"
	"github.com/rbmk-project/common/mocks"
)

func TestFsmodelFS(t *testing.T) {
	t.Run("Chmod", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockChmod: func(name string, mode fs.FileMode) error {
				return expected
			},
		}
		err := fs.Chmod("test.txt", 0644)
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Chown", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockChown: func(name string, uid, gid int) error {
				return expected
			},
		}
		err := fs.Chown("test.txt", 1000, 1000)
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Chtimes", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockChtimes: func(name string, atime, mtime time.Time) error {
				return expected
			},
		}
		err := fs.Chtimes("test.txt", time.Now(), time.Now())
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Create", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockCreate: func(name string) (fsmodel.File, error) {
				return nil, expected
			},
		}
		_, err := fs.Create("test.txt")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("DialUnix", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockDialUnix: func(name string) (net.Conn, error) {
				return nil, expected
			},
		}
		_, err := fs.DialUnix("test.sock")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("ListenUnix", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockListenUnix: func(name string) (net.Listener, error) {
				return nil, expected
			},
		}
		_, err := fs.ListenUnix("test.sock")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Lstat", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockLstat: func(name string) (fs.FileInfo, error) {
				return nil, expected
			},
		}
		_, err := fs.Lstat("test.txt")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Mkdir", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockMkdir: func(name string, perm fs.FileMode) error {
				return expected
			},
		}
		err := fs.Mkdir("testdir", 0755)
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("MkdirAll", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockMkdirAll: func(path string, perm fs.FileMode) error {
				return expected
			},
		}
		err := fs.MkdirAll("testdir/subdir", 0755)
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Open", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockOpen: func(name string) (fsmodel.File, error) {
				return nil, expected
			},
		}
		_, err := fs.Open("test.txt")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("OpenFile", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockOpenFile: func(name string, flag int, perm fs.FileMode) (fsmodel.File, error) {
				return nil, expected
			},
		}
		_, err := fs.OpenFile("test.txt", fsmodel.O_CREATE|fsmodel.O_WRONLY, 0644)
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("ReadDir", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockReadDir: func(dirname string) ([]fs.DirEntry, error) {
				return nil, expected
			},
		}
		_, err := fs.ReadDir("testdir")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockRemove: func(name string) error {
				return expected
			},
		}
		err := fs.Remove("test.txt")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("RemoveAll", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockRemoveAll: func(path string) error {
				return expected
			},
		}
		err := fs.RemoveAll("testdir")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Rename", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockRename: func(oldname, newname string) error {
				return expected
			},
		}
		err := fs.Rename("old.txt", "new.txt")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})

	t.Run("Stat", func(t *testing.T) {
		expected := errors.New("mocked error")
		fs := &mocks.FsmodelFS{
			MockStat: func(name string) (fs.FileInfo, error) {
				return nil, expected
			},
		}
		_, err := fs.Stat("test.txt")
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})
}

func TestFsmodelFile(t *testing.T) {
	t.Run("Read", func(t *testing.T) {
		expected := errors.New("mocked error")
		file := &mocks.FsmodelFile{
			MockRead: func(b []byte) (int, error) {
				return 0, expected
			},
		}
		count, err := file.Read(make([]byte, 128))
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
		if count != 0 {
			t.Fatal("expected 0 bytes")
		}
	})

	t.Run("Write", func(t *testing.T) {
		expected := errors.New("mocked error")
		file := &mocks.FsmodelFile{
			MockWrite: func(b []byte) (int, error) {
				return 0, expected
			},
		}
		count, err := file.Write(make([]byte, 128))
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
		if count != 0 {
			t.Fatal("expected 0 bytes")
		}
	})

	t.Run("Close", func(t *testing.T) {
		expected := errors.New("mocked error")
		file := &mocks.FsmodelFile{
			MockClose: func() error {
				return expected
			},
		}
		err := file.Close()
		if !errors.Is(err, expected) {
			t.Fatal("not the error we expected")
		}
	})
}
