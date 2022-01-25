# Structs

So far all of the examples in this section have been focused on functions, but struct types can utilize generics as well. For example, the function `PrintIDAndSum[T ~string, K Numeric](T, func(...K) K)` can be decomposed into a struct that:

* has fields for the:
  * ID
  * numeric values
  * sum function
* has a function for printing the ID and sum

Okay, let's start by defining a struct that as a field named `ID` which can be a string value or a type definition with an underlying type of `string`:

```golang
// Ledger is an identifiable, financial record.
type Ledger[T ~string] struct {

	// ID identifies the ledger.
	ID T
}
```

The example demonstrates that defining structs with generic fields uses the same syntax in between the brackets as when defining generic functions. Let's go ahead and add the other fields:

```golang
// Ledger is an identifiable, financial record.
type Ledger[T ~string, K Numeric] struct {

	// ID identifies the ledger.
	ID T

	// Amounts is a list of monies associated with this ledger.
	Amounts []K

	// SumFn is a function that can be used to sum the amounts
	// in this ledger.
	SumFn SumFn[K]
}
```

So far I think this should make sense as it does not differ _too_ much from what we have already covered. However, the next step _is_ new -- defining a function on a generic struct. One of the use cases is:

> has a function for printing the ID and sum

Normally a function would be defined on `Ledger[T, K]` like so:

```golang
func (l Ledger) PrintIDAndSum() {}
```

However, because `Ledger` has generic types, their symbols must be included in the function receiver:

```golang
// PrintIDAndSum emits the ID of the ledger and a sum of its amounts on a
// single line to stdout.
func (l Ledger[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}
```

Notice the constraints do not need to be included, only the symbols for the constraints. With all of this in place, the same example from the previous page can be rewritten using the new type `Ledger[T, K]` ([Golang playground](https://gotipplay.golang.org/p/9GpJbR897Pr)):

```golang
func main() {
	Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	}.PrintIDAndSum()
}
```

The above program produces the expected output:

```bash
acct-1 has a sum of 6
```

Please note that all of the types must be supplied explicitly. This is because type inference is a convenience feature for functions and does not apply to structs.

---

Next: [Structural constraints](./04-structural-constraints.md)
