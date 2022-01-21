# Golang

Let's find out if Golang is also subject to the same type erasure as Java or retains type information like Microsoft .NET. Please consider the following program:

```golang
package main

import (
	"fmt"
	"reflect"
)

type List[T any] []T

func (a *List[T]) add(val T) {
	*a = append(*a, val)
}

func printLen[T any](list List[T]) {
	fmt.Println(len(list))
}

func main() {
	// Declare a new List[int] normally and add some values to it.
	var ints List[int]
	ints.add(1)
	ints.add(2)
	ints.add(3)

	// Get the type of the generic List in order to build a new List[string]
	// at runtime using reflection.
	_ = reflect.TypeOf(List)

	// There is no point in going any further as the above statement simply
	// does not compile.
	//
	// Generics in Go do not exist at runtime. Like Java they are a compile-time
	// convenience feature. While we already claimed that generic types in Go
	// do *not* get type erased, that is not entirely true.
	//
	// The generic "template", in this case "List[T any]" is erased at runtime.
	//
	// The only type that exists at runtime is main.List[int], because that is
	// what was instantiated.
	//
	// The Go compiler transforms all instantiated, generic types into concrete
	// types and drops the generic "templates" from the compiled binary.
}
```

The code comments in the above program describe why Go does not support runtime instantiation from generic types. More importantly, the above program illustrates that Go does not maintain any knowledge of the generic "templates" used to instantiate concrete types once the code is compiled.

To run the above program, please follow these instructions:

---

:warning: **Please note**

The `go build` command below uses the flag `-tags invalid`. This flag instructs Go to include `05-internals/03-runtime-instantiation/golang/main.go` in the compilation. Normally this file has a build constraint that prevents it from being considered to prevent:

* Dev tools (ex. an IDE such as VS Code) from constantly warning this file is in error
* a failed `go test ./...` command from the root of this repository

However, the entire _point_ of _this_ example is to demonstrate that failure, so the file needs to be "activated."

---

1. Launch the container:

    ```bash
    docker run -it --rm go-generics-the-hard-way
    ```

1. Attempt to build the program:

    ```bash
    go build -tags invalid ./05-internals/03-runtime-instantiation/golang/
    ```

    and the following compiler error will occur:

    ```bash
    # go-generics-the-hard-way/05-internals/03-runtime-instantiation/golang
    05-internals/03-runtime-instantiation/golang/main.go:46:21: cannot use generic type List without instantiation
    ```

1. Type `exit` to stop and remove the container.

Generics are not useful in Go until instantiated, and at the moment there is no way to refer to generic "templates" using reflection, which means it is not possible to instantiate new types using generics at runtime.

Please continue reading to see how Java, .NET, and Golang held up under this section's scrutiny...

---

Next: [In summary](../04-summary.md)
