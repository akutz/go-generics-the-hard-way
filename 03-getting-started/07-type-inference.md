# Type inference

Wow, we have come so far already! We have created:

* a generic function
* a generic constraint
* a new type definition that satifies the constraint

Now we can use a single function to sum various types of numeric values ([Golang playground](https://gotipplay.golang.org/p/_PLH2reLokI)):

```golang
func main() {
	fmt.Println(Sum([]int{1, 2, 3}...))
	fmt.Println(Sum([]id{1, 2, 3}...))
	fmt.Println(Sum(1, 2, 3.0))
}
```

:exclamation::exclamation: Except instead of the expected output...

```bash
6
6
6
```

...the program fails to run with the following compiler error :exclamation::exclamation::

```bash
./prog.go:29:24: default type float64 of 3.0 does not match inferred type int for T
```

This error occurred because the Go compiler sees `Sum[T Numeric](...T) T` and needs to determine what concrete type should replace `T`.

In fact, have you noticed we have actually not had to specify the type for `T` in all our calls to `Sum[T Numeric](...T) T` on this page or any of the previous ones? This is because of a feature in the Go compiler known as, you guessed it, _type inference_.

The Go compiler attempts to infer the intended, concrete types for a generic function from its provided arguments by:

* looking at the first instance of a generic type
* determining the concrete type for the generic type
* ensuring all other instances of the generic type share the same, concrete type

That is a _**gross**_ oversimplification. The exact process by which type inference occurs is rather [in-depth](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#Type-inference), and you _are_ encouraged to read it. However, in the end the following key, takeaways will serve you well 90% of the time:

* Type inference is a convenience feature
* The Go compiler tries _really_ hard to infer the intended types, but it does not always work when you think it should
* If you are not sure why something written generically is not working, try providing the types explicitly

For that matter, how _does_ one specify types explicitly? 

---

Next: [Explicit types](./08-explicit-types.md)
