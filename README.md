# golang-linter

- errcheck - Check unchecked errors in Go code. These unchecked errors can be critical bugs in some cases.
- goimports - Check import statements are formatted according to the 'goimport' command. Reformat imports in autofix
  mode.
- revive - Fast, configurable, extensible, flexible, and beautiful linter for Go. Drop-in replacement of
  golint. [what can do](https://revive.run/r)
- govet - Vet examines Go source code and reports suspicious constructs. It is roughly the same as 'go vet' and uses its
  passes.
- staticcheck - It's a set of rules from [staticcheck](https://staticcheck.io/docs/configuration/options/#checks).

## Checks:

### Custom PA

* Variable usage with mutex check - possible with tags ad CFG. Highlighting: **WARNING** (cause of accuracy?)
* Null interface call check - possible with CFG. Highlighting: **WARNING** (cause of accuracy?)
* Cast interface to type check - possible with AST (golang-linter should check that first var is interface).
  Highlighting: **WARNING** (cause of accuracy?). !ATTENTION! `res = a.(b)` valid for checker if `b` implement `a`. Need
  TypeCheck? UPD: intllijIdea detect this rule

### [UBER](https://github.com/uber-go/guide/blob/master/style.md)

#### Performance

1. [StrConv instead fmt](https://github.com/uber-go/guide/blob/master/style.md#prefer-strconv-over-fmt) - possible with
   AST. Highlighting: **ERROR**. Grade: 😍
2. [String to byte repeating avoid](https://github.com/uber-go/guide/blob/master/style.md#avoid-repeated-string-to-byte-conversions) -
   possible with AST (we just want to find "str to byte" in cycle)? Highlighting: **WARNING** (cause of accuracy?)
3. [Map capacity](https://github.com/uber-go/guide/blob/master/style.md#avoid-repeated-string-to-byte-conversions) -
   possible with AST. Highlighting: **WARNING** (cause of "try") Grade: 😍

#### Guidelines

1. [Pointers to Interfaces](https://github.com/uber-go/guide/blob/master/style.md#pointers-to-interfaces) - possible
   with AST (pointer variable decl + assignment with new(...) or we need a TypeCheck info for resolve expressions?)?
   Highlighting: **WARNING** (cause of "usually")
2. [Verify Interface Compliance](https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance) -
   possible with AST (just find custom interface and decl with casting it to null)?
   Highlighting: **ERROR** (or it is not important?)
3. [Receivers and Interfaces](https://github.com/uber-go/guide/blob/master/style.md#receivers-and-interfaces) - possible
   with TypeCheck info (we should resolve expression type?)?. Highlighting: **ERROR**
4. [Zero-value Mutexes are Valid](https://github.com/uber-go/guide/blob/master/style.md#zero-value-mutexes-are-valid) -
   possible with AST. Highlighting: **ERROR**. Grade: 😍
5. [Copy Slices and Maps at Boundaries](https://github.com/uber-go/guide/blob/master/style.md#copy-slices-and-maps-at-boundaries) -
   possible with ??? (I think with AST we can check fact that input and output variables was copied or not, but if we
   create
   copies in some call - AST useless, it should be CFG analysis). Highlighting: **WARNING** (cause of accuracy?)
6. [Defer to Clean Up](https://github.com/uber-go/guide/blob/master/style.md#copy-slices-and-maps-at-boundaries) -
   possible with AST? We can just find usages of Unlock and Close operation outside `deffer`. Highlighting: **WARNING
   ** (cause of accuracy?). Grade: 😍
7. [Channel Size is One or None](https://github.com/uber-go/guide/blob/master/style.md#channel-size-is-one-or-none) -
   possible with AST. Highlighting: **WARNING** (cause of "usually"). Grade: 😍
8. [Start Enums at One](https://github.com/uber-go/guide/blob/master/style.md#start-enums-at-one) -
   possible with AST. Highlighting: **WARNING** (cause of "usually"). Grade: 😍
9. [Avoid Mutable Globals](https://github.com/uber-go/guide/blob/master/style.md#avoid-mutable-globals) - possible with
   AST? Detect global `var` declaration is enough. Highlighting: **WARNING** (cause of "usually"). Grade: 😍