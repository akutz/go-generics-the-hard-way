# Multiple generic types

Up to this point all of the examples have used a single generic type, `T`, in the function `Sum[T Numeric](...T) T`. But how are multiple generic, types expressed? Imagine we need to define a function that can:

* receive an ID that can be represented by a string value
* receive a function that can be used to sum numeric values
* receive zero to many numeric values
* print the ID and sum of the values on a single line

With what we have learned so far, it _should_ be possible to satisfy the above use case with multiple, generic types. Let's look at the first requirement:

> we need to define a function that can

So there needs to be a single function that can handle the remaining requirements, got it.

> receive an ID that can be represented by a string value

This sounds like a generic constraint that can be expressed as `~string`. With that we can probably start building our function:

```golang
func SomeFunc[T ~string](id T) {}
```

> receive a function that can be used to sum numeric values

This definitely seems like our old friend `Sum[T Numeric](...T) T` :smiley:, but it looks like `SomeFunc` needs to _receive_ our friend, so how would we do that normally? Ah yes, a function signature!

```golang
// SumFn defines a function that can return the sum of one to zero numbers.
type SumFn[T Numeric] func(...T) T

func SomeFunc[T ~string, K Numeric](id T, fn SumFn[K]) {}
```

> receive zero to many numeric values

Okay, this would be a variadic based on the generic constraint `Numeric`. Easy enough:

```golang
func SomeFunc[T ~string, K Numeric](id T, sum SumFn[K], values ...K) {}
```

> print the ID and sum of the values on a single line

Now we are getting to the heart of it! And since we know the intent, let's go ahead and rename the function:

```golang
// PrintIDAndSum prints the provided ID and sum of the given values to stdout.
func PrintIDAndSum[T ~string, K Numeric](id T, sum SumFn[K], values ...K) {

	// The format string uses "%v" to emit the sum since using "%d" would
	// be invalid if the value type was a float or complex variant.
	fmt.Printf("%s has a sum of %v\n", id, sum(values...))
}
```

Alright, let's put all this together into a program ([Golang playground](https://gotipplay.golang.org/p/ZeSWNvI-SQi)):

```golang
func main() {
	PrintIDAndSum("acct-1", Sum, 1, 2, 3)
}
```

If you have been reading from the beginning you are probably thinking _This is where he says "fails to compile"._ Am I really that predictable?

.
.
.

And the above example, of course, fails to compile:

```bash
./prog.go:36:26: cannot use generic function Sum without instantiation
```

One of the key points on the page about type inference was:

> The Go compiler tries _really_ hard to infer the intended types, but it does not always work when you think it should

Because of the way the `PrintIDAndSum[T ~string, K Numeric](T, SumFn[K])` is written, it would be easy to assume that the Go compiler _should_ be able to infer the numeric type for the `Sum` function, but it just does not always work when you think it should. Instead the program needs to specify the type for `Sum` explicitly:

```golang
func main() {
	PrintIDAndSum("acct-1", Sum[complex128], 1, 2, 3)
}
```

Now the program will produce the expected output:

```bash
acct-1 has a sum of (6+0i)
```

So far generic types have been used by functions as:

* argument types
* return types

But there has been one thing at the heart of each example that has used the `Sum[T Numeric](...T) T` function that has not yet been examined...

---

Next: [Getting going](../04-getting-going/)
