# .NET

The Microsoft .NET common language runtime (CLR) VM is responsible for running all of the .NET languages, but this example and others in this repository will use C#. Let's find out if the CLR is also subject to the same type erasure as Java. Please consider the following program:

```csharp
using System.Diagnostics;

namespace console
{
    class Program
    {
        static void printLen<T>(List<T> list) {
            Console.Out.WriteLine(list.Count);
        }

        static void Main(string[] args)
        {
            var ints = new List<Int32>();
            ints.Add(1);
            ints.Add(2);
            ints.Add(3);

            var strs = new List<String>();
            strs.Add("Hello");
            strs.Add("world");

            printLen(ints);
            printLen(strs);

            Debugger.Break();
        }
    }
}
```

Just like the Java example, the above program defines two variables using C#'s generic list type, `List`:

* `ints`: a list of `Int32` values
* `strs`: a list of `String` values

What do `ints` and `strs` look like at runtime? Follow the instructions below to find out:

1. Launch the container:

    ```bash
    docker run -it --rm --cap-add=SYS_PTRACE --security-opt seccomp=unconfined go-generics-the-hard-way
    ```

    Please note the `--cap-add=SYS_PTRACE --security-opt seccomp=unconfined` flags are required in order to use the `lldb` debugger to attach to a .NET process.

1. Compile the above program:

    ```bash
    dotnet build --debug -p:UseSharedCompilation=false -o ./05-internals/dotnet/bin ./05-internals/dotnet
    ```

1. Load the above program into the .NET debugger:

    ```bash
    lldb ./05-internals/dotnet/bin/dotnet
    ```

1. Launch a new process using the provided program and attach the debugger to it:

    ```bash
    process launch
    ```

1. The process should launch and continue until the predefined breakpoint is hit:

    ```bash
    Process 90 launched: '/go-generics-the-hard-way/05-internals/dotnet/bin/dotnet' (x86_64)
    3
    2
    Process 90 stopped
    * thread #1, name = 'dotnet', stop reason = signal SIGTRAP
        frame #0: 0x00007ffff77c0571 libcoreclr.so`___lldb_unnamed_symbol15647$$libcoreclr.so + 1
    libcoreclr.so`___lldb_unnamed_symbol15647$$libcoreclr.so:
    ->  0x7ffff77c0571 <+1>: retq   
        0x7ffff77c0572 <+2>: nop    

    libcoreclr.so`___lldb_unnamed_symbol15648$$libcoreclr.so:
        0x7ffff77c0574 <+0>: pushq  %rbp
        0x7ffff77c0575 <+1>: movq   0xd8(%rdi), %r12
    ```

1. Now that they are loaded into memory, print information about the `ints` and `strs` variables:

    ```bash
    clrstack -i -a
    ```

    ```bash
    Dumping managed stack and managed variables using ICorDebug.
    =============================================================================
    Child SP         IP               Call Site
    00007FFFFFFFDCC8 00007ffff77c0571 [NativeStackFrame]
    00007FFFFFFFDD18 (null) [Internal call: 00007FFFFFFFDD18]
    00007FFFFFFFDE40 00007fff7daa3777 [DEFAULT] Void System.Diagnostics.Debugger.Break() (/root/.dotnet/shared/Microsoft.NETCore.App/6.0.1/System.Private.CoreLib.dll)

    PARAMETERS: (none)

    LOCALS: (none)

    00007FFFFFFFDE50 00007fff7e25317a [DEFAULT] Void console.Program.Main(SZArray String) (/go-generics-the-hard-way/05-internals/dotnet/bin/dotnet.dll)

    PARAMETERS:
    + string[] args   (empty)

    LOCALS:
    + System.Collections.Generic.List`1&lt;int&gt; ints @ 0x7fff48008758
    + System.Collections.Generic.List`1&lt;string&gt; strs @ 0x7fff480087b8

    00007FFFFFFFDE90 00007ffff762aa27 [NativeStackFrame]
    Stack walk complete.
    =============================================================================
    ```

    In addition to being defined at compile-time as `List<Int32>` and `List<String>`, the variables `ints` and `strs` maintain their full type information at runtime as `List<int>` and `List<string>` (`Int32` and `int` are interchangeable as are `String` and `string`).

1. In fact, not only are the types not erased, but .NET maintains knowledge of the generic type from which these types were instantiated. To illustrate this we need to grab the metadata tokens for the underlying classes used by the `ints` and `strs` variables.

    In the previous step please note the memory addresses of each variable:

      * `ints`: `0x7fff48008758`
      * `strs`: `0x7fff480087b8`

1. Dump the `ints` object using its memory address:

    ```bash
    dumpobj 0x7fff48008758
    ```

    ```bash
    Name:        System.Collections.Generic.List`1[[System.Int32, System.Private.CoreLib]]
    MethodTable: 00007fff7e2f7d20
    EEClass:     00007fff7e373598
    Tracked Type: false
    Size:        32(0x20) bytes
    File:        /root/.dotnet/shared/Microsoft.NETCore.App/6.0.1/System.Private.CoreLib.dll
    Fields:
                MT    Field   Offset                 Type VT     Attr            Value Name
    00007fff7e2b8080  4001fa5        8       System.Int32[]  0 instance 00007fff48008790 _items
    00007fff7e2a9018  4001fa6       10         System.Int32  1 instance                3 _size
    00007fff7e2a9018  4001fa7       14         System.Int32  1 instance                3 _version
    00007fff7e2b8080  4001fa8        8       System.Int32[]  0   static dynamic statics NYI                 s_emptyArray
    ```

    Record the memory address for the `EEClass`, ex. `00007fff7e373598`.

