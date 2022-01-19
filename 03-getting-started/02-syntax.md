# Syntax

The basic syntax for writing generic code can be explained by taking a look at an example from the previous page:

```golang
// Sum returns the sum of the provided arguments.
func Sum(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
```

The above function:

* takes zero to many `int` arguments
* returns an `int` value

In fact, there is no type a part of the function signature or body that is not `int`. So one way to rewrite the above function generically would be to let `T` represent `int` ([Go playground](https://gotipplay.golang.org/p/UkF5oXDeW_i)):

```golang
// Sum returns the sum of the provided arguments.
func Sum[T int](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
```

The above function's name is suceeded by `[T int]`, the basic form of the syntax for expressing generics in Go:

* the `[]` brackets are always used to define the generics
* the most basic pattern is `[<ID> <CONSTRAINT>]` where:
  * **`<ID>`**: is the symbol used to represent the generic type
  * **`<CONSTRAINT>`**: is the constraint that indicates which concrete types can be used

In the above example the generic type `T` is constrainted to `int`, which means it is not possible to call `Sum[T int](...T) T` with a list of `int64` values ([Golang playground](https://gotipplay.golang.org/p/8FCzhpVpPUc)):

```golang
func main() {
	fmt.Println(Sum([]int64{1, 2, 3}...))
}
```

Trying to run the above example will produce the following compiler error:

```bash
./prog.go:17:17: int64 does not implement int
```

The next page will review how to rewrite the above function so that more than one type can satisfy its constraints.

---

Next: [Constraints](./03-constraints.md)
