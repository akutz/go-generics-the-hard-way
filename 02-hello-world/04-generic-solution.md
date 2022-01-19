# A generic solution

With Go generics, a truly elegant solution to the above problem is finally possible -- a single function that can return the address of a value of any type, ex. ([Go playground](https://gotipplay.golang.org/p/d0dgvmP6bII)):

```go
package main

import (
	"fmt"
)

type request struct {
	enabled  *bool
	nickname *string
}

func printRequest(r request) {
	e, n := "nil", "nil"
	if r.enabled != nil {
		e = fmt.Sprintf("%v", *r.enabled)
	}
	if r.nickname != nil {
		n = *r.nickname
	}
	fmt.Printf("request: enabled=%s, nickname=%s\n", e, n)
}

// ptr returns the address of t.
func ptr[T any](t T) *T {
	return &t
}

func main() {
	printRequest(request{
		enabled:  ptr(true),
		nickname: ptr("akutz"),
	})
}

```

The function `ptr[T any](t T) *T`:

* Takes any value of any type
* Returns the address for that value
* Does **_not_** box the value of `t`

How does this all work? What exactly is `T`? What is `any`?  Keep reading to find out :smiley:...

---

Next: [Getting started](../03-getting-started/)
