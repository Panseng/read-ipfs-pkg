go-blockservice
==================

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![Coverage Status](https://codecov.io/gh/ipfs/go-block-format/branch/master/graph/badge.svg)](https://codecov.io/gh/ipfs/go-block-format/branch/master)
[![Build Status](https://circleci.com/gh/ipfs/go-blockservice.svg?style=svg)](https://circleci.com/gh/ipfs/go-blockservice)

> go-blockservice 为本地存储 & 远程存储 提供无缝接口。
> > go-blockservice provides a seamless interface to both local and remote storage backends.

## Lead Maintainer

[Steven Allen](https://github.com/Stebalien)

## Table of Contents

- [TODO](#todo)
- [Contribute](#contribute)
- [License](#license)

## TODO

这里的接口可能会与 `blockstore` 接口合并。目前 `dagservice` 构造函数需要一个 `blockervice` ，但如果它能直接接受一个 `blockstore` ，并让这个包实现一个 `blockervice` ，那就真的太好了。
> The interfaces here really would like to be merged with the blockstore interfaces.
The 'dagservice' constructor currently takes a blockservice, but it would be really nice
if it could just take a blockstore, and have this package implement a blockstore.

## Contribute

PRs are welcome!

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Juan Batiz-Benet
