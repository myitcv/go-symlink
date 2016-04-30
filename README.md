## go-symlink example repository

See the associated [Go proposal](https://docs.google.com/document/d/1n5y3mZPs_4PjI80a0vZEaHLe7r9PeiiE9xsIrQFT8Is)

This repository exports three public packages:

* `mylib1` - a very boring package
* `mylib2` - another very boring package
* `cmd/a` - an uninteresting binary

`cmd/a` is `go get`-able because it [vendors](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)
(Go 1.5 definition) all of its external dependencies.

However, per best practice, `mylib{1,2}` do not vendor (Go 1.5 definition) their external dependencies. Instead, for
developers of this repository (and by extension `mylib{1,2}`, `cmd/a`) we follow a different approach.

## Developers of this repository

In order to have repeatable builds and consistent development experience for _developers of the repository_ all external
dependencies are "vendored" (**NOT** Go 1.5 definition) into `_vendor`. As described below, these developers are required
to augment their `GOPATH`.

Here, "vendored" means the _entire_ external dependency is copied (test files and all) into `_vendor`. This ensures that
at any point in time we can verify we have a "compatible" set of external dependencies.

To avoid having multiple copies of external dependencies, `cmd/a` shares the same copy of the external dependencies
via a symlink (in the [proposal](https://docs.google.com/document/d/1n5y3mZPs_4PjI80a0vZEaHLe7r9PeiiE9xsIrQFT8Is) this
Unix symlink is replaced with a Go "symlink")

## Getting started as a developer of this repository

_**Note:** this repository currently makes use of symlinks and so is not guaranteed to work on all platforms (read
Windows). The [following proposal](https://docs.google.com/document/d/1n5y3mZPs_4PjI80a0vZEaHLe7r9PeiiE9xsIrQFT8Is)
discusses a means by which symlink-like behaviour be introduced to `cmd/go` and friends_

Augmentation of `GOPATH` is required to develop this repository.

Code can be `go get` in the usual way:

```
go get github.com/myitcv/go-symlink/{mylib1,mylib2,cmd/a}
```

But when working on packages contained within `github.com/myitcv/go-symlink`, your `GOPATH` must be augmented to include `_vendor`:

```
$ pwd
/path/to/gopath/src/github.com/myitcv/go-symlink
$ echo $GOPATH
/path/to/gopath/src/github.com/myitcv/go-symlink/_vendor:/path/to/gopath
$ go test ./...
?       github.com/myitcv/go-symlink/cmd/a      [no test files]
?       github.com/myitcv/go-symlink/cmd/a/internal/cmdinternallib1     [no test files]
?       github.com/myitcv/go-symlink/mylib1     [no test files]
?       github.com/myitcv/go-symlink/mylib2     [no test files]
```

We can also ensure we have "compatible" external dependencies:

```
$ cd _vendor/src
$ go test ./...
ok      bitbucket.org/pkg/inflect       0.005s
ok      github.com/pborman/uuid 0.003s
```
