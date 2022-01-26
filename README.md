# Go generics the hard way

I started using Go back around 2015 and was immediately surprised by the lack of a generic type system. Sure, the empty `interface{}` existed, but that was hardly the same. At first I thought I ~~wanted~~ _needed_ generics in Go, but over time I began appreciating the simplicity of the language. Therefore I was ambivalent at best when I learned of discussions to introduce generics in Go 2.0, and once the timetable was accelerated to 1.18, I decided it was time to dig into the proposal.

After a while, I gained an appreciation for how generics are implemented with the same elegance as Golang itself, and this moved me to share my experience. _Go generics the hard way_ is a culmination of the time I spent playing with this new feature and provides a hands-on approach to learning all about generics in Go.

* [**Labs**](#labs): a hands-on approach to learning Go generics
* [**FAQ**](#FAQ): answers to some of the most frequently asked questions regarding Go generics
* [**Links**](#links): links to related reference material and projects that use generics

## Labs

1. [**Prerequisites**](./01-prereqs/): how to install the prerequisites required to run the examples in this repository
2. [**Hello world**](./02-hello-world/): a simple example using generics
3. [**Getting started**](./03-getting-started): an introduction to go generics
4. [**Getting going**](./04-getting-going): basic concepts explored
5. [**Internals**](./05-internals/): how generics are implemented in golang
6. [**Benchmarks**](./06-benchmarks/): basic benchmarks for common patterns using generics
7. [**Lessons learned**](./07-lessons-learned/): lessons learned from digging into generics

## FAQ

* [**How are you using generics in the Go playground?**](#how-are-you-using-generics-in-the-go-playground)
* [**What is `T`?**](#what-is-t)
* [**What is this `any` I keep seeing everywhere?**](#what-is-this-any-i-keep-seeing-everywhere)
* [**What does the tilde `~` do?**](#what-does-the-tilde-do)
* [**Do Go generics use _type erasure_?**](#do-go-generics-use-type-erasure)


### How are you using generics in the Go playground?
We can use [the Go playground in “Go dev branch” mode](https://go.dev/play/?v=gotip) to edit and run your program with generics.

### What is `T`?

The symbol `T` is often used when discussing generic types because `T` is the first letter of the word _**t**ype_. That is really all there is too it. Just like `x` or `i` are often the go-to variable names for loops, `T` is the go-to symbol for generic types.

For what is worth, `K` is often used when there is more than one generic type, ex. `T, K`.


### What is this `any` I keep seeing everywhere?

The word `any` is a new, [predeclared identifier](https://go.dev/ref/spec#Predeclared_identifiers) and is [equivalent to the empty interface in all ways](https://github.com/golang/go/blob/24239120bfbff9ebee8e8c344d9d3a8ce460b686/src/builtin/builtin.go#L94-L95). Simply put, writing and reading `any` is just more user friendly than `interface{}` :smiley:.


### What does the tilde `~` do?

The `~` symbol is used to express that `T` may be satisfied by a defined or named type directly or by a type definition that has the same, underlying type as another defined or named type. To learn more about type constraints and the `~` symbol, please refer to the section [_Tilde `~`_](./03-getting-started/06-tilde.md).


### Do Go generics use _type erasure_?

Generics in Go are not implemented with type erasure. Please jump to [_Internals_](./05-internals/) for more information.


## Links

### Additional reading

* [**Type parameter proposal**](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md): the accepted proposal for introdicing generics to go
* [**Getting started with generics**](https://go.dev/doc/tutorial/generics): a tutorial from the authors of go for getting started with generics
* [**Go language specification**](https://go.dev/ref/spec): the reference manual for the go language
* [**Go FAQ**](https://go.dev/doc/faq): frequently asked questions for go

### Projects using generics

* [**Controller-runtime**](https://gist.github.com/akutz/887fa677f2196c341d85595f14c6280b): a write-up and patchset for implementing conditions logic, patch helpers, and simple reconcilers using generics
* [**Go collections**](https://github.com/mikhailswift/go-collections): generic utility functions for dealing with collections in go
* [**go-generics-example**](https://github.com/mattn/go-generics-example): examples using generics

