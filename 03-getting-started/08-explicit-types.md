# Explicit types

While exploring _type inference_, we realized that up to this point we have not had to explicitly specify the type for calls to `Sum[T Numeric](...T) T`. Yet it _is_ possible to explicitly state what types to use with generic functions. In fact, this approach can be used to avoid compiler errors. For example, remember the example that used type inference and failed to compile ([Golang playground](https://gotipplay.golang.org/p/Jg9uCMglPC-))?

```golang
func main() {
	fmt.Println(Sum([]int{1, 2, 3}...))
	fmt.Println(Sum([]id{1, 2, 3}...))
	fmt.Println(Sum(1, 2, 3.0))
}
```

The above example failed with the following compiler error:

```bash
./prog.go:30:24: default type float64 of 3.0 does not match inferred type int for T
```

The Go compiler infered the type of `T` to be an `int` based on the first value to the function, `1`. However, if we explicitly assert the type this error can be avoided ([Golang playground](https://gotipplay.golang.org/p/TW6qmbRRZRa)):

```golang
func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum([]id{1, 2, 3}...))
	
	// Generic types can be specified explicitly by invoking a function
	// with the bracket notation and the list of types to use. Because
	// the Sum function only has a single, generic type -- "T" -- the
	// call "Sum[float64]" means that "T" will be replaced by "float64"
	// when compiling the code. Since the values "1" and "2" can both
	// be treated as "float64," the code is valid.
	fmt.Println(Sum[float64](1, 2, 3.0))
}

```

Now the program will print the expected output:

```bash
6
6
6
```

---

Next: [Multiple generic types](./09-multiple-generic-types.md)
