# Declaring a new instance of `T` with `var`

Let's take one more look at our old friend `Sum[T Numeric](...T) T`:

```golang
func Sum[T Numeric](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
```

The line `var sum T`:

* declares a new variable `sum`
* of type `T`
* on the stack

It would _also_ be possible to rewrite the function as ([Golang example](https://gotipplay.golang.org/p/ePMrDb_gfv1)):

```golang
func Sum[T Numeric](args ...T) T {
	var defaultT T
	var sum *T = &defaultT
	for i := 0; i < len(args); i++ {
		*sum += args[i]
	}
	return *sum
}
```

Now there are two, declared variables:

* `defaultT`
  * of type `T`
  * on the stack
* `sum`
  * of type `*T`
  * on the stack

It is necessary to declare `defaultT` in order to take the address of the empty value of type `T`, because otherwise `sum` would point to a `nil` value. And yes, pointers _can_ be on the stack. Go is able to optimize things pretty efficiently. To prove this simply take the above example and run it locally with the flags `-gcflags "-m"`:

```bash
$ docker run -it --rm go-generics-the-hard-way \
  go run -gcflags "-m" ./04-getting-going/01-var-t/stack
# go-generics-the-hard-way/04-getting-going/01-var-t/stack
04-getting-going/01-var-t/stack/main.go:16:6: can inline Sum[go.shape.int_0]
04-getting-going/01-var-t/stack/main.go:26:17: inlining call to Sum[go.shape.int_0]
04-getting-going/01-var-t/stack/main.go:26:13: inlining call to fmt.Println
04-getting-going/01-var-t/stack/main.go:26:13: ... argument does not escape
04-getting-going/01-var-t/stack/main.go:26:17: ~R0 escapes to heap
04-getting-going/01-var-t/stack/main.go:26:17: ... argument does not escape
6
```

Please note that 

If the function returned `*T` then `defaultT` would be moved to the heap ([Golang playground](https://gotipplay.golang.org/p/dk9fWDftW4Z)):

```bash
$ docker run -it --rm go-generics-the-hard-way \
  go run -gcflags "-m" ./04-getting-going/01-var-t/heap
# go-generics-the-hard-way/04-getting-going/01-var-t/heap
04-getting-going/01-var-t/heap/main.go:16:6: can inline Sum[go.shape.int_0]
04-getting-going/01-var-t/heap/main.go:26:19: inlining call to Sum[go.shape.int_0]
04-getting-going/01-var-t/heap/main.go:26:13: inlining call to fmt.Println
04-getting-going/01-var-t/heap/main.go:26:13: ... argument does not escape
04-getting-going/01-var-t/heap/main.go:26:14: *(~R0) escapes to heap
04-getting-going/01-var-t/heap/main.go:26:19: ... argument does not escape
04-getting-going/01-var-t/heap/main.go:17:6: moved to heap: defaultT
6
```

In other words, whether the `t` in `var t T` ends up moving/escaping to the heap is subject to all of the same rules as normal Go code. Speaking of the heap, `var t T` is only one way to declare a new instance of `T`...

---

Next: [`new(T)`](./02-new-t.md)
