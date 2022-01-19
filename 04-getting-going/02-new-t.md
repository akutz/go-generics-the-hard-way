# Declaring a new instance of `T` on the heap with `new`

For all intents and purposes the following are semantically identical:

* `var t *T`
* `new(T)`

In both of the following examples, `NewT` does the same thing:

**`var t *T`** ([Golang playground](https://gotipplay.golang.org/p/FKP-SlsmhAg))

```golang
func NewT[T any]() *T {
	var t *T
	return t
}
```

**`new(T)`** ([Golang playground](https://gotipplay.golang.org/p/RfSW5Pq9mXD))

```golang
func NewT[T any]() *T {
	return new(T)
}
```

For even more proof, the following example illustrate that Go can still, when possible, optimize a `T` value to the stack even when its address was created by `new(T)`:

```bash
$ docker run -it --rm go-generics-the-hard-way \
  go run -gcflags "-m" ./04-getting-going/02-new-t/stack
# go-generics-the-hard-way/04-getting-going/02-new-t/stack
04-getting-going/02-new-t/stack/main.go:16:6: can inline Sum[go.shape.int_0]
04-getting-going/02-new-t/stack/main.go:27:17: inlining call to Sum[go.shape.int_0]
04-getting-going/02-new-t/stack/main.go:27:13: inlining call to fmt.Println
04-getting-going/02-new-t/stack/main.go:27:13: ... argument does not escape
04-getting-going/02-new-t/stack/main.go:27:17: ~R0 escapes to heap
04-getting-going/02-new-t/stack/main.go:27:17: ... argument does not escape
04-getting-going/02-new-t/stack/main.go:27:17: new(go.shape.int_0) does not escape
04-getting-going/02-new-t/stack/main.go:18:12: new(go.shape.int_0) does not escape
6
```

However, there _is_ a ticking timebomb when discussing `*T` that will be covered at the end of this section.

---

Next: [Structs](./03-structs.md)
