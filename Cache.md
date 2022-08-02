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
  - [golangci-lint](#golangci-lint)
    - [Installation](#installation)
    - [Usage](#usage-2)
      - [Show Configuration](#show-configuration)
    - [Modify Configuration](#modify-configuration)
- [Visual Studio Code](#visual-studio-code)
  - [Language Server](#language-server)
- [Differences to JS/TS](#differences-to-jsts)

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

### golangci-lint

#### Installation

```shell
brew install golangci-lint

# brew reinstall golangci-lint
# brew upgrade golangci-lint
```

#### Usage

##### Show Configuration

```shell
golangci-lint help linters

Enabled by default linters:

deadcode: Finds unused code [fast: false, auto-fix: false]
errcheck: Errcheck is a program for checking for unchecked errors in go programs. These unchecked errors can be critical bugs in some cases [fast: false, auto-fix: false]
gosimple (megacheck): Linter for Go source code that specializes in simplifying code [fast: false, auto-fix: false]
govet (vet, vetshadow): Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string [fast: false, auto-fix: false]
ineffassign: Detects when assignments to existing variables are not used [fast: true, auto-fix: false]
staticcheck (megacheck): It's a set of rules from staticcheck. It's not the same thing as the staticcheck binary. The author of staticcheck doesn't support or approve the use of staticcheck as a library inside golangci-lint. [fast: false, auto-fix: false]
structcheck: Finds unused struct fields [fast: false, auto-fix: false]
typecheck: Like the front-end of a Go compiler, parses and type-checks Go code [fast: false, auto-fix: false]
unused (megacheck): Checks Go code for unused constants, variables, functions and types [fast: false, auto-fix: false]
varcheck: Finds unused global variables and constants [fast: false, auto-fix: false]

Disabled by default linters:

asasalint: check for pass []any as any in variadic func(...any) [fast: false, auto-fix: false]
asciicheck: Simple linter to check that your code does not contain non-ASCII identifiers [fast: true, auto-fix: false]
bidichk: Checks for dangerous unicode character sequences [fast: true, auto-fix: false]
bodyclose: checks whether HTTP response body is closed successfully [fast: false, auto-fix: false]
containedctx: containedctx is a linter that detects struct contained context.Context field [fast: true, auto-fix: false]
contextcheck: check the function whether use a non-inherited context [fast: false, auto-fix: false]
cyclop: checks function and package cyclomatic complexity [fast: false, auto-fix: false]
decorder: check declaration order and count of types, constants, variables and functions [fast: true, auto-fix: false]
depguard: Go linter that checks if package imports are in a list of acceptable packages [fast: true, auto-fix: false]
dogsled: Checks assignments with too many blank identifiers (e.g. x, _, _, _, := f()) [fast: true, auto-fix: false]
dupl: Tool for code clone detection [fast: true, auto-fix: false]
durationcheck: check for two durations multiplied together [fast: false, auto-fix: false]
errchkjson: Checks types passed to the json encoding functions. Reports unsupported types and optionally reports occasions, where the check for the returned error can be omitted. [fast: false, auto-fix: false]
errname: Checks that sentinel errors are prefixed with the `Err` and error types are suffixed with the `Error`. [fast: false, auto-fix: false]
errorlint: errorlint is a linter for that can be used to find code that will cause problems with the error wrapping scheme introduced in Go 1.13. [fast: false, auto-fix: false]
execinquery: execinquery is a linter about query string checker in Query function which reads your Go src files and warning it finds [fast: false, auto-fix: false]
exhaustive: check exhaustiveness of enum switch statements [fast: false, auto-fix: false]
exhaustivestruct [deprecated]: Checks if all struct's fields are initialized [fast: false, auto-fix: false]
exhaustruct: Checks if all structure fields are initialized [fast: false, auto-fix: false]
exportloopref: checks for pointers to enclosing loop variables [fast: false, auto-fix: false]
forbidigo: Forbids identifiers [fast: true, auto-fix: false]
forcetypeassert: finds forced type assertions [fast: true, auto-fix: false]
funlen: Tool for detection of long functions [fast: true, auto-fix: false]
gci: Gci controls golang package import order and makes it always deterministic. [fast: true, auto-fix: false]
gochecknoglobals: check that no global variables exist [fast: true, auto-fix: false]
gochecknoinits: Checks that no init functions are present in Go code [fast: true, auto-fix: false]
gocognit: Computes and checks the cognitive complexity of functions [fast: true, auto-fix: false]
goconst: Finds repeated strings that could be replaced by a constant [fast: true, auto-fix: false]
gocritic: Provides diagnostics that check for bugs, performance and style issues. [fast: false, auto-fix: false]
gocyclo: Computes and checks the cyclomatic complexity of functions [fast: true, auto-fix: false]
godot: Check if comments end in a period [fast: true, auto-fix: true]
godox: Tool for detection of FIXME, TODO and other comment keywords [fast: true, auto-fix: false]
goerr113: Golang linter to check the errors handling expressions [fast: false, auto-fix: false]
gofmt: Gofmt checks whether code was gofmt-ed. By default this tool runs with -s option to check for code simplification [fast: true, auto-fix: true]
gofumpt: Gofumpt checks whether code was gofumpt-ed. [fast: true, auto-fix: true]
goheader: Checks is file header matches to pattern [fast: true, auto-fix: false]
goimports: In addition to fixing imports, goimports also formats your code in the same style as gofmt. [fast: true, auto-fix: true]
golint [deprecated]: Golint differs from gofmt. Gofmt reformats Go source code, whereas golint prints out style mistakes [fast: false, auto-fix: false]
gomnd: An analyzer to detect magic numbers. [fast: true, auto-fix: false]
gomoddirectives: Manage the use of 'replace', 'retract', and 'excludes' directives in go.mod. [fast: true, auto-fix: false]
gomodguard: Allow and block list linter for direct Go module dependencies. This is different from depguard where there are different block types for example version constraints and module recommendations. [fast: true, auto-fix: false]
goprintffuncname: Checks that printf-like functions are named with `f` at the end [fast: true, auto-fix: false]
gosec (gas): Inspects source code for security problems [fast: false, auto-fix: false]
grouper: An analyzer to analyze expression groups. [fast: true, auto-fix: false]
ifshort: Checks that your code uses short syntax for if-statements whenever possible [fast: true, auto-fix: false]
importas: Enforces consistent import aliases [fast: false, auto-fix: false]
interfacer [deprecated]: Linter that suggests narrower interface types [fast: false, auto-fix: false]
ireturn: Accept Interfaces, Return Concrete Types [fast: false, auto-fix: false]
lll: Reports long lines [fast: true, auto-fix: false]
maintidx: maintidx measures the maintainability index of each function. [fast: true, auto-fix: false]
makezero: Finds slice declarations with non-zero initial length [fast: false, auto-fix: false]
maligned [deprecated]: Tool to detect Go structs that would take less memory if their fields were sorted [fast: false, auto-fix: false]
misspell: Finds commonly misspelled English words in comments [fast: true, auto-fix: true]
nakedret: Finds naked returns in functions greater than a specified function length [fast: true, auto-fix: false]
nestif: Reports deeply nested if statements [fast: true, auto-fix: false]
nilerr: Finds the code that returns nil even if it checks that the error is not nil. [fast: false, auto-fix: false]
nilnil: Checks that there is no simultaneous return of `nil` error and an invalid value. [fast: false, auto-fix: false]
nlreturn: nlreturn checks for a new line before return and branch statements to increase code clarity [fast: true, auto-fix: false]
noctx: noctx finds sending http request without context.Context [fast: false, auto-fix: false]
nolintlint: Reports ill-formed or insufficient nolint directives [fast: true, auto-fix: false]
nonamedreturns: Reports all named returns [fast: false, auto-fix: false]
nosnakecase: nosnakecase is a linter that detects snake case of variable naming and function name. [fast: true, auto-fix: false]
nosprintfhostport: Checks for misuse of Sprintf to construct a host with port in a URL. [fast: true, auto-fix: false]
paralleltest: paralleltest detects missing usage of t.Parallel() method in your Go test [fast: false, auto-fix: false]
prealloc: Finds slice declarations that could potentially be pre-allocated [fast: true, auto-fix: false]
predeclared: find code that shadows one of Go's predeclared identifiers [fast: true, auto-fix: false]
promlinter: Check Prometheus metrics naming via promlint [fast: true, auto-fix: false]
revive: Fast, configurable, extensible, flexible, and beautiful linter for Go. Drop-in replacement of golint. [fast: false, auto-fix: false]
rowserrcheck: checks whether Err of rows is checked successfully [fast: false, auto-fix: false]
scopelint [deprecated]: Scopelint checks for unpinned variables in go programs [fast: true, auto-fix: false]
sqlclosecheck: Checks that sql.Rows and sql.Stmt are closed. [fast: false, auto-fix: false]
stylecheck: Stylecheck is a replacement for golint [fast: false, auto-fix: false]
tagliatelle: Checks the struct tags. [fast: true, auto-fix: false]
tenv: tenv is analyzer that detects using os.Setenv instead of t.Setenv since Go1.17 [fast: false, auto-fix: false]
testpackage: linter that makes you use a separate _test package [fast: true, auto-fix: false]
thelper: thelper detects golang test helpers without t.Helper() call and checks the consistency of test helpers [fast: false, auto-fix: false]
tparallel: tparallel detects inappropriate usage of t.Parallel() method in your Go test codes [fast: false, auto-fix: false]
unconvert: Remove unnecessary type conversions [fast: false, auto-fix: false]
unparam: Reports unused function parameters [fast: false, auto-fix: false]
varnamelen: checks that the length of a variable's name matches its scope [fast: false, auto-fix: false]
wastedassign: wastedassign finds wasted assignment statements. [fast: false, auto-fix: false]
whitespace: Tool for detection of leading and trailing whitespace [fast: true, auto-fix: true]
wrapcheck: Checks that errors returned from external packages are wrapped [fast: false, auto-fix: false]
wsl: Whitespace Linter - Forces you to use empty lines! [fast: true, auto-fix: false]

Linters presets:

bugs: asasalint, asciicheck, bidichk, bodyclose, contextcheck, durationcheck, errcheck, errchkjson, errorlint, exhaustive, exportloopref, gosec, govet, makezero, nilerr, noctx, rowserrcheck, scopelint, sqlclosecheck, staticcheck, typecheck
comment: godot, godox, misspell
complexity: cyclop, funlen, gocognit, gocyclo, maintidx, nestif
error: errcheck, errorlint, goerr113, wrapcheck
format: decorder, gci, gofmt, gofumpt, goimports
import: depguard, gci, goimports, gomodguard
metalinter: gocritic, govet, revive, staticcheck
module: depguard, gomoddirectives, gomodguard
performance: bodyclose, maligned, noctx, prealloc
sql: execinquery, rowserrcheck, sqlclosecheck
style: asciicheck, containedctx, decorder, depguard, dogsled, dupl, errname, exhaustivestruct, exhaustruct, forbidigo, forcetypeassert, gochecknoglobals, gochecknoinits, goconst, gocritic, godot, godox, goerr113, goheader, golint, gomnd, gomoddirectives, gomodguard, goprintffuncname, gosimple, grouper, ifshort, importas, interfacer, ireturn, lll, makezero, misspell, nakedret, nilnil, nlreturn, nolintlint, nonamedreturns, nosnakecase, nosprintfhostport, paralleltest, predeclared, promlinter, revive, stylecheck, tagliatelle, tenv, testpackage, thelper, tparallel, unconvert, varnamelen, wastedassign, whitespace, wrapcheck, wsl
test: exhaustivestruct, exhaustruct, paralleltest, testpackage, tparallel
unused: deadcode, ineffassign, structcheck, unparam, unused, varcheck
```

#### Modify Configuration

Pass `-E/--enable` to enable linter and `-D/--disable` to disable:

```shell
golangci-lint run --disable-all -E errcheck
```

## Visual Studio Code

### Language Server

Details here: <https://microsoft.github.io/language-server-protocol/>

## Differences to JS/TS

- complex, real, imag,
- len, cap
- iota
