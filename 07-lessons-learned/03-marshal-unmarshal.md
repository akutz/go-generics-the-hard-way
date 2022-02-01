# Marshal/unmarshal

One of the early, incorrect takes I had on Go generics was based on some early [boxing benchmarks](../06-benchmarks/01-boxing.md) -- the performance improvement is phenomanal! I thought _Generics will benefit any function that takes an empty interface...that's code that uses an marshal/unmarshal pattern!_

Except I was wrong. The problem is that at the end of the day, that code is very likely using stdlib or some other library to actually marshal/unmarshal the object/data to/from its wire/object form. And until _that_ code takes `T`, boxing will still occur under the covers. And since once the value has been boxed it does not need to be boxed again, if boxing occurs once, the benefit of eliminating it everywhere else above the stack is negated.

To see what I mean, take a look at the `encoding/json.Marshal` and `Unmarshal` functions in 1.18beta2:

* [`func Marshal(v any) ([]byte, error)`](https://pkg.go.dev/encoding/json@go1.18beta2#Marshal)
* [`func Unmarshal(data []byte, v any) error`](https://pkg.go.dev/encoding/json@go1.18beta2#Unmarshal)

Both functions still take the empty interface (or its type alias, the `any` identifier). This means if at any point code that has been painstakingly refactored to use generic types to eliminate boxing calls into one of the above functions -- you've been boxed.

I am not entirely sure these functions _can_ be generic, but unless they are, they impact everything above them.

---

Next: [Impact to build times & file sizes](./04-builds.md)
