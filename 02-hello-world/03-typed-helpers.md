# Typed helper functions

The other, frequently seen pattern is a per-type helper function, ex. ([Go playground](https://gotipplay.golang.org/p/ouB4myMtHWS)):

```go
// PtrInt returns *i.
func PtrInt(i int) *int {
	return &i
}

// PtrStr returns *s.
func PtrStr(s string) *string {
	return &s
}

func main() {
	// Use the two helper functions that return pointers to their provided
	// values. Remember, this pattern must scale with the number of distinct,
	// defined types that need to be passed by pointer instead of value.
	print(request{
		host: PtrStr("local"),
		port: PtrInt(80),
	})
}
```

This approach requires a new function, per type. Surely there must be a more elegant solution...

---

Next: [A generic solution](./04-generic-solution.md)
