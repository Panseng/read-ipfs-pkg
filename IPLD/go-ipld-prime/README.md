go-ipld-prime
=============

IPLD 数据模型的 Golang 接口，包括核心编解码器、IPLD 模式支持，及一些方便的功能转换工具。

`go-ipld-prime` 是 IPLD 规范接口的实现，是用于 CBOR 和 JSON 数据格式的 IPLD 编解码器实现，用于 IPLD 对象基本操作的工具
- CBOR (Concise Binary Object Representation) : 简洁的二进制对象表示
> `go-ipld-prime` is an implementation of the IPLD spec interfaces,
a batteries-included codec implementations of IPLD for CBOR and JSON,
and tooling for basic operations on IPLD objects (traversals, etc).



API
---

根据代码功能，`API` 分为多个模块。最核心的接口是基础包，你需要导入更多的包，以便将实现具体的功能。
> The API is split into several packages based on responsibly of the code.
> The most central interfaces are the base package,
> but you'll certainly need to import additional packages to get concrete implementations into action.

粗略地说，核心包的接口都是关于IPLD数据模型的；`codec/*` 包含了将序列化数据 与 IPLD数据模型 互转的方法；`traversal` 是IPLD 数据模型的高阶函数；具体的 `ipld.Node` 实现可以在 `node/*` 目录下的包中找到；几个额外的软件包 包含高级功能，如IPLD模式。
> Roughly speaking, the core package interfaces are all about the IPLD Data Model;
> the `codec/*` packages contain functions for parsing serial data into the IPLD Data Model,
> and converting Data Model content back into serial formats;
> the `traversal` package is an example of higher-order functions on the Data Model;
> concrete `ipld.Node` implementations ready to use can be found in packages in the `node/*` directory;
> and several additional packages contain advanced features such as IPLD Schemas.

(Because the codecs, as well as higher-order features like traversals, are
implemented in a separate package from the core interfaces or any of the Node implementations,
you can be sure they're not doing any funky "magic" -- all this stuff will work the same
if you want to write your own extensions, whether for new Node implementations
or new codecs, or new higher-order order functions!)

- `github.com/ipld/go-ipld-prime` -- 仅作为 `ipld` 导入，包含了 IPLD 的核心接口。最重要的接口是 `Node`, `NodeBuilder`, `Path`, and `Link`。
> imported as just `ipld` -- contains the core interfaces for IPLD.  The most important interfaces are `Node`, `NodeBuilder`, `Path`, and `Link`.

- `github.com/ipld/go-ipld-prime/node/basicnode` -- `Node` and `NodeBuilder` 的具体实现，使用非结构化内存处理任何类型的数据。
> -- provides concrete implementations of `Node` and `NodeBuilder` which work for any kind of data, using unstructured memory.

- `github.com/ipld/go-ipld-prime/node/bindnode` -- `Node` and `NodeBuilder` 的具体实现，用 golang 结构体 存储数据，通过反射提供接口。也支持 IPLD 模式。
> -- provides concrete implementations of `Node` and `NodeBuilder` which store data in native golang structures, interacting with it via reflection.  Also supports IPLD Schemas!

- `github.com/ipld/go-ipld-prime/traversal` -- 包含高阶函数，用于轻松遍历数据图。
> -- contains higher-order functions for traversing graphs of data easily.

- `github.com/ipld/go-ipld-prime/traversal/selector` -- 包含选择器，有点像正则表达式，但用于 IPLD 数据的树和图！
> -- contains selectors, which are sort of like regexps, but for trees and graphs of IPLD data!

- `github.com/ipld/go-ipld-prime/codec` -- parent package of all the codec implementations!

- `github.com/ipld/go-ipld-prime/codec/dagcbor` -- CBOR ( 高效的二进制格式 ) 编码 & 解码实现
> -- implementations of marshalling and unmarshalling as CBOR (a fast, binary serialization format).

