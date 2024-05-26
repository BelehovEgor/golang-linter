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

* [ ] Variable usage with mutex check - possible with tags ad CFG. Highlighting: **WARNING** (cause of accuracy?)
* [ ] Null interface call check - possible with CFG. Highlighting: **WARNING** (cause of accuracy?)
* [x] Cast interface to type check - possible with AST (golang-linter should check that first var is interface).
  Highlighting: **WARNING** (cause of accuracy?). !ATTENTION! `res = a.(b)` valid for checker if `b` implement `a`. Need
  TypeCheck? UPD: Detected by existing linters

### [UBER](https://github.com/uber-go/guide/blob/master/style.md)

#### Performance

1. [StrConv instead fmt](https://github.com/uber-go/guide/blob/master/style.md#prefer-strconv-over-fmt) - possible with
   AST. Highlighting: **ERROR**.
   Grade: 😍.
   Status: test
2. [String to byte repeating avoid](https://github.com/uber-go/guide/blob/master/style.md#avoid-repeated-string-to-byte-conversions) -
   possible with AST (we just want to find "str to byte" in cycle)?
   Highlighting: **WARNING** (cause of accuracy?).
   Status: test
3. [Map capacity](https://github.com/uber-go/guide/blob/master/style.md#prefer-specifying-container-capacity) -
   possible with AST.
   Highlighting: **WARNING** (cause of "try")
   Grade: 😍.
   Status: not implemented

#### Guidelines

1. [Pointers to Interfaces](https://github.com/uber-go/guide/blob/master/style.md#pointers-to-interfaces) - possible
   with AST (pointer variable decl + assignment with new(...) or we need a TypeCheck info for resolve expressions?)?
   Highlighting: **WARNING** (cause of "usually").
   Status: not implemented
2. [Verify Interface Compliance](https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance) -
   possible with AST (just find custom interface and decl with casting it to null)?
   Highlighting: **ERROR** (or it is not important?).
   Status: not implemented
3. [Receivers and Interfaces](https://github.com/uber-go/guide/blob/master/style.md#receivers-and-interfaces) - possible
   with TypeCheck info (we should resolve expression type?)?.
   Highlighting: **ERROR**.
   Status: not implemented
4. [Zero-value Mutexes are Valid](https://github.com/uber-go/guide/blob/master/style.md#zero-value-mutexes-are-valid) -
   possible with AST.
   Highlighting: **ERROR**.
   Grade: 😍.
   Status: test
5. [Copy Slices and Maps at Boundaries](https://github.com/uber-go/guide/blob/master/style.md#copy-slices-and-maps-at-boundaries) -
   possible with ??? (I think with AST we can check fact that input and output variables was copied or not, but if we
   create
   copies in some call - AST useless, it should be CFG analysis).
   Highlighting: **WARNING** (cause of accuracy?)
   Status: not implemented
6. [Defer to Clean Up](https://github.com/uber-go/guide/blob/master/style.md#defer-to-clean-up) -
   possible with AST? We can just find usages of Unlock and Close operation outside `defer`.
   Highlighting: **WARNING** (cause of accuracy?).
   Grade: 😍.
   Status: draft
7. [Channel Size is One or None](https://github.com/uber-go/guide/blob/master/style.md#channel-size-is-one-or-none) -
   possible with AST. Highlighting: **WARNING** (cause of "usually").
   Grade: 😍.
   Status: test
8. [Start Enums at One](https://github.com/uber-go/guide/blob/master/style.md#start-enums-at-one) -
   possible with AST. Highlighting: **WARNING** (cause of "usually").
   Grade: 😍.
   Status: test
9. [Use "time" to handle time](https://github.com/uber-go/guide/blob/master/style.md#use-time-to-handle-time) -
   impossible. We can't detect places where `Int/Float...` use like `Time`. We can try to detect write `Time.time` to
   Json without specific tag, but it is doubtful.
10. [Errors](https://github.com/uber-go/guide/blob/master/style.md#errors) - possible with AST. Highlighting:
    **WARNING** (cause of accuracy). We can check error.New() outside VarDecl and fmt.Errorf outside Struct, but it may
    be useless. Grade: 😐
11. [Error wrapping](https://github.com/uber-go/guide/blob/master/style.md#error-wrapping) - I don't understand
    example. I think they mean just apply context to error if possible. It looks like impossible for detection.
12. [Error naming](https://github.com/uber-go/guide/blob/master/style.md#error-naming) - Possible with AST.
    Highlighting:
    **Error** (because we can?).
    Grade: 😍.
13. [Handle Errors Once](https://github.com/uber-go/guide/blob/master/style.md#handle-errors-once) - Possible with AST
    but better with CFG. All we want is check log error with return in one scope? Highlighting: **WARNING** (cause of
    accuracy).
    Grade: 😐.
14. [Handle Type Assertion Failures](https://github.com/uber-go/guide/blob/master/style.md#handle-type-assertion-failures) -
    possible with AST. Highlighting: **WARNING** (cause of optional?).
    Grade: 😍.
    Status: test
15. [Don't Panic](https://github.com/uber-go/guide/blob/master/style.md#dont-panic) - possible with AST. Highlighting:
    **Error**.
    Grade: 😍.
    Status: test
16. [Use go.uber.org/atomic](https://github.com/uber-go/guide/blob/master/style.md#use-gouberorgatomic) - possible with
    AST. We need detect types from [this](https://pkg.go.dev/go.uber.org/atomic) and suggest to use uber atomic.
    Highlighting:
    **WARNING** (cause of optional?).
    Grade: 😐-😍.
17. [Avoid Mutable Globals](https://github.com/uber-go/guide/blob/master/style.md#avoid-mutable-globals) - possible with
    AST? Detect global `var` declaration is enough. Highlighting: **WARNING** (cause of "usually").
    Grade: 😍.
    Status: test
18. [Avoid Embedding Types in Public Structs](https://github.com/uber-go/guide/blob/master/style.md#avoid-embedding-types-in-public-structs) -
    Possible with AST but a more complex than other AST tasks. We need to find structures that use other structures from
    our project and check that property declared? I think I'm wrong.
    Highlighting: **WARNING** (cause of accuracy).
    Grade: 😐
19. [Avoid Using Built-In Names](https://github.com/uber-go/guide/blob/master/style.md#avoid-using-built-in-names) -
    possible with AST? I don't know how look go parser rules, but it may be very easy.
    Highlighting: **Error**.
    Grade: 😍.
    Status: **may be ready in default linters.**
20. [Avoid init()](https://github.com/uber-go/guide/blob/master/style.md#avoid-init) -
    possible with AST. Highlighting: **WARNING** (cause of optional?).
    Grade: 😍.
21. [Exit in Main](https://github.com/uber-go/guide/blob/master/style.md#exit-in-main) -
    possible with AST. Highlighting: **Error**.
    Grade: 😍.
22. [Exit Main Once](https://github.com/uber-go/guide/blob/master/style.md#exit-once)
    possible with AST. Highlighting: **WARNING** (cause of "if possible").
    Grade: 😍.
23. [Use field tags in marshaled structs](https://github.com/uber-go/guide/blob/master/style.md#use-field-tags-in-marshaled-structs) -
    for this task we need to get types (struct) of variables and find these structs. Go ast/cfg hold this data?
    Highlighting: **WARNING** (cause of accuracy).
    Grade: 😐.
24. [Don't fire-and-forget goroutines](https://github.com/uber-go/guide/blob/master/style.md#dont-fire-and-forget-goroutines) -
    may be possible with AST, but I think it is confusing task.
    Highlighting: **WARNING** (cause of accuracy).
    Grade: 😐.
25. [No goroutines in init()](https://github.com/uber-go/guide/blob/master/style.md#no-goroutines-in-init) -
    possible with AST. Highlighting: **Error**.
    Grade: 😍.

#### Style

Highlighting: **WARNING** (cause of it is style?).

1. [Avoid overly long lines](https://github.com/uber-go/guide/blob/master/style.md#avoid-overly-long-lines) - 99 chars <
   IntelliJ default limit for line. Possible with readLines)
   Grade: 😍.
   Status: **need to check default linter!**
2. [Be Consistent](https://github.com/uber-go/guide/blob/master/style.md#be-consistent) -
   general recommendation
3. [Group Similar Declarations](https://github.com/uber-go/guide/blob/master/style.md#group-similar-declarations) -
   partly possible with AST (last subrule).
   Grade: 😍.
4. [Import Group Ordering](https://github.com/uber-go/guide/blob/master/style.md#import-group-ordering) -
   possible with AST. Already in golangci-lint?
   Grade: 😐.
   Status: **need to check default linter!**
5. [Package Names](https://github.com/uber-go/guide/blob/master/style.md#package-names) -
   possible with AST. Already in golangci-lint (partly).
   Grade: 😍.
   Status: **need to check default linter!**
6. [Function Names](https://github.com/uber-go/guide/blob/master/style.md#function-names) -
   possible with AST. Already in golangci-lint?
   Grade: 😍.
   Status: **need to check default linter!**
7. [Import Aliasing](https://github.com/uber-go/guide/blob/master/style.md#import-aliasing) -
   possible with AST, but we need to library sources.  
   Grade: 😐.
8. [Function Grouping and Ordering](https://github.com/uber-go/guide/blob/master/style.md#function-grouping-and-ordering) -
   I think it is possible with AST, because we need only approximate order
   of intra-procedure execution, but CFG more correct.
   Grade: 😐.
9. [Reduce Nesting](https://github.com/uber-go/guide/blob/master/style.md#reduce-nesting) -
   partly possible with CFG.
   Grade: 😐.
10. [Unnecessary Else](https://github.com/uber-go/guide/blob/master/style.md#unnecessary-else) -
    possible with CFG (mb AST).
    Grade: 😐.
11. [Top-level Variable Declarations](https://github.com/uber-go/guide/blob/master/style.md#top-level-variable-declarations) -
    potential possible with AST, but I think we can't support this right now - for stable work we should check type of
    any expression or identifier.
    Grade: ☹️.
12. [Prefix Unexported Globals with _](https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_) -
    possible with AST, but we need to analyze all project files for collect `import`.
    Grade: 😐.
13. 