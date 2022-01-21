# Local variables

The first solution often used is to declare local variables in order to derive pointers, ex. ([Go playground](https://gotipplay.golang.org/p/dddTfDcdapn)):

```go
func main() {
	// Declare "host" and "port" in order to create pointers to satisfy the
	// fields in the "request" struct.
	host, port := "local", 80

	print(request{
		host: &host,
		port: &port,
	})
}
```

This leads to cluttered, hard-to-read code where the only purpose variables serve is for deriving pointers.

---

Next: [Typed, helper functions](./03-typed-helpers.md)
