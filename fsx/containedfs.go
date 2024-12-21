//
// SPDX-License-Identifier: Apache-2.0
//
// Adapted from: https://github.com/spf13/afero
//

package fsx

import (
	"io/fs"
	"net"
	"path/filepath"
	"strings"
	"time"
)

// NewContainedFS creates a new [FS] rooted at the given path
// using the given child [FS] as the dependency.
//
// Any file name (after [filepath.Clean]) outside this base
// path will be treated as a non-existing file.
//
// Any absolute file name will be treated as a non-existing file.
//
// We return [fs.ErrNotExist] in these cases.
//
// Note: This implementation cannot prevent symlink traversal
// attacks. The caller must ensure the base directory does not
// contain symlinks if this is a security requirement.
func NewContainedFS(dep FS, path string) *ContainedFS {
	return &ContainedFS{basepath: path, dep: dep}
}

// ContainedFS is the [FS] type returned by [NewContainedFS].
//
// The zero value IS NOT ready to use; construct using [NewContainedFS].
type ContainedFS struct {
	// basepath is the base path.
	basepath string

	// dep is the dependency [FS].
	dep FS
}

// Ensure [basePathFS] implements [FS].
var _ FS = &ContainedFS{}

// realPath returns the real path of a given file name or an error.
func (rfs *ContainedFS) realPath(name string) (string, error) {
	// 1. entirely reject absolute path names
	if filepath.IsAbs(name) {
		return "", fs.ErrNotExist
	}

	// 2. clean the path and make sure it is not outside the base path
	bpath := filepath.Clean(rfs.basepath)
	fullpath := filepath.Clean(filepath.Join(bpath, name))
	if !strings.HasPrefix(fullpath, bpath) {
		return name, fs.ErrNotExist
	}
	return fullpath, nil
}

// Chmod implements [FS].
func (rfs *ContainedFS) Chmod(name string, mode fs.FileMode) error {
	name, err := rfs.realPath(name)
	if err != nil {
		return &fs.PathError{Op: "chmod", Path: name, Err: err}
	}
	return rfs.dep.Chmod(name, mode)
}

// Chown implements [FS].
func (rfs *ContainedFS) Chown(name string, uid, gid int) error {
	name, err := rfs.realPath(name)
	if err != nil {
		return &fs.PathError{Op: "chown", Path: name, Err: err}
	}
	return rfs.dep.Chown(name, uid, gid)
}

// Chtimes implements [FS].
func (rfs *ContainedFS) Chtimes(name string, atime, mtime time.Time) error {
	name, err := rfs.realPath(name)
	if err != nil {
		return &fs.PathError{Op: "chtimes", Path: name, Err: err}
	}
	return rfs.dep.Chtimes(name, atime, mtime)
}

// Create implements [FS].
func (rfs *ContainedFS) Create(name string) (File, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "create", Path: name, Err: err}
	}
	return rfs.dep.Create(name)
}

// DialUnix implements [FS].
func (rfs *ContainedFS) DialUnix(name string) (net.Conn, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "dialunix", Path: name, Err: err}
	}
	return rfs.dep.DialUnix(name)
}

// ListenUnix implements [FS].
func (rfs *ContainedFS) ListenUnix(name string) (net.Listener, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "listenunix", Path: name, Err: err}
	}
	return rfs.dep.ListenUnix(name)
}

// Lstat implements [FS].
func (rfs *ContainedFS) Lstat(name string) (fs.FileInfo, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "lstat", Path: name, Err: err}
	}
	return rfs.dep.Lstat(name)
}

// Mkdir implements [FS].
func (rfs *ContainedFS) Mkdir(name string, mode fs.FileMode) error {
	name, err := rfs.realPath(name)
	if err != nil {
		return &fs.PathError{Op: "mkdir", Path: name, Err: err}
	}
	return rfs.dep.Mkdir(name, mode)
}

// MkdirAll implements [FS].
func (rfs *ContainedFS) MkdirAll(name string, mode fs.FileMode) error {
	name, err := rfs.realPath(name)
	if err != nil {
		return &fs.PathError{Op: "mkdir", Path: name, Err: err}
	}
	return rfs.dep.MkdirAll(name, mode)
}

// Open implements [FS].
func (rfs *ContainedFS) Open(name string) (File, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "open", Path: name, Err: err}
	}
	return rfs.dep.Open(name)
}

// OpenFile implements [FS].
func (rfs *ContainedFS) OpenFile(name string, flag int, mode fs.FileMode) (File, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "openfile", Path: name, Err: err}
	}
	return rfs.dep.OpenFile(name, flag, mode)
}

// ReadDir implements [FS].
func (rfs *ContainedFS) ReadDir(name string) ([]fs.DirEntry, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: err}
	}
	return rfs.dep.ReadDir(name)
}

// Remove implements [FS].
func (rfs *ContainedFS) Remove(name string) error {
	name, err := rfs.realPath(name)
	if err != nil {
		return &fs.PathError{Op: "remove", Path: name, Err: err}
	}
	return rfs.dep.Remove(name)
}

// RemoveAll implements [FS].
func (rfs *ContainedFS) RemoveAll(name string) error {
	name, err := rfs.realPath(name)
	if err != nil {
		return &fs.PathError{Op: "removeall", Path: name, Err: err}
	}
	return rfs.dep.RemoveAll(name)
}

// Rename implements [FS].
func (rfs *ContainedFS) Rename(oldname, newname string) error {
	oldname, err := rfs.realPath(oldname)
	if err != nil {
		return &fs.PathError{Op: "rename", Path: oldname, Err: err}
	}
	newname, err = rfs.realPath(newname)
	if err != nil {
		return &fs.PathError{Op: "rename", Path: newname, Err: err}
	}
	return rfs.dep.Rename(oldname, newname)
}

// Stat implements [FS].
func (rfs *ContainedFS) Stat(name string) (fs.FileInfo, error) {
	name, err := rfs.realPath(name)
	if err != nil {
		return nil, &fs.PathError{Op: "stat", Path: name, Err: err}
	}
	return rfs.dep.Stat(name)
}
