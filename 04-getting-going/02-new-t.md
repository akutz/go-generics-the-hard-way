# Declaring a new instance of `T` on the heap with `new`

Consider the following example ([Golang playground](https://gotipplay.golang.org/p/Y_n2i8O0YQ1)):

```golang
func NewT[T any]() *T {
	var t *T
	return t
}

func main() {
	fmt.Println(NewT[int]())
}
```

The above program will print `<nil>` because `NewT[T] *T` returns a `*T` with a `nil` value since there was no instance of a `T` allocated. Now consider ([Golang playground](https://gotipplay.golang.org/p/T_vKX69p0l_0)):

```golang
func NewT[T any]() *T {
	return new(T)
}

func main() {
	fmt.Println(NewT[int]())
}
```

Now the output will be a memory address that references a new instance of an `int` value. This means we can also clean up our example from the previous page so there is no longer a need to assign `sum` to the default value of `T` since `new(T)` allocates that for us ([Golang playground](https://gotipplay.golang.org/p/LysgqC56CBU)):

```golang
func Sum[T Numeric](args ...T) T {
	sum := new(T)
	for i := 0; i < len(args); i++ {
		*sum += args[i]
	}
	return *sum
}
```

And just like with `var t`, a new `T` allocated with `new` can still be inlined and optimized to the stack instead of escaping to the heap:

```bash
$ docker run -it --rm go-generics-the-hard-way \
  go run -gcflags "-m" ./04-getting-going/02-new-t/stack
# go-generics-the-hard-way/04-getting-going/02-new-t/stack
04-getting-going/02-new-t/stack/main.go:32:6: can inline Sum[go.shape.int_0]
04-getting-going/02-new-t/stack/main.go:41:17: inlining call to Sum[go.shape.int_0]
04-getting-going/02-new-t/stack/main.go:41:13: inlining call to fmt.Println
04-getting-going/02-new-t/stack/main.go:41:13: ... argument does not escape
04-getting-going/02-new-t/stack/main.go:41:17: ~R0 escapes to heap
04-getting-going/02-new-t/stack/main.go:41:17: ... argument does not escape
04-getting-going/02-new-t/stack/main.go:41:17: new(go.shape.int_0) does not escape
04-getting-going/02-new-t/stack/main.go:33:12: new(go.shape.int_0) does not escape
6
```

Still, regardless of how a new `T` is allocated, there _is_ a ticking timebomb when discussing `*T` that will be covered at the end of this section.

---

Next: [Structs](./03-structs.md)
