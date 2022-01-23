# Golang

Let's find out if Golang is also subject to the same type erasure as Java or retains type information like Microsoft .NET. Please consider the following program:

```golang
package main

import (
	"fmt"
	"runtime"
)

type List[T any] []T

func (a *List[T]) add(val T) {
	*a = append(*a, val)
}

func printLen[T any](list List[T]) {
	fmt.Println(len(list))
}

func main() {
	var ints List[int]
	ints.add(1)
	ints.add(2)
	ints.add(3)

	var strs List[string]
	strs.add("Hello")
	strs.add("world")

	printLen(ints)
	printLen(strs)

	runtime.Breakpoint()
}
```

Just like the Java and .NET examples, the above program defines two variables using a very hacky, generic `List`:

* `ints`: a list of `int` values
* `strs`: a list of `string` values

What do `ints` and `strs` look like at runtime? Follow the instructions below to find out:

1. Launch the container:

    ```bash
    docker run -it --rm go-generics-the-hard-way
    ```

1. Load the above program into the Golang debugger:

    ```bash
    dlv debug ./05-internals/golang/main.go
    ```

1. Continue until the predefined breakpoint is hit:

    ```bash
    continue
    ```

    ```bash
    3
    2
    > main.main() ./05-internals/golang/main.go:48 (PC: 0x495c0d)
        43:	
        44:		printLen(ints)
        45:		printLen(strs)
        46:	
        47:		runtime.Breakpoint()
    =>  48:	}
    ```

1. Now that they are loaded into memory we can test the type safety of `ints` and `strs`. Unfortunately the Golang debugger `dlv` does not let us invoke the `ints.add(int)` function, but we _can_ show that `ints` definitely maintains its type safety at runtime by trying to assign `strs` to `ints`:

    ```bash
    set ints=strs
    ```

    The above command will fail with the following error:

    ```bash
    Command failed: can not convert value of type main.List[string] to main.List[int]
    ```

1. Type `quit` to exit the debugger.

1. Type `exit` to stop and remove the container.

In other words, Go does enforce type safety for generic types at runtime. In fact, it is almost as if generic types do not exist at all in compiled Golang binaries. Check out the next section to find out more!

---

Next: [Runtime instantiation](../03-runtime-instantiation/)
