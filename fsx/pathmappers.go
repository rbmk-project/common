//
// SPDX-License-Identifier: Apache-2.0
//
// Adapted from: https://github.com/spf13/afero
//

package fsx

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// RealPathMapper maps a virtual file name to its real path name.
type RealPathMapper interface {
	RealPath(virtualPath string) (realPath string, err error)
}

// RealPathMapperFunc is a function type that implements [RealPathMapper].
type RealPathMapperFunc func(virtualPath string) (realPath string, err error)

// Ensure [RealPathMapperFunc] implements [RealPathMapper].
var _ RealPathMapper = RealPathMapperFunc(nil)

// RealPath implements [RealPathMapper].
func (fx RealPathMapperFunc) RealPath(virtualPath string) (realPath string, err error) {
	return fx(virtualPath)
}

// Mockable [filepath.Abs] function for testing.
var filepathAbs = filepath.Abs

// ChdirPathMapper is a [RealPathMapper] that prepends
// a base directory to the virtual path.
//
// The zero value is invalid. Use [NewRelativeChdirPathMapper] or
// [NewAbsoluteChdirPathMapper] to construct a new instance.
type ChdirPathMapper struct {
	// baseDir is the base directory to prepend.
	baseDir string
}

// NewAbsoluteChdirPathMapper converts the given directory
// to an absolute path and, on success, returns a new
// [*ChdirPathMapper] instance. On failure, it returns and error.
//
// # Usage Considerations
//
// Use this constructor when you want your [*ChdirPathMapper] to
// be robust against concurrent invocations of [os.Chdir].
func NewAbsoluteChdirPathMapper(baseDir string) (*ChdirPathMapper, error) {
	absBaseDir, err := filepathAbs(baseDir)
	if err != nil {
		return nil, err
	}
	return &ChdirPathMapper{baseDir: absBaseDir}, nil
}

// NewRelativeChdirPathMapper returns a new [*ChdirPathMapper]
// instance without bothering to check if the given directory
// is relative or absolute.
//
// # Usage Considerations
//
// Use this constructor when you know your program is not going
// to invoke [os.Chdir] so you can avoid building potentially long
// paths that could break Unix domain sockets as documented in
// the top-level package documentation.
func NewRelativeChdirPathMapper(baseDir string) *ChdirPathMapper {
	return &ChdirPathMapper{baseDir: baseDir}
}

// Ensure [ChdirPathMapper] implements [RealPathMapper].
var _ RealPathMapper = &ChdirPathMapper{}

// RealPath implements [RealPathMapper].
func (b *ChdirPathMapper) RealPath(virtualPath string) (realPath string, err error) {
	return filepath.Join(b.baseDir, virtualPath), nil
}

// ContainedDirPathMapper is a [RealPathMapper] that prevents
// accessing file names outside of a given base directory.
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
//
// The zero value is invalid. Use [NewRelativeContainedDirPathMapper] or
// [NewAbsoluteContainedDirPathMapper] to construct a new instance.
type ContainedDirPathMapper struct {
	// baseDir is the base directory to contain.
	baseDir string
}

// NewAbsoluteContainedDirPathMapper converts the given directory
// to an absolute path and, on success, returns a new [*ContainedDirPathMapper]
// instance. On failure, it returns and error.
//
// # Usage Considerations
//
// Use this constructor when you want your [*ContainedDirPathMapper] to
// be robust against concurrent invocations of [os.Chdir].
func NewAbsoluteContainedDirPathMapper(baseDir string) (*ContainedDirPathMapper, error) {
	absBaseDir, err := filepathAbs(baseDir)
	if err != nil {
		return nil, err
	}
	return &ContainedDirPathMapper{baseDir: absBaseDir}, nil
}

// NewRelativeContainedDirPathMapper returns a new [*ContainedDirPathMapper]
// instance without bothering to check if the given directory
// is relative or absolute.
//
// # Usage Considerations
//
// Use this constructor when you know your program is not going
// to invoke [os.Chdir] so you can avoid building potentially long
// paths that could break Unix domain sockets as documented in
// the top-level package documentation.
func NewRelativeContainedDirPathMapper(baseDir string) *ContainedDirPathMapper {
	return &ContainedDirPathMapper{baseDir: baseDir}
}

// Ensure [ContainedDirPathMapper] implements [RealPathMapper].
var _ RealPathMapper = &ContainedDirPathMapper{}

// RealPath implements [RealPathMapper].
func (c *ContainedDirPathMapper) RealPath(virtualPath string) (realPath string, err error) {
	// 1. entirely reject absolute path names
	if filepath.IsAbs(virtualPath) {
		return "", fs.ErrNotExist
	}

	// 2. clean the path and make sure it is not outside the base path
	bpath := filepath.Clean(c.baseDir)
	fullpath := filepath.Clean(filepath.Join(bpath, virtualPath))
	if !strings.HasPrefix(fullpath, bpath) {
		return "", fs.ErrNotExist
	}
	return fullpath, nil
}
