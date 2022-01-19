# What is a generic?

A generic is a symbol that represents one or more concrete types. Whereas variables are placeholders for referencing values, generics are placeholders for concrete types (there is a major difference which will be covered later). For example, the following code declares a variable and assigns it a value:

```golang
var x int   // declare "x"
x = 2       // assign "x" the value of "2"
```

A developer can now refer to the value `2` using the variable `x` and/or assign a different value to `x`. In other words, the variable `x`:

* can be used in any operation that takes a whole number
* can be assigned the value of any whole number

In the same way, a generic type:

* can be used in any operation in place of the concrete type(s) the generic type represents
* can be assigned any value that is valid for the concrete type(s) the generic type represents

The following function is a quick and easy way to return the sum of zero to many `int` values:

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

However, what if a developer wants to implement the same logic for `int64`? For `uint32`? Before generics it would likely mean:

* Using the empty interface and asserting type ([Go playground](https://go.dev/play/p/Q6G0Gdw6Bdx)),
* or multiple functions, likely with the intended type name as their suffix ([Go playground](https://go.dev/play/p/WySbIpSvFz6)).

What is needed is some way to rewrite the original `Sum` function in a _generic_ fashion so that it can:

* be friendly to developers writing code (take multiple types)
* provide compile-time type safety
* provide runtime type safety

And that is exactly what Golang generics provides. Please keep reading to learn how to write a generic version of the `Sum` function.

---

Next: [Syntax](./02-syntax.md)
