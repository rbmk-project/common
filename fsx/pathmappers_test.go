// SPDX-License-Identifier: Apache-2.0

package fsx

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestPathMappers(t *testing.T) {
	curdir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	errMocked := errors.New("mocked error")

	type testCase struct {
		name      string
		construct func(baseDir string) (RealPathMapper, error)
		baseDir   string
		path      string
		mockAbs   func(string) (string, error)
		want      string
		wantError error
	}

	tests := []struct {
		group string
		cases []testCase
	}{
		{
			group: "ChdirPathMapper",
			cases: []testCase{
				{
					name: "absolute mapper with relative path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewAbsoluteChdirPathMapper(baseDir)
					},
					baseDir: "testdata",
					path:    "file.txt",
					want:    filepath.Join(curdir, "testdata", "file.txt"),
				},

				{
					name: "absolute mapper with error",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewAbsoluteChdirPathMapper(baseDir)
					},
					baseDir: "testdata",
					mockAbs: func(path string) (string, error) {
						return "", errMocked
					},
					wantError: errMocked,
				},

				{
					name: "relative mapper with relative path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewRelativeChdirPathMapper(baseDir), nil
					},
					baseDir: "testdata",
					path:    "file.txt",
					want:    filepath.Join("testdata", "file.txt"),
				},
			},
		},

		{
			group: "ContainedDirPathMapper",
			cases: []testCase{
				{
					name: "absolute mapper with relative path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewAbsoluteContainedDirPathMapper(baseDir)
					},
					baseDir: "testdata",
					path:    "file.txt",
					want:    filepath.Join(curdir, "testdata", "file.txt"),
				},

				{
					name: "absolute mapper with absolute path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewAbsoluteContainedDirPathMapper(baseDir)
					},
					baseDir:   "testdata",
					path:      "/file.txt",
					wantError: fs.ErrNotExist,
				},

				{
					name: "absolute mapper with outside path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewAbsoluteContainedDirPathMapper(baseDir)
					},
					baseDir:   "testdata",
					path:      "../file.txt",
					wantError: fs.ErrNotExist,
				},

				{
					name: "absolute mapper with error",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewAbsoluteContainedDirPathMapper(baseDir)
					},
					baseDir: "testdata",
					mockAbs: func(path string) (string, error) {
						return "", errMocked
					},
					wantError: errMocked,
				},

				{
					name: "relative mapper with relative path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewRelativeContainedDirPathMapper(baseDir), nil
					},
					baseDir: "testdata",
					path:    "file.txt",
					want:    filepath.Join("testdata", "file.txt"),
				},

				{
					name: "relative mapper with absolute path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewRelativeContainedDirPathMapper(baseDir), nil
					},
					baseDir:   "testdata",
					path:      "/file.txt",
					wantError: fs.ErrNotExist,
				},

				{
					name: "relative mapper with outside path",
					construct: func(baseDir string) (RealPathMapper, error) {
						return NewRelativeContainedDirPathMapper(baseDir), nil
					},
					baseDir:   "testdata",
					path:      "../file.txt",
					wantError: fs.ErrNotExist,
				},
			},
		},
	}

	for _, group := range tests {
		t.Run(group.group, func(t *testing.T) {
			for _, tt := range group.cases {
				t.Run(tt.name, func(t *testing.T) {
					if tt.mockAbs != nil {
						saved := filepathAbs
						filepathAbs = tt.mockAbs
						defer func() { filepathAbs = saved }()
					}

					pmap, err := tt.construct(tt.baseDir)
					if err != nil {
						if !errors.Is(err, tt.wantError) {
							t.Fatalf("unexpected construction error: got %v, want %v", err, tt.wantError)
						}
						return
					}

					got, err := pmap.RealPath(tt.path)
					if !errors.Is(err, tt.wantError) {
						t.Fatalf("unexpected error: got %v, want %v", err, tt.wantError)
					}
					if err == nil && got != tt.want {
						t.Fatalf("got %q, want %q", got, tt.want)
					}
				})
			}
		})
	}
}
