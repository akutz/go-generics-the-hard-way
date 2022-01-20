# The `any` constraint

Previously we illustrated how to write a function that can sum both `int` and `int64` types ([Go playground](https://gotipplay.golang.org/p/c6jzxJo0s7S)):

```golang
// Sum returns the sum of the provided arguments.
func Sum[T int | int64](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
```

But why limit the function only `int` and `int64`? One way to increase the number of supported types is by continuing to use the `|` operator:

```golang
func Sum[T int | int8 | int32 | int64](args ...T) T
```

However, this would become cumbersome if we wanted to include _any_ numeric types. Huh, _any_? Is there not some new `any` identifier in Go? Can we use it to rewrite `Sum[T](...T) T`? Let's try ([Go playground](https://gotipplay.golang.org/p/pjq15Sro_SQ)):

```golang
package main

import (
	"fmt"
)

// Sum returns the sum of the provided arguments.
func Sum[T any](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}...))
	fmt.Println(Sum([]int8{1, 2, 3}...))
	fmt.Println(Sum([]uint32{1, 2, 3}...))
	fmt.Println(Sum([]float64{1.1, 2.2, 3.3}...))
	fmt.Println(Sum([]complex128{1.1i, 2.2i, 3.3i}...))
}
```

Unfortunately the above program will fail to compile with the following error:

```bash
./prog.go:11:3: invalid operation: operator + not defined on sum (variable of type T constrained by any)
```

The `any` identifier is [equivalent to the empty interface in all ways](https://github.com/golang/go/blob/24239120bfbff9ebee8e8c344d9d3a8ce460b686/src/builtin/builtin.go#L94-L95). Thus `T` in the above example might not represent a type for which the addition operator is valid.

Instead we need a type that represents all, possible _numeric_ types...

---

Next: [Composite constraints](./05-composite-constraints.md)
