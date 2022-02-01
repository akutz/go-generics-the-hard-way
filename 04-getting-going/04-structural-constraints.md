# Structural constraints

So far all constraints have referred to primitive types or interfaces expressing a composite set of defined types or type definitions with the same underlying type of defined types. However, generics in Go also supports structural constraints. Imagine the following use case:

* we need a function that accepts _any_ struct that has fields for:
  * ID
  * numeric values
  * sum function

In other words what is needed is a function that takes an instance of `Ledger[T, K]` ([Go playground](https://gotipplay.golang.org/p/Mf1pTRmx-FO)):

```golang
func SomeFunc[T ~string, K Numeric](l Ledger[T, K]) {}

func main() {
	SomeFunc[string, int](Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	})
}
```

---

:warning: Eagle-eyed readers may notice that `SomeFunc` does not actually do anything. That is because there is currently a bug in Go 1.18beta2 that _should_ be resolved by the time 1.18 is released. For now, the issue [golang/go#50417](https://github.com/golang/go/issues/50417) prevents reading and writing fields with structural constraints.

---

The program compiles and runs fine, except the use case was not for a function that takes an instance of `Ledger[T, K]` but for a function that matches _any_ struct with those fields. That sounds a lot like the tilde `~` operator... ([Go playground](https://gotipplay.golang.org/p/HSqEpRWkr-8)):

```golang
func SomeFunc[T ~string, K Numeric, L ~Ledger[T, K]](l L) {}

func main() {
	SomeFunc[string, int, Ledger[string, int]](Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	})
}
```

But of course this does not compile:

```bash
./prog.go:47:39: invalid use of ~ (underlying type of Ledger[T, K] is struct{ID T; Amounts []K; SumFn SumFn[K]})
./prog.go:50:24: cannot implement ~Ledger[string, int] (empty type set)
```

The secret here is _all_ structs implement the anonymous struct, so if we want to match all structs with the aforementioned fields, we want to use `~` with an anonymous struct ([Go playground](https://gotipplay.golang.org/p/6DwJnBiYD4J)):

```golang
func SomeFunc[
	T ~string,
	K Numeric,
	L ~struct {
		ID      T
		Amounts []K
		SumFn   SumFn[K]
	},
](l L) {
}

func main() {
	SomeFunc[string, int, Ledger[string, int]](Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	})
}
```

In fact, the above example even be rewritten so the call to `SomeFunc` relies on type inference ([Go playground](https://gotipplay.golang.org/p/6ds7bDq2_ep)):

```golang
func main() {
	SomeFunc(Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	})
}
```

Now that `SomeFunc` can accept any struct that matches the constraint, it's possible to send other types as well ([Go playground](https://gotipplay.golang.org/p/Mng1uoHqKg5)):


```golang
type ID string

type CustomLedger struct {
	ID      ID
	Amounts []uint64
	SumFn   SumFn[uint64]
}

func main() {
	// Call SomeFunc with an anonymous struct that uses a type
	// alias for "string" as the type for the "ID" field.
	//
	// Please note that because a type alias is used in a nested
	// type, inference does not work.
	SomeFunc[ID, float32, struct {
		ID      ID
		Amounts []float32
		SumFn   SumFn[float32]
	}](struct {
		ID      ID
		Amounts []float32
		SumFn   SumFn[float32]
	}{
		ID:      ID("fake"),
		Amounts: []float32{1, 2, 3},
		SumFn:   Sum[float32],
	})

	// Compare that to this call which *also* uses an anonymous
	// struct, but gets to rely on type inference. This is because
	// the "ID" field is a "string" and can be infered from the
	// nested type.
	SomeFunc(struct {
		ID      string
		Amounts []float32
		SumFn   SumFn[float32]
	}{
		ID:      "fake",
		Amounts: []float32{1, 2, 3},
		SumFn:   Sum[float32],
	})

	// Call SomeFunc a Ledger[T, K].
	SomeFunc(Ledger[string, complex64]{
		ID:      "fake",
		Amounts: []complex64{1, 2, 3},
		SumFn:   Sum[complex64],
	})

	// SomeFunc can also be used with other concrete struct types
	// as long as they satisfy the constraint.
	SomeFunc[ID, uint64, CustomLedger](CustomLedger{
		ID:      ID("fake"),
		Amounts: []uint64{1, 2, 3},
		SumFn:   Sum[uint64],
	})
}
```

However, please note the following will _not_ work ([Go playground](https://gotipplay.golang.org/p/7DlFxBI2rEz)):

```golang
type LedgerNode struct {
	ID      string
	Amounts []uint64
	SumFn   SumFn[uint64]
	Next    *LedgerNode
}

func main() {
	// This will fail because LedgerNode does not match the constraint
	// exactly. The presence of the additional field will result in a
	// compiler error.
	SomeFunc(LedgerNode{
		ID:      "fake",
		Amounts: []uint64{1, 2, 3},
		SumFn:   Sum[uint64],
	})
}
```

Instead the above fails to compile with the following error:

```bash
./prog.go:69:10: L does not match struct{ID T; Amounts []K; SumFn SumFn[K]}
```

Structural constraints must match the struct _exactly_, and this means even if all of the fields in the constraint are present, the presence of additional fields in the provided value means the type does not satisfy the constraint.

So how does one express the presence of additional information? With an interface of course!

---

Next: [Interface constraints](./05-interface-constraints.md)
