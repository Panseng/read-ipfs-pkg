# go-datastore

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![GoDoc](https://godoc.org/github.com/ipfs/go-datastore?status.svg)](https://godoc.org/github.com/ipfs/go-datastore)

> 键值对 数据存储接口
> > key-value datastore interfaces

## Lead Maintainer

[Steven Allen](https://github.com/Stebalien)

## Table of Contents

- [Background](#background)
- [Documentation](#documentation)
- [Contribute](#contribute)
- [License](#license)

## Background
Datastore 是 数据存储 & 数据访问 的通用抽象层。它是一个简单的 API，旨在以与**数据存储无关**的方式启用应用程序开发，允许在不更改应用程序代码的情况下**无缝交换数据存储**。因此，可以利用具有不同优势的不同数据存储，而不需要将应用程序在其整个生命周期中投入一个数据存储。
> Datastore is a generic layer of abstraction for data store and database access. It is a simple API with the aim to enable application development in a datastore-agnostic way, allowing datastores to be swapped seamlessly without changing application code. Thus, one can leverage different datastores with different strengths without committing the application to one datastore throughout its lifetime.

此外，分组的数据存储大大简化了对感兴趣数据的访问模式
> In addition, grouped datastores significantly simplify interesting data access patterns (such as caching and sharding).

Based on [datastore.py](https://github.com/datastore/datastore).

## Documentation

https://godoc.org/github.com/ipfs/go-datastore

## Contribute

Feel free to join in. All welcome. Open an [issue](https://github.com/ipfs/go-datastore/issues)!

This repository falls under the IPFS [Code of Conduct](https://github.com/ipfs/community/blob/master/code-of-conduct.md).

### Want to hack on IPFS?

[![](https://cdn.rawgit.com/jbenet/contribute-ipfs-gif/master/img/contribute.gif)](https://github.com/ipfs/community/blob/master/contributing.md)

## License

MIT

