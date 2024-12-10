//
// SPDX-License-Identifier: Apache-2.0
//
// Adapted from: https://github.com/spf13/afero
//

/*
Package fsx allows abstracting the file system.

This package is derived from [afero].

In addition to [afero], we also implement support for dialing
and listening unix domain sockets, and for Lstat.

The [NewRelativeFS] creates a filesystem relative to a given
directory that only allows:

1. Accessing files within the base directory.

2. Changing working directory to a subdirectory of the base directory.

[afero]: https://github.com/spf13/afero
*/
package fsx

import (
	"errors"
	"io/fs"
	"os"

	"github.com/rbmk-project/common/internal/fsmodel"
)

// Forward file system constants.
const (
	O_CREATE = fsmodel.O_CREATE
	O_RDONLY = fsmodel.O_RDONLY
	O_RDWR   = fsmodel.O_RDWR
	O_TRUNC  = fsmodel.O_TRUNC
	O_WRONLY = fsmodel.O_WRONLY
)

// IsNotExist combines the [os.ErrNotExist] check with
// checking for the [fs.ErrNotExist] error.
func IsNotExist(err error) bool {
	return errors.Is(err, fs.ErrNotExist) || os.IsNotExist(err)
}

// File is an alias for [fsmodel.File].
type File = fsmodel.File

// Ensure [*os.File] implements [File].
var _ File = &os.File{}

// FS is an alias for [fsmodel.FS].
type FS = fsmodel.FS
