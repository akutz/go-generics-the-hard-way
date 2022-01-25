# Constraints

In the example on the previous page we tried to sum a series of 64-bit integers with the generic `Sum[T int](...T) T` function ([Go playground](https://gotipplay.golang.org/p/8FCzhpVpPUc)):

```golang
package main

import (
	"fmt"
)

// Sum returns the sum of the provided arguments.
func Sum[T int](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int64{1, 2, 3}...))
}
```

However, when trying to run this example the compiler balks with the following error:

```bash
./prog.go:17:17: int64 does not implement int
```

This occurs because while the function `Sum[T int](...T) T` is written generically, the type of `T` is constrained to be only an `int`. The function needs to be rewritten so it can accept an `int` _or_ an `int64`:

```golang
func Sum[T int | int64](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}...))
	fmt.Println(Sum([]int64{1, 2, 3}...))
}
```

Using the `|` operator, the constraints for `T` can be expressed such that `T` can be satisfied by an `int` or `int64`. Now when the example is run, it produces the expected output ([Golang playground](https://gotipplay.golang.org/p/c6jzxJo0s7S)):

```bash
6
6
```

---

:sparkles: Wait a minute, isn't the `|` operator in Go supposed to be the [_bitwise_](https://go.dev/ref/spec#Arithmetic_operators) `OR` operation and `||` the [_conditional_](https://go.dev/ref/spec#Logical_operators) `OR`? 

You are not wrong. However, it seems that the type parameter proposal elected to follow the lead of languages like [PHP](https://en.wikipedia.org/wiki/Union_type#PHP), [Python](https://en.wikipedia.org/wiki/Union_type#Python), and [TypeScript](https://en.wikipedia.org/wiki/Union_type#TypeScript), which all use the `|` operator to express a union type. Indeed `T | K` is referred to as a [_union constraint_](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#union-constraint-element) in the type parameter proposal.

---

But what about wanting a `Sum` function that can handle _any_ numeric value? While it is possible to just list them all in `Sum[T int | int32 | int64 |...`, this would result in difficult-to-read code. Luckily there is a better way...

---

Next: [The any constraint](./04-the-any-constraint.md)
