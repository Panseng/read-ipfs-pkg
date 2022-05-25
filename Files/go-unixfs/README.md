go-unixfs
==================

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![Coverage Status](https://codecov.io/gh/ipfs/go-unixfs/branch/master/graph/badge.svg)](https://codecov.io/gh/ipfs/go-unixfs/branch/master)
[![Travis CI](https://travis-ci.org/ipfs/go-unixfs.svg?branch=master)](https://travis-ci.org/ipfs/go-unixfs)

> go-unixfs 是在 IPLD merkledag 之上的类 Unix 文件系统。
> > go-unixfs implements unix-like filesystem utilities on top of an ipld merkledag

## Lead Maintainer

[Steven Allen](https://github.com/Stebalien)

## Table of Contents

- [Directory](#directory)
- [Install](#install)
- [Contribute](#contribute)
- [License](#license)

## Package Directory
此包 包含多个子包，每个子包都可以是非常大的个体。
> This package contains many subpackages, each of which can be very large on its own.

### Top Level
顶层 unixfs 包定义数据结构 和 围绕它的辅助方法。
> The top level unixfs package defines the unixfs format datastructures, and some helper methods around it.

### importers
当您想将普通文件转换为 unixfs 文件时将使用 `importer` 子包。
> The `importer` subpackage is what you'll use when you want to turn a normal file into a unixfs file.

### io
`io` 提供 文件读取 & 目录操作。
> The `io` subpackage provides helpers for reading files and manipulating directories.

`DagReader` 将关联一个 unixfs file 文件，并返回 文件句柄（用于查找 & 读取 内容）。
> The `DagReader` takes a reference to a unixfs file and returns a file handle that can be read from and seeked through. \
> DAG: 有向无环图

`Directory` 类，简化目录的 内容读取、添入、查找。
> The `Directory` interface allows you to easily read items in a directory, add items to a directory, and do lookups.

### mod
`mod` 是 `DagModifier` 类的实现，用于 写入已存在 或 创建新 的 unixfs file。它的逻辑明显要比 `DagReader` 复杂，因此剥离。
> The `mod` subpackage implements a `DagModifier` type that can be used to write to an existing unixfs file, or
create a new one. The logic for this is significantly more complicated than for the dagreader, so its a separate
type. (TODO: maybe it still belongs in the `io` subpackage though?)

### hamt
`hamt` 是 CHAMP hamt 的实现，用于 unixfs directory 切片。
> The `hamt` subpackage implements a CHAMP hamt that is used in unixfs directory sharding.

### archive
`archive` 是 `tar` 压缩 & 解压 的实现。目前还不是正式的 unixfs 子包，但未来可能会合并。
> The `archive` subpackage implements a `tar` importer and exporter. The objects created here are not officially unixfs,
but in the future, this may be integrated more directly.

### test
The `test` subpackage provides several utilities to make testing unixfs related things easier.

## Install

```sh
go get github.com/ipfs/go-unixfs
```

## Contribute

PRs are welcome!

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Juan Batiz-Benet