- `github.com/ipld/go-ipld-prime/codec/dagjson` -- JSON 编码 & 解码实现
> -- implementations of marshalling and unmarshalling as JSON (a popular human readable format).

- `github.com/ipld/go-ipld-prime/linking/cid` -- CID `Link` 的实现
> -- imported as `cidlink` -- provides concrete implementations of `Link` as a CID.  Also, the multicodec registry.

- `github.com/ipld/go-ipld-prime/schema` -- `schema.Type` and `schema.TypedNode` 接口声明，作为 IPLD 模式类型信息。
> -- contains the `schema.Type` and `schema.TypedNode` interface declarations, which represent IPLD Schema type information.

- `github.com/ipld/go-ipld-prime/node/typed` -- schema.TypedNode的具体实现，在运行时对基本节点(`Node`)进行装饰，使其具有IPLD模式描述的额外功能。
> -- provides concrete implementations of `schema.TypedNode` which decorate a basic `Node` at runtime to have additional features described by IPLD Schemas.


Getting Started
---------------

假设您想以编程方式创建一些数据，然后对其进行序列化，或者将其保存为 `blocks`。
> Let's say you want to create some data programmatically,
and then serialize it, or save it as [blocks].

您有很多不同的选项，具体取决于您要使用的 golang 约定
> You've got a ton of different options, depending on what golang convention you want to use:

