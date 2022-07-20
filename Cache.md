# Cache

<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Go Command](#go-command)
- [Install a Package or Update](#install-a-package-or-update)
- [Online Documentation & Links](#online-documentation-links)
- [Tools & Co](#tools-co)
  - [Overview](#overview)
  - [Callvis](#callvis)
  - [Go Imports](#go-imports)
    - [Install](#install)
    - [Usage](#usage)
  - [hey - Load Tests](#hey-load-tests)
    - [Installing](#installing)
    - [Usage](#usage-1)

<!-- /code_chunk_output -->

## Go Command

```shell
$ go help
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

  bug         start a bug report
  build       compile packages and dependencies
  clean       remove object files and cached files
  doc         show documentation for package or symbol
  env         print Go environment information
  fix         update packages to use new APIs
  fmt         gofmt (reformat) package sources
  generate    generate Go files by processing source
  get         add dependencies to current module and install them
  install     compile and install packages and dependencies
  list        list packages or modules
  mod         module maintenance
  work        workspace maintenance
  run         compile and run Go program
  test        test packages
  tool        run specified go tool
  version     print Go version
  vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

  buildconstraint build constraints
  buildmode       build modes
  c               calling between Go and C
  cache           build and test caching
  environment     environment variables
  filetype        file types
  go.mod          the go.mod file
  gopath          GOPATH environment variable
  gopath-get      legacy GOPATH go get
  goproxy         module proxy protocol
  importpath      import path syntax
  modules         modules, module versions, and more
  module-get      module-aware go get
  module-auth     module authentication using go.sum
  packages        package lists and patterns
  private         configuration for downloading non-public code
  testflag        testing flags
  testfunc        testing functions
  vcs             controlling version control with GOVCS

Use "go help <topic>" for more information about that topic.
```

```shell
$ go tool

  addr2line
  asm
  buildid
  cgo
  compile
  cover
  dist
  doc
  fix
  link
  nm
  objdump
  pack
  pprof
  test2json
  trace
  vet
```

## Install a Package or Update

```shell
go install https://github.com/rakyll/hey
go install https://github.com/rakyll/hey@latest
```

## Online Documentation & Links

- [What is Software Engineering with Go?](https://research.swtch.com/vgo-eng)
- [Go und Versioning](https://research.swtch.com/vgo)

## Tools & Co

### Overview

<https://pkg.go.dev/golang.org/x/tools>

### Callvis

Visualisation of Structure

```shell
go install github.com/ofabry/go-callvis@latest

```

### Go Imports

#### Install

```shell
go install golang.org/x/tools/cmd/goimports@latest
```

#### Usage

```shell
goimports -f -w .
```

### hey - Load Tests

<https://github.com/rakyll/hey>

#### Installing

```shell
brew install hey or
go install https://github.com/rakyll/hey
```

#### Usage

```shell
hey https://myserver.example.com
```
