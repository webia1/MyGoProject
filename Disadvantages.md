# Negative Sides of Go

- Generics from v.1.18
- No Function/Method Overloading
- No Inheritance (weird extending rules)
- It has function types and closures but functional programming is not possible
- Not for numerical computing
- `:=` Re-Assigning can cause `Shadowing Variables`
- No checks for unused constants
  - constants are not constants like in JS/TS
    - they are a kind of compile time calculations
- You cannot debug, if it does not build/run
- Unvisible, unintuitiv rules

  - Initial letter uppercase -> Public function
    - like export in JS/TS
  - `er` at the end of the word for interfaces ??
  - filename musst be package name, all lowercase

- Some ridiculous rules
  - error strings should not be capitalized (ST1005)go-staticcheck
