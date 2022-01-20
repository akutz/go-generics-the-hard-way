# Tilde `~`

The tilde `~` symbol is used to express a constraint that may be satisfied by:

* a defined or named type
* a [type definition](https://go.dev/ref/spec#Type_definitions) with the same underlying type as another defined or named type

For example, `~int` would match both of the following:

* `int`: the built-in `int` type
* `type Integer int`: a type definition that has the same, underlying type as `int`

For what it is worth, this works for `struct` types as well, but those will be covered later. For now, please consider the example from the previous page:

```golang
// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64 |
		complex64 | complex128
}

// Sum returns the sum of the provided arguments.
func Sum[T Numeric](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// id is a new type definition for an int64
type id int64

func main() {
	fmt.Println(Sum([]id{1, 2, 3}...))
}
```

As it is written, the example fails to compile because `id` does not satisfy the constraint `Numeric`. In fact, the error produced by the compiler gives a hint to the issue:

```bash
./prog.go:28:17: id does not implement Numeric (possibly missing ~ for int64 in constraint Numeric)
```

By prefixing the composited type `int64` with `~`, ex. `~int64`, the constraint can now be satisified by the defined type `int64` or a type definition that has an underlying type of `int64`. Now the program will produce the expected output ([Golang playground](https://gotipplay.golang.org/p/Lk720EthSS0)):

```bash
6
```


---

Next: [Type inference](./07-type-inference.md)