- the `qp` package -- [example](https://pkg.go.dev/github.com/ipld/go-ipld-prime/fluent/qp#example-package)
- the `bindnode` system, if you want to use golang types -- [example](https://pkg.go.dev/github.com/ipld/go-ipld-prime/node/bindnode#example-Wrap-NoSchema), [example with schema](https://pkg.go.dev/github.com/ipld/go-ipld-prime/node/bindnode#example-Wrap-WithSchema)
- or the [`NodeBuilder`](https://pkg.go.dev/github.com/ipld/go-ipld-prime/datamodel#NodeBuilder) interfaces, raw (verbose; not recommended)
- or even some codegen systems!

Once you've got a Node full of data,
you can serialize it:

https://pkg.go.dev/github.com/ipld/go-ipld-prime#example-package-CreateDataAndMarshal

但你可能想做的不止这些；你可能想把这些数据存储为一个块，并得到一个链接到它的CID。你可以用 `LinkSystem`
> But probably you want to do more than that;
probably you want to store this data as a block,
and get a CID that links back to it.
For this you use `LinkSystem`:

https://pkg.go.dev/github.com/ipld/go-ipld-prime/linking#example-LinkSystem.Store

希望这些案例对你有用。 API 文档应该有助于你。我们也强烈建议详读 [godocs](https://pkg.go.dev/github.com/ipld/go-ipld-prime)  中的其他范例代码，在不同的包中都有!
Hopefully these pointers give you some useful getting-started focal points.
The API docs should help from here on out.
We also highly recommend scanning the  for other pieces of example code, in various packages!

Let us know in [issues](https://github.com/ipld/go-ipld-prime/issues), [chat, or other community spaces](https://ipld.io/docs/intro/community/) if you need more help,
or have suggestions on how we can improve the getting-started experiences!



Other IPLD Libraries
--------------------

The IPLD specifications are designed to be language-agnostic.
Many implementations exist in a variety of languages.

For overall behaviors and specifications, refer to the IPLD website, or its source, in IPLD meta repo:
- https://ipld.io/
- https://github.com/ipld/ipld/
You should find specs in the `specs/` dir there,
human-friendly docs in the `docs/` dir,
and information about _why_ things are designed the way they are mostly in the `design/` directories.

There are also pages in the IPLD website specifically about golang IPLD libraries,
and your alternatives: https://ipld.io/libraries/golang/


### distinctions from go-ipld-interface&go-ipld-cbor

This library ("go ipld prime") is the current head of development for golang IPLD,
and we recommend new developments in golang be done using this library as the basis.

However, several other libraries exist in golang for working with IPLD data.
Most of these predate go-ipld-prime and no longer receive active development,
but since they do support a lot of other software, you may continue to seem them around for a while.
go-ipld-prime is generally **serially compatible** with these -- just like it is with IPLD libraries in other languages.

在编程API和功能方面，go-ipld-prime对IPLD的接口进行了简洁的处理，并选择了与老一代库非常不同的几个设计策略。
> In terms of programmatic API and features, go-ipld-prime is a clean take on the IPLD interfaces,
> and chose to address several design decisions very differently than older generation of libraries:

- **节点接口与IPLD数据模型的映射很清晰**;
- 许多已知的遗留功能被放弃了;
- `Link` 是 CIDs 的纯粹实现（没有了 name 、size 属性）
- 在同一个包中实现了 Path;
- 在同一个包中实现了 JSON & CBOR;
- 一些依赖 `blockstore` 及其他与 IPFS 紧密耦合的接口，被更简单、更少耦合的接口所取代;
- IPLD 选择器等**新功能** 仅可从 go-ipld-prime 获得;
- 像ADL（高级数据布局）这样的新功能，提供了透明分片和大数据索引等功能，仅可从 go-ipld-prime 获得;
- 可以使用 go-ipld-prime 将声明性转换应用于 IPLD 数据（根据 IPLD 数据模型定义）;

> - **The Node interfaces map cleanly to the IPLD Data Model**;
> - Many features known to be legacy are dropped;
> - The Link implementations are purely CIDs (no "name" nor "size" properties);
> - The Path implementations are provided in the same box;
> - The JSON and CBOR implementations are provided in the same box;
> - Several odd dependencies on blockstore and other interfaces that were closely coupled with IPFS are replaced by simpler, less-coupled interfaces;
> - New features like IPLD Selectors are only available from go-ipld-prime;
> - New features like ADLs (Advanced Data Layouts), which provide features like transparent sharding and indexing for large data, are only available from go-ipld-prime;
> - Declarative transformations can be applied to IPLD data (defined in terms of the IPLD Data Model) using go-ipld-prime;
> - and many other small refinements.

In particular, the clean and direct mapping of "Node" to concepts in the IPLD Data Model
ensures a much more consistent set of rules when working with go-ipld-prime data, regardless of which codecs are involved.
(Codec-specific embellishments and edge-cases were common in the previous generation of libraries.)
This clarity is also what provides the basis for features like Selectors, ADLs, and operations such as declarative transformations.

Many of these changes had been discussed for the other IPLD codebases as well,
but we chose clean break v2 as a more viable project-management path.
Both go-ipld-prime and these legacy libraries can co-exist on the same import path, and both refer to the same kinds of serial data.
Projects wishing to migrate can do so smoothly and at their leisure.

We now consider many of the earlier golang IPLD libraries to be defacto deprecated,
and you should expect new features *here*, rather than in those libraries.
(Those libraries still won't be going away anytime soon, but we really don't recomend new construction on them.)

### migrating

**For recommendations on where to start when migrating:**
see [README_migrationGuide](./README_migrationGuide.md).
That document will provide examples of which old concepts and API names map to which new APIs,
and should help set you on the right track.

### unixfsv1

Lots of people who hear about IPLD have heard about it through IPFS.
IPFS has IPLD-native APIs, but IPFS *also* makes heavy use of a specific system called "UnixFSv1",
so people often wonder if UnixFSv1 is supported in IPLD libraries.

The answer is "yes" -- but it's not part of the core.

UnixFSv1 is now treated as an [ADL](https://ipld.io/glossary/#adl),
and a go-ipld-prime compatible implementation can be found
in the [ipfs/go-unixfsnode](https://github.com/ipfs/go-unixfsnode) repo.

Additionally, the codec used in UnixFSv1 -- dag-pb --
can be found implemented in the [ipld/go-codec-dagpb](https://github.com/ipld/go-codec-dagpb) repo.

A "some assembly required" advisory may still be in effect for these pieces;
check the readmes in those repos for details on what they support.

The move to making UnixFSv1 a non-core system has been an arduous retrofit.
However, framing it as an ADL also provides many advantages:

- it demonstrates that ADLs as a plugin system _work_, and others can develop new systems in this pattern!
- it has made pathing over UnixFSv1 much more standard and well-defined
- this standardization means systems like [Selectors](https://ipld.io/glossary/#selectors) work naturally over UnixFSv1...
- ... which in turn means anything using them (ex: CAR export; graphsync; etc) can very easily be asked to produce a merkle-proof
  for a path over UnixFSv1 data, without requiring the querier to know about the internals.  Whew!

We hope users and developers alike will find value in how these systems are now layered.



Change Policy
-------------

The go-ipld-prime library is ready to use, and we value stability highly.

We make releases periodically.
However, using a commit hash to pin versions precisely when depending on this library is also perfectly acceptable.
(Only commit hashes on the master branch can be expected to persist, however;
depending on a commit hash in a branch is not recommended.  See [development branches](#development-branches).)

We maintain a [CHANGELOG](CHANGELOG.md)!
Please read it, when updating!

We do make reasonable attempts to minimize the degree of changes to the library which will create "breaking change" experiences for downstream consumers,
and we do document these in the changelog (often, even with specific migration instructions).
However, we do also still recommend running your own compile and test suites as a matter of course after updating.

You can help make developing this library easier by staying up-to-date as a downstream consumer!
When we do discover a need for API changes, we typically try to introduce the new API first,
and do at least one release tag in which the old API is deprecated (but not yet removed).
We will all be able to develop software faster, together, as an ecosystem,
if libraries can keep reasonably closely up-to-date with the most recent tags.


### Version Names

When a tag is made, version number steps in go-ipld-prime advance as follows:

1. the number bumps when the lead maintainer says it does.
2. even numbers should be easy upgrades; odd numbers may change things.
3. the version will start with `v0.` until further notice.

[This is WarpVer](https://gist.github.com/warpfork/98d2f4060c68a565e8ad18ea4814c25f).

These version numbers are provided as hints about what to expect,
but ultimately, you should always invoke your compiler and your tests to tell you about compatibility,
as well as read the [changelog](CHANGELOG.md).


### Updating

**Read the [CHANGELOG](CHANGELOG.md).**

Really, read it.  We put exact migration instructions in there, as much as possible.  Even outright scripts, when feasible.

An even-number release tag is usually made very shortly before an odd number tag,
so if you're cautious about absorbing changes, you should update to the even number first,
run all your tests, and *then* upgrade to the odd number.
Usually the step to the even number should go off without a hitch, but if you *do* get problems from advancing to an even number tag,
A) you can be pretty sure it's a bug, and B) you didn't have to edit a bunch of code before finding that out.


### Development branches

The following are norms you can expect of changes to this codebase, and the treatment of branches:

- The `master` branch will not be force-pushed.
    - (exceptional circumstances may exist, but such exceptions will only be considered valid for about as long after push as the "$N-second-rule" about dropped food).
    - Therefore, commit hashes on master are gold to link against.
- All other branches *can* be force-pushed.
    - Therefore, commit hashes not reachable from the master branch are inadvisable to link against.
- If it's on master, it's understood to be good, in as much as we can tell.
    - Changes and features don't get merged until their tests pass!
    - Packages of "alpha" developmental status may exist, and be more subject to change than other more finalized parts of the repo, but their self-tests will at least pass.
- Development proceeds -- both starting from and ending on -- the `master` branch.
    - There are no other long-running supported-but-not-master branches.
    - The existence of tags at any particular commit do not indicate that we will consider starting a long running and supported diverged branch from that point, nor start doing backports, etc.
