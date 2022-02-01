# Delve

The [delve](https://github.com/go-delve/delve) program is a debugger for Golang and should feel at home to people familiar with `gdb`.

* [**Install delve**](#install-delve): install the go debugger
* [**Configure delve**](#configure-delve): ensure delve uses Go 1.18beta2

## Install delve

Install the latest release (1.18+) to support Go 1.18beta2 with the following command:

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

## Configure delve

Even with the latest release, `dlv debug` may not work as expected"

1. Consider the following Go program:

    ```go
    package main

    import "fmt"

    func print[T any](t T) {
    	fmt.Println(t)
    }

    func main() {
    	print("Hello, world.")
    }
    ```

1. Save the program as `main.go`

1. Run the program through the debugger:

    ```bash
    $ dlv debug main.go
    # command-line-arguments
    ./main.go:5:6: missing function body
    ./main.go:5:11: syntax error: unexpected [, expecting (
    exit status 2
    ```

This fails because delve [uses the `go` binary from the shell's `PATH`](https://github.com/go-delve/delve/blob/master/pkg/gobuild/gobuild.go) to build a binary suitable for debugging. To fix this issue the Go 1.18 binary needs to be the program called when delve invokes the `go` command:

1. If Go 1.18beta2 was installed from the instructions on the previous page, then use the following command to ensure that Go 1.18beta2's `go` binary appears first in the shell's `PATH`:

    ```bash
    export PATH="$(go1.18beta2 env GOROOT)/bin:${PATH}"
    ```

2. Verify that `go version` shows Go 1.18beta2:

    ```bash
    $ go version
    go version go1.18beta2 linux/amd64
    ```

3. Run the delve debugger again:

    ```bash
    $ dlv debug main.go       
    Type 'help' for list of commands.
    (dlv) 
    ```

And that's it! Delve should now be set up to successfully build and debug Go1.18 programs!

---

Next: [Configuring VS Code](./03-vscode.md)
