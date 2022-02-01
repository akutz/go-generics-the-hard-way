# Native

All of the below software must be installed and configured correctly to run the examples in this repository:

* [**Go 1.18beta2+**](https://go.dev/dl/#go1.18beta2): the version of Go with support for generics
* [**delve 1.18+**](https://github.com/go-delve/delve): the Go debugger
* [**.NET 6.0+**](https://dotnet.microsoft.com/en-us/download/dotnet/6.0): the .NET runtime and SDK
* [**lldb 11+**](https://lldb.llvm.org/man/lldb.html): a high-performance debugger
* [**sos**](https://docs.microsoft.com/en-us/dotnet/core/diagnostics/dotnet-sos): a .NET debugger extension for `lldb`
* [**OpenJDK 11+**](https://openjdk.java.net/): the OSS implementation of Java
* [**expect 5.45+**](https://linux.die.net/man/1/expect): talks to other programs
* [**Python 3**](https://python.org): for parsing Go benchmark results

While it _is_ possible to do so successfully both on macOS and Linux, it is highly recommended to use Docker.

* For those who _really_ want to run the examples on their local system, please proceed to the [_Native_](./native/) section. 
* For those who just want to run the examples without fuss or muss, please proceed normally.

---

Next: [Docker](./02-docker.md)
