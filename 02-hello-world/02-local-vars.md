# Local variables

The first solution often used is to declare local variables with the values intended for the optional fields and take the address of those variables, ex. ([Go playground](https://gotipplay.golang.org/p/phhmn6btClM)):

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

func main() {
	e, n := true, "akutz"
	printRequest(request{
		enabled:  &e,
		nickname: &n,
	})
}
```

This leads to cluttered code where the only purpose the variables serve is to their address can be taken.

---

Next: [Typed, helper functions](./03-typed-helpers.md)