1. Dump the `EEClass` for the `ints` variable using the address, ex. `00007fff7e373598`:

    ```bash
    dumpclass 00007fff7e373598
    ```

    ```bash
    Class Name:      System.Collections.Generic.List`1[[System.Int32, System.Private.CoreLib]]
    mdToken:         0000000002000865
    File:            /root/.dotnet/shared/Microsoft.NETCore.App/6.0.1/System.Private.CoreLib.dll
    Parent Class:    00007fff7e1fa230
    Module:          00007fff7d6d4000
    Method Table:    00007fff7e2f7d20
    Vtable Slots:    1e
    Total Method Slots:  4e
    Class Attributes:    102001  
    NumInstanceFields:   3
    NumStaticFields:     1
                MT    Field   Offset                 Type VT     Attr            Value Name
    00007fff7e2b8080  4001fa5        8       System.Int32[]  0 instance           _items
    00007fff7e2a9018  4001fa6       10         System.Int32  1 instance           _size
    00007fff7e2a9018  4001fa7       14         System.Int32  1 instance           _version
    00007fff7e2b8080  4001fa8        8       System.Int32[]  0   static dynamic statics NYI                 s_emptyArray
    ```

    Note the `ints` variable ultimately uses a class with the `mdToken` at memory address `0000000002000865`.

1. Print the metadata for the `ints` variable:

    ```bash
    sos Token2EE System.Private.CoreLib.dll 0000000002000865
    ```

    ```bash
    Module:      00007fff7d6d4000
    Assembly:    System.Private.CoreLib.dll
    Token:       0000000002000865
    MethodTable: 00007fff7e2dd7e8
    EEClass:     00007fff7e2ec598
    Name:        System.Collections.Generic.List`1
    ```

1. Repeat the above steps for the `strs` variable, starting by dumping its object at address `0x7fff480087b8`:

    ```bash
    dumpobj 0x7fff480087b8
    ```

    It has an `EEClass` address of `00007fff7e2ee0b0`, so dump that too:

    ```bash
    dumpclass 00007fff7e2ee0b0
    ```

    Which prints a metadata token of...`0000000002000865`. The same as `ints`! This is because both `ints` and `strs` ultimately use the same generic template, `List<T>`, to build their types, and .NET maintains that information at runtime.

1. Detach from the process:

    ```
    process detach
    ```

1. Type `quit` to exit the debugger.

1. Type `exit` to stop and remove the container.

In other words, generics in .NET **do** retain their type information at runtime. So what about Golang?

---

Next: [Golang](./04-golang.md)
