## The problem

If you have ever worked with the Amazon Web Services SDK for Go or Kubernetes Custom Resource Definitions (CRD) in Go, then you are no doubt quite familiar when some argument or field takes a pointer to a `string` or `int`, for example ([Go playground](https://gotipplay.golang.org/p/r6Dv9T0kCvJ)):

```go
package main

import (
	"fmt"
)

type request struct {
	host *string
	port *int
}

func print(r request) {
	fmt.Print("request: host=")
	if r.host != nil {
		fmt.Print(*r.host)
	}
	fmt.Print(", port=")
	if r.port != nil {
		fmt.Printf("%d", *r.port)
	}
	fmt.Println()
}

func main() {
	print(request{
		host: nil, // needs a *string
		port: nil, // needs a *int
	})
}
```

There are two, common -- albeit far from perfect -- solutions for addressing the issue:

* [**Local variables**](./02-local-vars.md): a solution using temporary, local variables
* [**Typed helper functions**](./03-typed-helpers.md): a solution using typed, helper functions

---

Next: [Local variables](./02-local-vars.md)
