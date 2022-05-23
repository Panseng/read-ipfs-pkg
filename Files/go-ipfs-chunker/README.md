# go-ipfs-chunker

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/ipfs/go-ipfs-chunker?status.svg)](https://godoc.org/github.com/ipfs/go-ipfs-chunker)
[![Build Status](https://travis-ci.org/ipfs/go-ipfs-chunker.svg?branch=master)](https://travis-ci.org/ipfs/go-ipfs-chunker)

> go ipfs 的 数据拆分器。
> > go-ipfs-chunker implements data Splitters for go-ipfs.

`go-ipfs-chunker` 提供 `Splitter` 接口。IPFS 拆分从 reader 获取的数据，并创建 **块(chunks)** 。这些 **块** 被用于构建 **ipfs DAGs (Merkle Tree)**，也是 ipfs 用于内容寻址的总和的基本单元
> `go-ipfs-chunker` provides the `Splitter` interface. IPFS splitters read data from a reader an create "chunks". These chunks are used to build the ipfs DAGs (Merkle Tree) and are the base unit to obtain the sums that ipfs uses to address content.

在大部分场景默认使用 `SizeSplitter` 创建相同大小的块。 \
`rabin` 指纹分块器。 \
存在重复内容的场景下，分块器将尝试以生成**相同块**的方式拆分数据，从而优化生成的 DAGs 。
> The package provides a `SizeSplitter` which creates chunks of equal size and it is used by default in most cases, and a `rabin` fingerprint chunker. This chunker will attempt to split data in a way that the resulting blocks are the same when the data has repetitive patterns, thus optimizing the resulting DAGs. \
> `rabin` : 一种非对称的加密算法

## Lead Maintainer

[Steven Allen](https://github.com/Stebalien)

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-ipfs-chunker` works like a regular Go module:

```
> go get github.com/ipfs/go-ipfs-chunker
```

## Usage

```
import "github.com/ipfs/go-ipfs-chunker"
```

Check the [GoDoc documentation](https://godoc.org/github.com/ipfs/go-ipfs-chunker)

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Protocol Labs, Inc.
