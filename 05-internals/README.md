# Internals

The _Getting started_ section of this repository defined generics as a pattern that:

* is friendly to developers writing code (take multiple types)
* provides compile-time type safety
* provides runtime type safety

With that in mind, this section explores how generics are implemented in Go by comparing it to Java and .NET. However:

* readers can draw their own conclusions regarding the developer-friendliness of generics in Go versus other languages
* and all of the aforementioned languages, including Go, provide compile-time type safety with generics

This section focus on _compiled_ code:

* [**Type erasure**](./01-type-erasure/): exploring type erasure
* [**Runtime type safety**](./02-runtime-type-safety/): enforcing generic type constraints at runtime
* [**Runtime instantiation**](./03-runtime-instantiation/): instantiating new generic types at runtime
* [**In summary**](./04-summary.md): a summary of how generics in Java, .NET, and Go stack up against one another

---

Next: [Type erasure](./01-type-erasure/)
