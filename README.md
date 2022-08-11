# GoLang

Own Notices

<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=4 orderedList=false} -->

<!-- code_chunk_output -->

- [Getting Started](#getting-started)
  - [Install Go via Brew](#install-go-via-brew)
  - [Uninstall Go](#uninstall-go)
  - [Go Versions Manager](#go-versions-manager)
  - [VsCode Install/Update Tools](#vscode-installupdate-tools)
- [Basics](#basics)
  - [Pre-Conditions](#pre-conditions)
  - [Hello World](#hello-world)
- [Development Environment](#development-environment)
  - [Linting & Vetting](#linting-vetting)
  - [Makefiles](#makefiles)
- [Types and Declarations](#types-and-declarations)
  - [Common Concepts](#common-concepts)
    - [Zero](#zero)
    - [var vs. `:=`](#var-vs)
    - [Explicit Type Conversion (= Automatic Type Promotion)](#explicit-type-conversion-automatic-type-promotion)
    - [Literals](#literals)
  - [Built-in Types](#built-in-types)
    - [Booleans](#booleans)
    - [Integer](#integer)
    - [Floating Point Types](#floating-point-types)
    - [Strings](#strings)
  - [Const](#const)
  - [Unused Variables & Constants](#unused-variables-constants)
  - [Naming Variables and Constants](#naming-variables-and-constants)
- [Composite Types](#composite-types)
  - [Arrays](#arrays)
  - [Slices](#slices)
    - [append (similar to push in JS)](#append-similar-to-push-in-js)
    - [Runtime Capacity](#runtime-capacity)
    - [make](#make)
    - [`nil` vs `zero` Declarations](#nil-vs-zero-declarations)
    - [Slice-slicing](#slice-slicing)
    - [Sharing Memory](#sharing-memory)
    - [Converting Array to Slices](#converting-array-to-slices)
    - [`copy` helps you to avoid memory sharing problems](#copy-helps-you-to-avoid-memory-sharing-problems)
  - [Strings, Runes, Bytes](#strings-runes-bytes)
    - [Conversion](#conversion)
    - [Strings to Slices](#strings-to-slices)
  - [Maps](#maps)
    - [With `make`](#with-make)
    - [`ok` &rarr; Comma `ok` idiom](#ok-rarr-comma-ok-idiom)
    - [Deleting](#deleting)

<!-- /code_chunk_output -->

## Getting Started

### Install Go via Brew

```shell
brew install go
```

### Uninstall Go

Online: <https://blog.dharnitski.com/2019/04/06/uninstall-go-on-mac/>

##### If previously installed via Brew

```shell
brew uninstall dep
brew uninstall go
```

##### If previously installed via Pkgutil

```shell
pkgutil --pkgs | grep go   # find in the list
sudo pkgutil --forget org.golang.go
```

### Go Versions Manager

Online: <https://github.com/kevincobain2000/gobrew>

##### Installing

```shell
curl -sLk https://git.io/gobrew | sh -  # Installation
gobrew use 1.16.4 # Download, install and use in one step

gobrew install 1.16.4  # install only
gobrew use 1.16.4 # change to this version

go uninstall 1.16 # uninstall a certain version
```

##### ENV VARIABLES (Important for VSCode)

VSCode needs GOPATH and GOBIN to detect the currently used version, if e.g. a
package manager like gobrew is installed:

```shell
# ~/.zshrc excerpt

export PATH="$HOME/.gobrew/current/bin:$HOME/.gobrew/bin:$PATH"
export GOPATH="$HOME/.gobrew/current"
export GOBIN="$HOME/.gobrew/current/bin"
```

### VsCode Install/Update Tools

> **Notice**: You have to repeat this step if you use a go versions manager and
> and change the current version!

> Further info for VSCode: <https://github.com/golang/vscode-go/wiki/tools>

> Further General Infos: <https://pkg.go.dev/golang.org/x/tools>

1. F1 -> Go: Install/Update Tools
2. Select all and click OK

![VSCode installing Go Tools](assets/go-install-tools-vscode.png)

You see then in output window, something like:

```zsh

Installing 10 tools at the configured GOBIN:
  gotests
  gomodifytags
  impl
  goplay
  gopkgs
  go-outline  # within gopls in newer versions
  dlv
  dlv-dap
  staticcheck
  gopls

.. github.com/cweill/gotests/gotests@latest (..go/bin/gotests) SUCCEEDED
.. github.com/fatih/gomodifytags@latest (..go/bin/gomodifytags) SUCCEEDED
.. github.com/josharian/impl@latest (..go/bin/impl) SUCCEEDED
.. github.com/haya14busa/goplay/cmd/goplay@latest (..go/bin/goplay) SUCCEEDED
.. github.com/go-delve/delve/cmd/dlv@latest (..go/bin/dlv) SUCCEEDED
.. honnef.co/go/tools/cmd/staticcheck@latest (..go/bin/staticcheck) SUCCEEDED
.. golang.org/x/tools/gopls@v0.9.1 (..go/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go. :)

```

More information to the tools online:

##### Go Ppkgs

`gopkgs` is a tool that provides list of available Go packages that can be
imported. This is an alternative to `go list all`, just faster.

Online: <https://github.com/uudashr/gopkgs/>

##### Go Outline

Simple utility for extracting a JSON representation of the declarations in a Go
source file.

<https://github.com/ramya-rao-a/go-outline>

##### Go Imports

```shell
go install golang.org/x/tools/cmd/goimports@latest
```

##### Go Tests

`gotests` makes writing Go tests easy. It's a Golang commandline tool that
generates table driven tests based on its target source files' function and
method signatures. Any new dependencies in the test files are automatically
imported.

<https://github.com/cweill/gotests/>

##### Go Modify Tags

Go tool to modify/update field tags in structs. `gomodifytags` makes it easy to
update, add or delete the tags in a struct field. You can easily add new tags,
update existing tags (such as appending a new key, i.e: db, xml, etc..) or
remove existing tags. It also allows you to add and remove tag options. It's
intended to be used by an editor, but also has modes to run it from the
terminal. Read the usage section below for more information.

![Go Modify Tags Use](assets/go-modify-tags-use.gif)

<https://www.github.com/fatih/gomodifytags>

##### Go Impl

`impl` generates method stubs for implementing an interface.

```go
$ impl 'f *File' io.ReadWriteCloser
func (f *File) Read(p []byte) (n int, err error) {
	panic("not implemented")
}

func (f *File) Write(p []byte) (n int, err error) {
	panic("not implemented")
}

func (f *File) Close() error {
	panic("not implemented")
}

# You can also provide a full name by specifying the package path.
# This helps in cases where the interface can't be guessed
# just from the package name and interface name.
$ impl 's *Source' golang.org/x/oauth2.TokenSource
func (s *Source) Token() (*oauth2.Token, error) {
    panic("not implemented")
}
```

<https://www.github.com/josharian/impl>

##### Go Play Ground Client

```shell
goplay some-example.go # opens browser
```

<https://www.github.com/haya14busa/goplay/>

##### Go Delve

A Debugger for the Go Programming Language

GitHub: <https://www.github.com/go-delve/delve/> Getting Started:
<https://github.com/go-delve/delve/blob/master/Documentation/cli/getting_started.md>

##### Go Tools from Dominik Honnef: Go `staticcheck`

`Staticcheck` is a state of the art linter for the Go programming language.
Using static analysis, it finds bugs and performance issues, offers
simplifications, and enforces style rules.

GitHub: <https://github.com/dominikh/go-tools> Online Documentation:
<https://staticcheck.io/docs/>

##### gopls, the Go Language Server

> MacOS: XCode and Command Line Tools may be necessary

`gopls` (pronounced "Go please") is the official Go language server developed by
the Go team. It provides IDE features to any LSP-compatible editor.

You should not need to interact with gopls directly--it will be automatically
integrated into your editor.

<https://pkg.go.dev/golang.org/x/tools/gopls>

<https://www.golang.org/x/tools/gopls>

## Basics

### Pre-Conditions

1. You have installed go and registered necessary PATHs in your shell
   configuration, e.g:

```shell
PATH="/usr/local/sbin:$PATH"
PATH="$HOME/.gobrew/current/bin:$HOME/.gobrew/bin:$PATH"
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
```

If you want to use private packages and bypass proxy & co:

```shell
export GOPRIVATE=example.com/*,example2.com/*,ex3.com/whatever
export GONOSUMDB=example.com/*,example2.com/*,ex3.com/whatever
export GONOPROXY=example.com/*,example2.com/*,ex3.com/whatever
```

2. Choose your workspace within your path, e.g:

```shell
mkdir $GOPATH/src/MyGoProject/
cd    $GOPATH/src/MyGoProject/
```

3. Initialise Module and Create a Main File

```shell
go mod init
touch main.go
```

### Hello World

A complete program is created by linking a single, unimported package called the
`main package` with all the packages it imports, transitively. The main package
must have package name main and `declare a function main` that takes no
arguments and returns no value.

```go
// main.go

package main

import "fmt"

func main() {
  fmt.Println(("Hello World"))
}
```

```shell
go run main.go    # Runs directly without building
go build          # Works if there is a module
go build main.go  # or so
./main.go         # prints Hello World or
./MyGoProject      # the same like above
```

At this time you have the following files in your workspace:

```shell
MyGoProject       # Executable (created by `go build` without file name)
go.mod           # Text File, e.g. ModuleName and used Go Version etc.
main             # Executable (created by `go build main.go`)
main.go          # Source File
```

That's the `go.mod`

```go
module MyGoProject

go 1.17
```

You can also use your GitHub repo as a module like:

```shell
go mod init github.com/webia1/my-go-project
```

## Development Environment

### Linting & Vetting

`golint` lints the Go source files named on its command line.

```shell
go install golang.org/x/lint/golint@latest
go lint ./...           # 3 DOTS
```

`vet` examines Go source code and reports suspicious constructs, such as Printf
calls whose arguments do not align with the format string. Vet uses heuristics
that do not guarantee all reports are genuine problems, but it can find errors
not caught by the compilers.

Vet is normally invoked through the go command. This command vets the package in
the current directory (no installation required):

```shell
go vet                    # or
go vet my/project/...     # or
go vet ./...              # 3 DOTS!
```

`golangci-lint` combines `golint` and `go wet`, it runs linters in parallel,
uses caching, supports yaml config, has integrations with all major IDE and has
dozens of linters included. (Documentation: <https://golangci-lint.run/>)

```shell
brew install golangci-lint

```

`Staticcheck` is a state of the art linter for the Go programming language.
Using static analysis, it finds bugs and performance issues, offers
simplifications, and enforces style rules. <https://staticcheck.io/>.
Configuration: <https://staticcheck.io/docs/configuration/>. (Installed by
VSCode Tools or see below).

```shell
go install honnef.co/go/tools/cmd/staticcheck@latest
```

### Makefiles

> **Makefiles for Go Developers**:
> <https://tutorialedge.net/golang/makefiles-for-go-developers/>

Go developers have adopted `make` as their solution (save as `Makefile`):

```Makefile
.DEFAULT_GOAL := build

fmt:
    go fmt ./...
.PHONY:fmt          # <-- Self chosen name for `fmt` part

lint: fmt            # <-- `fmt` is the pre-condition for `lint`
    golint ./...
.PHONY:lint

vet: fmt
    go vet ./...
.PHONY:vet
```

**Makefiles are extremely picky:** You must indent the steps in a target with a
`tab`.

> A **`PHONY`** target is one that is not really the name of a file; rather it
> is just a name for a recipe to be executed when you make an explicit request.
> There are two reasons to use a phony target: to avoid a conflict with a file
> of the same name, and to improve performance.

Once the `Makefile` is in the `"src"` directory (any name can be chosen), type:

```shell
make
```

## Types and Declarations

### Common Concepts

#### Zero

Variable is declared but not assigned a value. (Like `null` in JS)

#### var vs. `:=`

```go
var x int = 10          // is the same like
var x = 10              // because it is assigned, no need for type
x := 10                 // is the same like the two declarations above
var x int               // something like "let x: int = null" in TS
var x, y int = 10, 20   // more than one declarations
var x, y int            // with zero values
var x, y = 10, "Hi"     // with different types
x, y := 10, "Hi"        // same like above

// also a "declaration list" would be possible

var (
  a     int
  b           = 20
  c     int   = 30
  d, e        = 40, "Hi"
  f, g  string
)
```

The `:=` operator can reassign (not possible by using `var`):

```go
x     := 10
x, y  := 30, "Hi"
```

One limitation for `:=`: At package level you must use var because it is not
legal outside of functions.

**Important notices:**

- Initialisation with zero values -> better `var` than `:=`
- prefer something like `var x byte = 20` to `x := byte(20)`
- `:=` allows you to reassign too. Attention: &rarr; `Shadowing Variables`

_**As a general rule: Declare variables in the package block that are
effectively immutable.**_

#### Explicit Type Conversion (= Automatic Type Promotion)

Go doesn't allow automatic type conversion, when variable types do not match.

#### Literals

- Integer Literals, based on 10, except:
  - `Ob` binary
  - `Oo` octal (0 with no letter after it is octal too, but don't use it)
  - `0x` hexadecimal
- Floating Point Literal
  - e.g. 7.11e23
  - `0x` hexadecimal
  - `p` exponent
  - `_` formatting big numbers
- Rune Literals (Chars in JS)

  - e.g. `('a')`, `('\171')`, `('\x47')`
  - 16 Bit Hexadecimal `('\u0061')`
  - 32 Bit Unicode `('\U00000061')`
  - Newline `('\n')`
  - Tabulator `('\t')`
  - Octal (rare)

- String Literals (very similar to JS)
  - Double Quotes
    - `"Hello World"`
    - `"My \"Hello World\""` If double quotes within -> escape them
  - or Backtick (also called Backquotes)
    - In this case, you don't have to escape double quotes within strings:
    - `` `My "Hello World"` ``

### Built-in Types

- boolean: `bool`
- integer
- float
- string

#### Booleans

```go
var myFlag bool       // no value assigned -> false
var myFlag = true     // it is bool and true
```

#### Integer

`NaN` -> Similar to JS

> **Important:** If you assign a type and then use a number **larger than the
> types range** to assign it, **it will fail**.

> If you convert to a type that has range lower than your current range, **data
> loss will occur**.

Special name `byte` is an alias for `uint8`, and the other special name `int` is
CPU dependent (e.g. int32 or int64).

> There are some uncommon 64Bit CPU architectures with 32 bit signed integer: Go
> supports: `amd64p32`, `mip64p32`, and `mips64p32le`

> Go does not have generics and function overloadings (yet?).

Source: <https://gosamples.dev/int-min-max/>

To get the maximum and minimum value of various integer types in Go, use the
[`math`](https://pkg.go.dev/math) package constants. For example, to get the
minimum value of the `int64` type, which is **\-9223372036854775808**, use the
[`math.MinInt64`](https://pkg.go.dev/math#pkg-constants) constant. To get the
maximum value of the `int64` type, which is **9223372036854775807**, use the
[`math.MaxInt64`](https://pkg.go.dev/math#pkg-constants). To check the minimum
and maximum values of different int types, see the following example and its
output.

> For unsigned integer types, only the max constant is available because the
> minimum value of unsigned types is always 0.

##### Signed integers in Go

Source: <https://golangdocs.com/integers-in-golang>

Signed integer types supported by Go is shown below.

```go
int8    // is -128 to 127
int16   // is -32768 to 32767
int32   // is -2147483648 to 2147483647
int64   // is -9223372036854775808 to 9223372036854775807
```

##### Unsigned integers in Go

```go
uint8   // 0 to 255
uint16  // 0 to 65535
uint32  // is 0 to 4294967295
uint64  // 0 to 18446744073709551615
```

##### Type Conversion

We do typecast by directly using the name of the variable as a function to
convert types:

```go
package main

import ("fmt")

func main() {
    var x int32
    var y uint32     // range 0 to 4294967295
    var z uint8      // range 0 to 255
    fmt.Println("Type Conversion")
    x = 26700
    y = uint32(x)       // data preserved because number is inside range
    z = uint8(x)        // data loss due to out of range conversion
    fmt.Println(y, z)   // prints 26700 76
}
```

##### Integer Operations

- `+, -, *, /, % for modulus` Division by 0 causes so called `panic`
- `==, !=, >, >=, <, <=` Comparisons
- `<<, >>, &, |, ^, &^ ` Bit Manipulations (^ = XOR, &^ = AND NOT)

#### Floating Point Types

Go supports the `IEEE-754` 32-bit and 64-bit floating-point numbers. You can use
all standard number operators with floats except `%`.

IEEEE-754: <https://en.wikipedia.org/wiki/IEEE_754>

> **Do not use them to represent money or whatever needs to have an exact
> decimal representation**

```go
float32
float64
```

Source: https://gosamples.dev/float64-min-max/

The maximum value of the `float64` type in Go is
**1.79769313486231570814527423731704356798070e+308** and you can get this value
using the [`math.MaxFloat64`](https://pkg.go.dev/math#pkg-constants) constant.

The minimum value above zero (smallest positive, non-zero value) of the
`float64` type in Go is **4.9406564584124654417656879286822137236505980e-324**
and you can get this value using the
[`math.SmallestNonzeroFloat64`](https://pkg.go.dev/math#pkg-constants) constant.

The maximum value of the `float32` type in Go is
**3.40282346638528859811704183484516925440e+38** and you can get this value
using the [`math.MaxFloat32`](https://pkg.go.dev/math#pkg-constants) constant.

The minimum value above zero (smallest positive, non-zero value) of the
`float32` type in Go is **1.401298464324817070923729583289916131280e-45** and
you can get this value using the
[`math.SmallestNonzeroFloat32`](https://pkg.go.dev/math#pkg-constants) constant.

##### Type Conversion

Loss of precision will occur when a 64-bit floating-point number is converted to
32-bit float.

Source: <https://golangdocs.com/floating-point-numbers-in-golang>

```go
package main

import (
    "fmt"
)

func main() {
    var f1 float32
    var f2 float64

    f2 = 1.234567890123
    f1 = float32(f2)

    fmt.Println(f1)         // prints "1.2345679"
}
```

##### Complex Numbers

Floating-point numbers are used in complex numbers as well. The real and
imaginary parts are floats.

More information: <https://golangdocs.com/complex-numbers-in-golang>

```go
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))
}

```

##### Matrix (Matrizen)

No matrix support :/

#### Strings

Similar to JS. `Zero value` is empty string.

### Const

Very similar to TS. Constants can be typed or untyped.

If constants are untyped e.g. the following is allowed:

```go
// it is a number but there is no specific type
const x = 10;

// therefore the following assignments are OK
var a int = x       // OK
var b float64 = x   // OK
var c byte = x      // OK
```

But if you give a type to it, you have to consider:

```go
// it is an integer
const x int = 10;
var a int = x       // OK
var b float64 = x   // not OK because int != float
```

### Unused Variables & Constants

Every declared local variable must be read. It is a compile-time error to
declare a local variable and to not read its value.

But the compiler's unused-variable-check is not precise enough; it accepts a
single read, even if there were writes afterwards, same with `go vet`. But
`golangci-lint` can detect them.

The Go compiler does not prevent you from creating unread
package-level-variables.

Suprisingly: Unused constants are OK :)

### Naming Variables and Constants

Very similar to JS/TS, any Unicode (letter/digit) is allowed. See other parts
above (e.g. begining with digit).

`_` is a valid character but Go prefers `camelCase` instead of `snake_case`.

> And undercore `_` by itself is a special identifier name in Go (ignoring a
> parameter, or prop etc., see examples in the coming sections).

Preferred Go Style (as short as possible within block code):

- k, v &rarr; key, value
- i, j &rarr; common names for index variables

## Composite Types

### Arrays

> Completely different than arrays in JS/TS

Confusing definition (By Jon Bodner): "All of the elements in the array must be
of the type that's specified but this does not mean they are always of the same
type". Another quote: **Don't use arrays unless you know the exact length you
need ahead of time.**

```go
var x [3]int                // 3 is the size of the array
                            // No values specified, i.e -> x = [0,0,0]
                            // Zero value for int is 0

var x = [3]int{10, 20, 30} // Values set



// Here is a so called sparse array :)

var x = [8]int{1, 3: 7, 5, 6: 8, 9}  // [1 0 0 7 5 0 8 9]

/**
  First value (index 0) is 1,
  7 is index no 3 (the next one, has the next index no)
  5 is index no 4,
  8 is index no 6 and so on, that means
  9 is the index no 7
  everything else is 0
*/

// You can also leave off number by using `...`

var x = [...]int{1, 2, 3}

x[0] = 10
fmt.Println(x)       // [10 2 3]
fmt.Println(len(x))  // 3

// Simulating more dimensional arrays:
// 2 arrays of length 3 with zero values

var x [2][3]int  // [[0 0 0] [0 0 0]]  -> How to modify them, see link below:

```

See more about: <https://codezup.com/arrays-in-golang-multi-dimensional-arrays/>

### Slices

> Slices looks like arrays with some differences (notice the missing `...`).

> **Zero value for a slice** is **`nil`** (and not `0`).

> (To my mind: Zero is something like `null` in JS, `nil` is like `undefined`)

```go
var x = [...]int{1, 2, 3}               // --> ARRAY
var x = []int{1, 2, 3}                  // --> SLICE

var x = [8]int{1, 3: 7, 5, 6: 8, 9}     // --> ARRAY
var x = []int{1, 3: 7, 5, 6: 8, 9}      // --> SLICE

var x [2][3]int                         // --> ARRAY
var x [][]int                           // --> SLICE
```

**The only thing you can compare a slice with is `nil`.** The `reflect package`
contains a function calles `DeepEqual` can compare almost anything, including
slices.

```go
var x []int
fmt.Println (x == nil) // true
```

#### append (similar to push in JS)

`...` like spread operator in JS but different syntax (postfix instead of
prefix).

> It is a compile-time error if you forget to assign the value returned from
> append. (Go is a **`call by value`** language &rarr; no object references like
> in JS, but real copies)

```go
var x = []int{1, 2, 3}
x = append(x, 10)         // [1 2 3 10]
x = append(x, 5, 7, 9)    // [1 2 3 10 5 7 9]
y := []int{20, 30, 40}
x = append(x, y...)       // [1 2 3 10 5 7 9 20 30 40]

```

#### Runtime Capacity

Runtime capacity is like in C++:

> The rules as of Go 1.14 are to double the size of the slice when the capacity
> is less than 1,024 and then grow by at least 25% afterward.

Just as the built-in **`len`** function returns the current length of a slice,
the built-in **`cap`** function returns the current capacity of a slice. It is
used far less frequently than len.

Cap is typically used to determine whether a slice is big enough to accommodate
new data or whether a call to make is required to create a new slice.

> The cap function also accepts an array as a parameter, although for arrays,
> cap always returns the same value as len.

```go
var x []int
fmt.Println(x, len(x), cap(x))
x = append(x, 10)
fmt.Println(x, len(x), cap(x))
x = append(x, 11)
fmt.Println(x, len(x), cap(x))
x = append(x, 12)
fmt.Println(x, len(x), cap(x))
x = append(x, 13)
fmt.Println(x, len(x), cap(x))
x = append(x, 14)
fmt.Println(x, len(x), cap(x))
```

**outputs:**

```shell
[] 0 0
[10] 1 1
[10 11] 2 2
[10 11 12] 3 4
[10 11 12 13] 4 4
[10 11 12 13 14] 5 8
```

#### make

The built-in `make` function is responsible for creating an empty slice with a
specified length or capacity.

> Your program will panic at runtime if you use a variable to set a capacity
> that is less than the length.

```go
x := make([]int, 5)         // length of 5 and a capacity of 5.
x := make([]int, 5, 10)     // length of 5 and a capacity of 10.


/** The following is tricky:
We cannot directly index into it because it has length 0,
but we can append values to it instead:
*/

x := make([]int, 0, 10)     // length of 0 and a capacity of 10.
x = append(x, 1, 3, 7)

```

A slice's length always increases after an `append`! Make sure that you set the
slice's length before using the `make`; otherwise, your slice may start off with
a surprising number of zero values.

#### `nil` vs `zero` Declarations

```go

// import fmt and reflect before

var data []int            // nil slice declaration
var data  = []int{}       // empty slice with zero-length

var x []int
var y = []int{}

fmt.Println(x, len(x))	  // [] 0  Debugger:  []int len: 0, cap: 0, nil
fmt.Println(y, len(y))	  // [] 0  Debugger:  []int len: 0, cap: 0, []

fmt.Println(x == nil)     // true
fmt.Println(y == nil)     // false

fmt.Println(reflect.TypeOf(x))                      // []int
fmt.Println(reflect.TypeOf(y))                      // []int
fmt.Println(reflect.TypeOf(y) == reflect.TypeOf(x)) // true

fmt.Println(reflect.ValueOf(x).Kind()) // slice
fmt.Println(reflect.ValueOf(y).Kind()) // slice

/**
  You cannot compare x == y
  invalid operation: cannot compare x == y
  (var x []int -> slice can only be compared to nil)
  compilerUndefinedOp
*/

// It is possible to create an int slice
// with zero length but greater capacity:

z := make([]int, 0, 10)		// Debugger: []int len: 0, cap: 10, []
fmt.Println(z, len(z), cap(z))  // [] 0 10

/**
Since its length is 0, we cannot directly
index into it, but we can append values to it:
*/

z := make([]int, 0, 10)
z = append(x,3,5,7); // []int len: 3, cap: 3, [3,5,7]

Same with nil-able declarations:

var v []int             // []]int len: 0, cap: 0, nil
v = append(v, 3, 5, 7)  // []int len: 3, cap: 3, [3,5,7]

```

Use `make` if you roughly know how big your slice needs to be but don't know
what values it will get.

The question is:

- whether you should specify a `nonzero length` or
- a `zero-length and a nonzero capacity` in the call to make.

There are three alternatives:

1. Slice as buffer &rarr; nonzero length
2. You know the size &rarr; specity the length and index into it. But if the set
   size was not big enough, you will get `panic`.
3. Or specify zero length, nonzero capacity and append to it. If the real size
   is smaller then there will be zero values at the end of the slice, if larger,
   your code will not panic.

#### Slice-slicing

See the following `slice expressions` that creates a slice form a slice:

```go
//            0  1  2  3  4
var x = []int{2, 3, 5, 7, 9}
var a = x[:3]   // 0 (incl) till 3 (excl) -> [2 3 5]
var b = x[2:]   // 2 (incl) till end  -> [5 7 9]
var c = x[1:4]  // 1 (incl) till 4 (excl)  -> [3 5 7]
var d = x[:] // all -> [2 3 5 7 9]
```

#### Sharing Memory

Important: Slices are not copies, they are references.

```go
var x = []int{2, 3, 5, 7, 9}
var b = x[0:2]
x[0] = 100

fmt.Println(x)  // [100 3 5 7 9]
fmt.Println(b)  // [100 3]
```

Many funny things happen:

```go
var x = []int{2, 3, 5, 7, 9}
var b = x[:2]

fmt.Println(cap(x), cap(b))   // 5 5

x[0] = 100
b = append(b, 30)

fmt.Println("x:", x)    // x: [100 3 30 7 9]
fmt.Println("b:", b)    // b: [100 3 30]
```

A more confusing example:

> Never `append` to a `slice` if you want to avoid surprises, or use the trick
> (third parameter with position) after the example below.

```go
x := make([]int, 0, 10)
x = append(x, 3, 5, 7, 9)
b := x[:2]
c := x[2:]

fmt.Println("x:", x)        // x: [3 5 7 9]
fmt.Println("b:", b)        // b: [3 5]
fmt.Println("c:", c)        // c: [7 9]

b = append(b, 20, 30, 40)
x = append(x, 11)
c = append(c, 13)

fmt.Println("x:", x)        // x: [3 5 20 30 13]
fmt.Println("b:", b)        // b: [3 5 20 30 13]
fmt.Println("c:", c)        // c: [20 30 13]
```

Notice the 3rd parameter in the slide expression: We limit the capacity of the
subslices to their length.

```go
x := make([]int, 0, 10)
x = append(x, 3, 5, 7, 9)

b := x[:2:2]          // <-- 3rd parameter
c := x[2:4:4]         // <-- 3rd parameter

fmt.Println("x:", x)  // x: [3 5 7 9]
fmt.Println("b:", b)  // b: [3 5]
fmt.Println("c:", c)  // c: [7 9]

b = append(b, 20, 30, 40)
x = append(x, 11)
c = append(c, 13)

fmt.Println("x:", x)  // x: [3 5 7 9 11]
fmt.Println("b:", b)  // b: [3 5 20 30 40]
fmt.Println("c:", c)  // c: [7 9 13]
```

#### Converting Array to Slices

You can take a slice from an Array too (same problems - memory sharing - see
above, using third parameter helps here too).

```go
x := [5]int{1, 3, 5, 7, 9}
b := x[:2:2]
c := x[2:4:4]

fmt.Println("x:", x)          // x: [1 3 5 7 9]
fmt.Println("b:", b)          // b: [1 3]
fmt.Println("c:", c)          // c: [5 7]

b = append(b, 20, 30, 40)
//	x = append(x, 11)  // you cannot append to an array
c = append(c, 13)

fmt.Println("x:", x)          // x: [1 3 5 7 9]
fmt.Println("b:", b)          // b: [1 3 20 30 40]
fmt.Println("c:", c)          // c: [5 7 13]
```

without the 3rd param you will get surprising results (like it is the case with
slices). The same example above this time without the 3rd param:

```go
x := [5]int{1, 3, 5, 7, 9}
b := x[:2]
c := x[2:4]

fmt.Println("x:", x)            // x: [1 3 5 7 9]
fmt.Println("b:", b)            // b: [1 3]
fmt.Println("c:", c)            // c: [5 7]

b = append(b, 20, 30, 40)
//	x = append(x, 11)  // you cannot append to an array
c = append(c, 13)

fmt.Println("x:", x)            // x: [1 3 20 30 13]
fmt.Println("b:", b)            // b: [1 3 20 30 13]
fmt.Println("c:", c)            // c: [20 30 13]
```

#### `copy` helps you to avoid memory sharing problems

##### Same size:

```go
a := []int{1, 3, 5, 7, 9}
b := make([]int, 5)

x := copy(b, a)         // (destination <- source)

fmt.Println("a:", a)    // a: [1 3 5 7 9] source
fmt.Println("b:", b)    // b: [1 3 5 7 9] destination
fmt.Println("x:", x)    // x: 5 (number of copied elems)
```

##### Smaller size: from the beginning of source array

```go
a := []int{1, 3, 5, 7, 9}
b := make([]int, 2)

x := copy(b, a)

fmt.Println("a:", a)    // a: [1 3 5 7 9] source
fmt.Println("b:", b)    // b: [1 3] destination
fmt.Println("x:", x)    // x: 2 (number of copied elems)
```

##### Bigger size: zero values at the end

```go
a := []int{1, 3, 5, 7, 9}
b := make([]int, 7)

x := copy(b, a)

fmt.Println("a:", a)    // a: [1 3 5 7 9] source
fmt.Println("b:", b)    // b: [1 3 5 7 9 0 0] destination
fmt.Println("x:", x)    // x: 5 (number of copied elems)
```

##### From anywhere of the source slice

```go
a := []int{1, 3, 5, 7, 9}
b := make([]int, 2)

x := copy(b, a[3:])

fmt.Println("a:", a)    // a: [1 3 5 7 9]
fmt.Println("b:", b)    // b: [7 9]
fmt.Println("x:", x)    // x: 2
```

##### From overlapping sections of the source slice to the source slice (copy and replace within)

Try out to explain it :)

```go
a := []int{1, 3, 5, 7, 9, 11, 13}

x := copy(a[:3], a[3:])

fmt.Println("a:", a)  // a: [7 9 11 7 9 11 13]
fmt.Println("x:", x)  // x: 3
```

Explanation:

```text
1  3  5  replace with 7 9 11 13
different size therefore sub-result => 7 9 11

0  1  2  3  4  5  6
1  3  5  7  9  11 13
7  9  11 7  9  11 13  <- result
```

Second Example:

```go
a := []int{1, 3, 5, 7, 9, 11, 13}

x := copy(a[4:], a[1:4])

fmt.Println("a:", a) // [1 3 5 7 3 5 7]
fmt.Println("x:", x) // 3
```

### Strings, Runes, Bytes

Strings (re-assignable but immutable) have no capacity only length, written
within double quotes. Runes has no length, written within single quotes.

```go
str := "Hello World 🫢"
r := 'h'

fmt.Println(str, len(str), r, string(r))
// Hello World 11 104 h
```

```go
var a string = "Hi there"
var b byte = a[6]
var x string = a[3:]

fmt.Println(a, b, x) // Hi there 114 there
```

#### Conversion

In general: Take care of the length!

```go
str := "Oh 🫢"        // Smiley is 4 bytes
fmt.Println(len(str)) // 7
```

#### Strings to Slices

> Slices of runes -> uncommon

```go
var str string = "Oh 🫢"
var x []byte = []byte(str)
var y []rune = []rune(str)
fmt.Println(x, y)

  // outputs
  // [79 104 32 240 159 171 162] [79 104 32 129762]

  // UTF8: little-endian versus big-endian -> no problems here

```

### Maps

The zero value for a `map` is `nil`. Maps are not comparable. You can only check
if they are equal to `nil`.

`len` is OK. Key can be any comparable type.

`map[keyType]valueType:`

If we want to read a value of a non-existing key, we get its zero value (e.g. if
`int` &rarr; 0).

> you can use ++ increment operator to increment the numeric value for a map
> key. (&rarr; Example)

```go
var myFirstMap = map[string]int{
  "foo": 0,  // Attention: Comma is always required
}

var mySecondMap = map[string][]string{
  "foo": {"one", "two"},
  "bar": {"one", "two", "three"}, // Attention: Comma required
}

fmt.Println(myFirstMap, mySecondMap)
// map[foo:0] map[bar:[one two three] foo:[one two]]
```

#### With `make`

```go
	myMap1 := make(map[int][]string, 10)
	myMap2 := make(map[string]int, 10)

	myMap1[13] = []string{"Hi", "there"}
	myMap2["foo"] = 4716

	fmt.Println(myMap1, myMap2)

  // map[13:[Hi there]]
  // map[foo:4716]

```

#### `ok` &rarr; Comma `ok` idiom

It can be any literal, must not be `ok` only.

```go
myMap1 := map[string]int{
  "foo": 7,
  "bar": 11,
}

value, ok := myMap1["foo"]
fmt.Println(myMap1, value, ok)

value, ok = myMap1["bar"]
fmt.Println(myMap1, value, ok)

value, ok = myMap1["what"]
fmt.Println(myMap1, value, ok)

// outputs
// map[bar:11 foo:7] 7 true
// map[bar:11 foo:7] 11 true
// map[bar:11 foo:7] 0 false

```

#### Deleting

```go
	myMap1 := map[string]int{
		"foo": 7,
		"bar": 11,
		"baz": 2014,
	}

	delete(myMap1, "baz")

	fmt.Println(myMap1) // map[bar:11 foo:7]
```

#### Maps as Sets

```go
//
```
