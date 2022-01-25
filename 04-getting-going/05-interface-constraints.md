# Interface constraints

The previous page discussed structural constraints, but those are limited to constraints based on the fields of a struct. What about its function receivers? Let's first rewrite the example from the previous page as a distinct constraint that is not inlined to the function that uses it:

```golang
// Ledgerish expresses a constraint that may be satisfied by types that have
// ledger-like qualities.
type Ledgerish[T ~string, K Numeric] interface {
	~struct {
		ID      T
		Amounts []K
		SumFn   SumFn[K]
	}
}
```

Again, we know that `Ledger[T, K]` _does_ have a function receiver:

```golang
// PrintIDAndSum emits the ID of the ledger and a sum of its amounts on a
// single line to stdout.
func (l Ledger[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}
```

So how do we write the following function to match `Ledgerish` in such a way that we can also invoke `PrintIDAndSum`:

```golang
// PrintLedger emits a ledger's ID and total amount on a single line
// to stdout.
func PrintLedger[T ~string, K Numeric, L Ledgerish[T, K]](l L) {
	l.PrintIDAndSum()
}
```

If we tried it as it is now, it would fail ([Golang playground](https://gotipplay.golang.org/p/AT2mw5btoYo)):

```golang
func main() {
	PrintLedger(Ledger[string, complex64]{
		ID:      "fake",
		Amounts: []complex64{1, 2, 3},
		SumFn:   Sum[complex64],
	})
}
```

The above example will produce the following compilier error:

```bash
./prog.go:60:4: l.PrintIDAndSum undefined (type L has no field or method PrintIDAndSum)
```

The secret this time is to remember that `Ledgerish` _is_ a Go interface, and as such, can have methods defined on it:

```golang
// Ledgerish expresses a constraint that may be satisfied by types that have
// ledger-like qualities.
type Ledgerish[T ~string, K Numeric] interface {
	~struct {
		ID      T
		Amounts []K
		SumFn   SumFn[K]
	}

	PrintIDAndSum()
}
```

Now the above example works as intended ([Golang playground](https://gotipplay.golang.org/p/N2xwtM91D-E)):

```bash
fake has a sum of (6+0i)
```

But hey, wait a minute! If interfaces can be used to express constraints, when and why would the strictness of a structural constraint be desired over a traditional, functional, interface-based constraint?

---

Next: [Careful constructors](./06-careful-constructors.md)
