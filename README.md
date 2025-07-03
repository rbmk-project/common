# General purpose libraries

[![GoDoc](https://pkg.go.dev/badge/github.com/rbmk-project/common)](https://pkg.go.dev/github.com/rbmk-project/common) [![Build Status](https://github.com/rbmk-project/common/actions/workflows/go.yml/badge.svg)](https://github.com/rbmk-project/common/actions) [![codecov](https://codecov.io/gh/rbmk-project/common/branch/main/graph/badge.svg)](https://codecov.io/gh/rbmk-project/common)

This repository contains several general purpose libraries
reused across different [rbmk](https://github.com/rbmk-project)
repositories.

## 📦 Archived

This repository has been merged into the main [`rbmk`](
https://github.com/rbmk-project/rbmk) repository with
[rbmk#79](https://github.com/rbmk-project/rbmk/pull/79).

Please, update the import paths as follows:

```
github.com/rbmk-project/common/FOO => github.com/rbmk-project/rbmk/pkg/common/FOO
```

This repository is archived for historical reference.

## Minimum Required Go Version

We maintain compatibility with the oldest supported version of Go as
documented by the [Go Release Policy](https://golang.org/doc/devel/release.html#policy)
and update our minimum required version of Go as needed.

## Installation

```sh
go get github.com/rbmk-project/common
```

## Design

See [DESIGN.md](DESIGN.md) for an overview of the design.

## Contributing

Contributions are welcome! Please submit a pull requests
using GitHub. Use [rbmk-project/issues](https://github.com/rbmk-project/issues)
to create issues and discuss features related to this package.

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```
