go-bitswap
==================

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![Matrix](https://img.shields.io/badge/matrix-%23ipfs%3Amatrix.org-blue.svg?style=flat-square)](https://matrix.to/#/#ipfs:matrix.org)
[![IRC](https://img.shields.io/badge/freenode-%23ipfs-blue.svg?style=flat-square)](http://webchat.freenode.net/?channels=%23ipfs)
[![Discord](https://img.shields.io/discord/475789330380488707?color=blueviolet&label=discord&style=flat-square)](https://discord.gg/24fmuwR)
[![Coverage Status](https://codecov.io/gh/ipfs/go-bitswap/branch/master/graph/badge.svg)](https://codecov.io/gh/ipfs/go-bitswap/branch/master)
[![Build Status](https://circleci.com/gh/ipfs/go-bitswap.svg?style=svg)](https://circleci.com/gh/ipfs/go-bitswap)

> An implementation of the bitswap protocol in go!

## Lead Maintainer

[Dirk McCormick](https://github.com/dirkmc)

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Implementation](#implementation)
- [Contribute](#contribute)
- [License](#license)


## Background

Bitswap 是 ipfs 的数据交换模块。它管理互联网上 peers 的请求、发送。它有两个主要工作：
- 客户端从互联网获取 块 资源请求
- 理智地将自身拥有的 块 发送给需要的对等点
> Bitswap is the data trading module for ipfs. It manages requesting and sending
> blocks to and from other peers in the network. Bitswap has two main jobs:
> - to acquire blocks requested by the client from the network
> - to judiciously send blocks in its possession to other peers who want them

Bitswap 是基于消息的协议，与 请求-响应相反。消息包括：想要的列表 或 块。
> Bitswap is a message based protocol, as opposed to request-response. All messages
> contain wantlists or blocks.

节点发送 想要列表，告诉对等点，它所想要的内容。当一个节点收到 想要列表，它需要检查是否含有响应的 块，并考虑是否发送匹配的 块 给请求者。
> A node sends a wantlist to tell peers which blocks it wants. When a node receives
> a wantlist it should check which blocks it has from the wantlist, and consider
> sending the matching blocks to the requestor.

当节点收到它请求的 块 ，节点应向其他对等点发出**取消**的通知，告诉它们，它不再需要这个 块。
When a node receives blocks that it asked for, the node should send out a
notification called a 'Cancel' to tell its peers that the node no longer
wants those blocks.

`go-bitswap` provides an implementation of the Bitswap protocol in go.

[Learn more about how Bitswap works](./docs/how-bitswap-works.md)

## Install

`go-bitswap` requires Go >= 1.11 and can be installed using Go modules

## Usage

### Initializing a Bitswap Exchange

```golang
import (
  "context"
  bitswap "github.com/ipfs/go-bitswap"
  bsnet "github.com/ipfs/go-bitswap/network"
  blockstore "github.com/ipfs/go-ipfs-blockstore"
  "github.com/libp2p/go-libp2p-core/routing"
  "github.com/libp2p/go-libp2p-core/host"
)

var ctx context.Context
var host host.Host
var router routing.ContentRouting
var bstore blockstore.Blockstore

network := bsnet.NewFromIpfsHost(host, router)
exchange := bitswap.New(ctx, network, bstore)
```

Parameter Notes:

1. `ctx` is just the parent context for all of Bitswap
2. network 是在 libp2p & 内容路由 之上提供给 Bitswap 的网络抽象 
> `network` is a network abstraction provided to Bitswap on top of libp2p & content routing. 
3. `bstore` is an IPFS blockstore

### Get A Block Synchronously

```golang
var c cid.Cid
var ctx context.Context
var exchange bitswap.Bitswap

block, err := exchange.GetBlock(ctx, c)
```

Parameter Notes:

1. `ctx` is the context for this request, which can be cancelled to cancel the request
2. `c` is the content ID of the block you're requesting

### Get Several Blocks Asynchronously

```golang
var cids []cid.Cid
var ctx context.Context
var exchange bitswap.Bitswap

blockChannel, err := exchange.GetBlocks(ctx, cids)
```

Parameter Notes:

1. `ctx` is the context for this request, which can be cancelled to cancel the request
2. `cids` is a slice of content IDs for the blocks you're requesting

### Get Related Blocks Faster With Sessions

在 IPFS 中，内容块通常通过 MerkleDAG 相互连接。如果您提前知道块请求是相关的，Bitswap 可以在内部对请求这些块的方式进行一些优化，以便更快地获得它们。... 每当您打算发出一系列相关的块请求并且其响应可能来自相同的对等方时，您都应该初始化一个 Bitswap 会话。
 > In IPFS, content blocks are often connected to each other through a MerkleDAG. If you know ahead of time that block requests are related, Bitswap can make several optimizations internally in how it requests those blocks in order to get them faster. Bitswap provides a mechanism called a Bitswap Session to manage a series of block requests as part of a single higher level operation. You should initialize a Bitswap Session any time you intend to make a series of block requests that are related -- and whose responses are likely to come from the same peers.

```golang
var ctx context.Context
var cids []cids.cid
var exchange bitswap.Bitswap

session := exchange.NewSession(ctx)
blocksChannel, err := session.GetBlocks(ctx, cids)
// later
var relatedCids []cids.cid
relatedBlocksChannel, err := session.GetBlocks(ctx, relatedCids)
```
请注意，NewSession 返回一个带有 GetBlock 和 GetBlocks 方法的接口，这些方法与整个 Bitswap 交换具有相同的签名。
> Note that `NewSession` returns an interface with `GetBlock` and `GetBlocks` methods that have the same signature as the overall Bitswap exchange.

### Tell bitswap a new block was added to the local datastore

```golang
var blk blocks.Block
var exchange bitswap.Bitswap

err := exchange.HasBlock(blk)
```

## Contribute

PRs are welcome!

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Juan Batiz-Benet
