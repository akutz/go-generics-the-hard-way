## The problem

If you have ever worked with the Amazon Web Services SDK for Go or Kubernetes Custom Resource Definitions (CRD) in Go, then you are very aware of all the occasions where some optional field takes a `string` or `bool` by address, for example ([Go playground](https://gotipplay.golang.org/p/ZRvbHRYAodL)):

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
	printRequest(request{
		enabled:  nil, // needs address of a bool
		nickname: nil, // needs address of a string
	})
}
```

There are two, common solutions for solving the above conundrum:

* [**Local variables**](./02-local-vars.md): a solution using temporary, local variables
* [**Typed helper functions**](./03-typed-helpers.md): a solution using typed, helper functions

---

Next: [Local variables](./02-local-vars.md)
