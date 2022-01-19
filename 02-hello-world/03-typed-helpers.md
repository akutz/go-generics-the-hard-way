# Typed helper functions

The others solutions we commonly see is to create a helper function for each type, ex. ([Go playground](https://gotipplay.golang.org/p/U9otocywibi)):

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

func ptrBool(b bool) *bool {
	return &b
}

func ptrString(s string) *string {
	return &s
}

func main() {
	printRequest(request{
		enabled:  ptrBool(true),
		nickname: ptrString("akutz"),
	})
}
```

This approach means needing to redefine variants of a function for taking the address of a value for every type of value whose address needs to be taken.

---

Next: [A generic solution](./04-generic-solution.md)
