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
	"time"
)

// NewChdirFS creates a new [FS] where each file name is
// prefixed with the given directory path.
func NewChdirFS(dep FS, path string) *ChdirFS {
	return &ChdirFS{basepath: path, dep: dep}
}

// ChdirFS is the [FS] type returned by [NewChdirFS].
//
// The zero value IS NOT ready to use; construct using [NewChdirFS].
type ChdirFS struct {
	// basepath is the base path.
	basepath string

	// dep is the dependency [FS].
	dep FS
}

// Ensure [basePathFS] implements [FS].
var _ FS = &ChdirFS{}

// realPath returns the real path of a given file name or an error.
func (rfs *ChdirFS) realPath(name string) string {
	return filepath.Join(rfs.basepath, name)
}

// Chmod implements [FS].
func (rfs *ChdirFS) Chmod(name string, mode fs.FileMode) error {
	return rfs.dep.Chmod(rfs.realPath(name), mode)
}

// Chown implements [FS].
func (rfs *ChdirFS) Chown(name string, uid, gid int) error {
	return rfs.dep.Chown(rfs.realPath(name), uid, gid)
}

// Chtimes implements [FS].
func (rfs *ChdirFS) Chtimes(name string, atime, mtime time.Time) error {
	return rfs.dep.Chtimes(rfs.realPath(name), atime, mtime)
}

// Create implements [FS].
func (rfs *ChdirFS) Create(name string) (File, error) {
	return rfs.dep.Create(rfs.realPath(name))
}

// DialUnix implements [FS].
//
// See also the limitations documented in the top-level package docs.
func (rfs *ChdirFS) DialUnix(name string) (net.Conn, error) {
	return rfs.dep.DialUnix(rfs.realPath(name))
}

// ListenUnix implements [FS].
//
// See also the limitations documented in the top-level package docs.
func (rfs *ChdirFS) ListenUnix(name string) (net.Listener, error) {
	return rfs.dep.ListenUnix(rfs.realPath(name))
}

// Lstat implements [FS].
func (rfs *ChdirFS) Lstat(name string) (fs.FileInfo, error) {
	return rfs.dep.Lstat(rfs.realPath(name))
}

// Mkdir implements [FS].
func (rfs *ChdirFS) Mkdir(name string, mode fs.FileMode) error {
	return rfs.dep.Mkdir(rfs.realPath(name), mode)
}

// MkdirAll implements [FS].
func (rfs *ChdirFS) MkdirAll(name string, mode fs.FileMode) error {
	return rfs.dep.MkdirAll(rfs.realPath(name), mode)
}

// Open implements [FS].
func (rfs *ChdirFS) Open(name string) (File, error) {
	return rfs.dep.Open(rfs.realPath(name))
}

// OpenFile implements [FS].
func (rfs *ChdirFS) OpenFile(name string, flag int, mode fs.FileMode) (File, error) {
	return rfs.dep.OpenFile(rfs.realPath(name), flag, mode)
}

// ReadDir implements [FS].
func (rfs *ChdirFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return rfs.dep.ReadDir(rfs.realPath(name))
}

// Remove implements [FS].
func (rfs *ChdirFS) Remove(name string) error {
	return rfs.dep.Remove(rfs.realPath(name))
}

// RemoveAll implements [FS].
func (rfs *ChdirFS) RemoveAll(name string) error {
	return rfs.dep.RemoveAll(rfs.realPath(name))
}

// Rename implements [FS].
func (rfs *ChdirFS) Rename(oldname, newname string) error {
	return rfs.dep.Rename(rfs.realPath(oldname), rfs.realPath(newname))
}

// Stat implements [FS].
func (rfs *ChdirFS) Stat(name string) (fs.FileInfo, error) {
	return rfs.dep.Stat(rfs.realPath(name))
}
