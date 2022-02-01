# Visual Studio Code

Setting up VS Code to use Go 1.18beta2 is fairly straight-forward:

* [**Update the Go language server**](#update-the-go-language-server): build `gopls` with Go 1.18beta2
* [**Configure `GOROOT`**](#configure-goroot): update VS Code to use Go 1.18beta2


## Update the Go language server

VS Code uses the Go language server, `gopls`, for dot completion, validating the source code, etc. Suffice it to say, for VS Code to function correctly with Go 1.18beta2, the Go language server needs to know about Go 1.18beta2. To update `gopls`, please follow the commands below:

1. Open a console window.

1. Execute the following command:

    ```bash
    go install golang.org/x/tools/gopls@latest
    ```

## Configure `GOROOT`

Now that the language server is updated, we can configure VS Code to use Go 1.18beta2:

1. Open up a console window.

---

:warning: **Please note** the next step assumes the `go` program is 1.18beta2, which it _should_ be if the instructions in the previous section were followed.

---

1. Execute the following command:

    ```bash
    go env GOROOT
    ```

1. Write down the path printed by the above command.

1. Open VS Code's preferences/settings.

1. Search for `GOROOT`

1. Click on `Edit in settings.json`

1. Modify or add the property `go.goroot` and set it to the directory recorded up above.

Voilá, VS Code should now be able to sucessfully handle Go 1.18beta2 code, including generics!

---

Next: [Hello world](../../02-hello-world/)
