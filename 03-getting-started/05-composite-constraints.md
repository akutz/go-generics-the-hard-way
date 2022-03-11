# Composite constraints

We know that if the `Sum[T](...T) T` function is going to be able to handle all, possible numeric types that:

* inlining all those possibilities would be poorly written code
* the `any` keyword will not work because the addition operator is invalid for the `any` type

Instead something _like_ `any` is needed but that is constrained to just the possible numeric types. The solution is to create a composite constraint using the following syntax:

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
```

Now the function `Sum[T Numeric](...T) T` can be used to sum values of any numeric type ([Go playground](https://gotipplay.golang.org/p/8YcJ9KkVZvg)):

```golang
func main() {
	fmt.Println(Sum([]int{1, 2, 3}...))
	fmt.Println(Sum([]int8{1, 2, 3}...))
	fmt.Println(Sum([]uint32{1, 2, 3}...))
	fmt.Println(Sum([]float64{1.1, 2.2, 3.3}...))
	fmt.Println(Sum([]complex128{1.1i, 2.2i, 3.3i}...))
}
```

The above program will compile and print the expected result:

```bash
6
6
6
6.6
(0+6.6i)
```

Yet...let's rexamine the following statement:

> the function `Sum[T Numeric](...T) T` can be used to sum values of any numeric type

Is this really _true_? What about a type definition with an underlying type of `int`? In fact, that will not work as `Sum[T Numeric](...T) T` is currently written ([Go playground](https://gotipplay.golang.org/p/YohUg7xdUIl)):

```golang
// id is a new type definition for an int64
type id int64

func main() {
	fmt.Println(Sum([]id{1, 2, 3}...))
}
```

In fact the above program fails to compile. The next section explores the tilde `~` operator and how it can be used to further expand how constraints are expressed.

---

Next: [Tilde `~`](./06-tilde.md)
