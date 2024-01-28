# Church Encoding in Go

***Note**. This is a piece of code written just for fun.*

## Lambda Term

It is amazing that in Go it is allowed to declare strange
recursive types, which one is `Term`:
```go
package church

// Term is an abstraction of a function
// that can be applied any number of
// times to itself. It is the base
// concept in lambda calculus. There
// will be aliases for this type for
// better readability.
type Term func(Term) Term
```
It means that we can implement pure lambda calculus. 'purity'
means following rules of code writing:
1. Use only `Term` declarations/definitions and function application.
2. All types except `Term`, constructions like `if` and `for` are prohibited.

So, we only can use `return` and `func` keywords, but
(as you can see) they are quite sufficient for implementation
of basics like integers, booleans and it's operations, and it
shows incredible beauty of lambda calculus.

## Testing

There are some files that break our rules: `dirty.go` and `binary.go`.
First contains tricky mappings from lambda abstractions to Go built-in
types for tests, second one exists only for syntax sugaring (in tests).

## Reading order

If you want to explore this funny code, I have an order for you:
1. `basic.go` — obviously
2. `numeral.go` — until you see something new and awkward
3. `boolean.go` — nice guys
4. `numeral.go` — before division
5. `compare.go` — they will save us
6. `dirty.go` — disgusting hacks
7. `numeral.go` — the main boss, division and mod
8. `klop.go` — are you sure?
