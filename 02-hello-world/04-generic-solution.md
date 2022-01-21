# A generic solution

With Go generics, a truly elegant solution to the above problem is finally possible -- a single function that can return the address of a value of any type, ex. ([Go playground](https://gotipplay.golang.org/p/dqCpdQjbmZ8)):

```go
// Ptr returns *value.
func Ptr[T any](value T) *T {
	return &value
}

func main() {
	// No local variables and the typed helper functions can be collapsed into
	// a single, generic function for getting a pointer to a value.
	print(request{
		host: Ptr("local"),
		port: Ptr(80),
	})
}
```

The function `Ptr[T any](T) *T`:

* Takes any value of any type
* Does **_not_** box the value
* Returns a pointer for that value

How does this all work? What exactly is `T`? What is `any`?  Keep reading to find out :smiley:...

---

Next: [Getting started](../03-getting-started/)
