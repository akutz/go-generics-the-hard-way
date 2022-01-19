# .NET

Not only does .NET enforce runtime type safety, but it is actually possible to create new generic types at runtime! Please take a look at the following program:

```csharp
using System;
using System.Runtime;
using System.Reflection;
using System.Collections;

namespace console
{
    class Program
    {
        static void printLen(IList list) {
            Console.Out.WriteLine(list.Count);
        }

        static IList newGenericListFor(String className) {
            // We can get the Class based on a string value -- a useful feature
            // for some meta programming.
            var clazz = TypeInfo.GetType(className);

            // Build a new list constructor using the class information.
            var listCtor = typeof(List<>).MakeGenericType(clazz);

            // Construct a new List<clazz> like we did before, only this time
            // using reflection at runtime.
            //
            // Please note the result is cast to a List. This is to ensure
            // the standard List methods such as Add, and fields such as Count,
            // are accessible via the ints variable.
            return (IList) Activator.CreateInstance(listCtor);
        }

        static void Main(string[] args) {

            // Create a List<Int32>.
            var ints = newGenericListFor("System.Int32");

            // Add some numbers to the list.
            ints.Add(1);
            ints.Add(2);
            ints.Add(3);

            if (args.Length > 0 && args[0] == "fail") {
                // If this program is executed with a single argument "fail",
                // then we try to add "Hello" to the ints list.
                //
                // This will compile because ints is not known to the compiler
                // yet as a List<Int32>, even though that is what it is.
                //
                // At runtime,however, this will fail because you cannot add a
                // String to a List<Int32>.
                ints.Add("Hello");
            }

            // Create a List<String>.
            var strs = newGenericListFor("System.String");

            // Add some strings to the list.
            strs.Add("Hello");
            strs.Add("world");

            printLen(ints);
            printLen(strs);
        }
    }
}
```

The code comments in the above program demonstrate that .NET does in fact allow runtime instantiation from generic types, and maintains their runtime type safety, even if this does not exist at compile time due to the meta programming involved.

To run the above program, please follow these instructions:

1. Launch the container:

    ```bash
    docker run -it --rm --cap-add=SYS_PTRACE --security-opt seccomp=unconfined go-generics-the-hard-way
    ```

    Please note the `--cap-add=SYS_PTRACE --security-opt seccomp=unconfined` flags are required in order to use the `lldb` debugger to attach to a .NET process.

1. Compile the above program:

    ```bash
    dotnet build --debug -p:UseSharedCompilation=false -o ./05-internals/03-runtime-instantiation/dotnet/bin ./05-internals/03-runtime-instantiation/dotnet
    ```

1. Run the program normally:

    ```bash
    ./05-internals/03-runtime-instantiation/dotnet/bin/dotnet
    ```

    ```bash
    3
    2
    ```

1. Run the program again forcing a runtime type check error:

    ```bash
    ./05-internals/03-runtime-instantiation/dotnet/bin/dotnet fail
    ```

    ```bash
    Unhandled exception. System.ArgumentException: The value "Hello" is not of type "System.Int32" and cannot be used in this generic collection. (Parameter 'value')
       at System.Collections.Generic.List`1.System.Collections.IList.Add(Object item)
       at console.Program.Main(String[] args) in /go-generics-the-hard-way/05-internals/03-runtime-instantiation/dotnet/main.cs:line 66
    Aborted
    ```

1. Type `exit` to stop and remove the container.

.NET clearly knows what it is doing when it comes to generics and runtime instantiation. Does Go keep up?

---

Next: [Golang](./03-golang.md)
