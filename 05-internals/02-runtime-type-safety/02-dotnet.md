# .NET

Since .NET eschews type erasure and maintains type information at runtime, is it safe to assume that runtime type safety also exists in the CLR? Let's take a look by considering the following program:

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

We know that at runtime both `ints` and `strs` retain their type information. Does this mean it is not possible to add a `String` value to the `ints` list at runtime? Follow the instructions below to find out:

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
    Process 70 launched: '/go-generics-the-hard-way/05-internals/dotnet/bin/dotnet' (aarch64)
    3
    2
    Process 70 stopped
    * thread #1, name = 'dotnet', stop reason = signal SIGTRAP
        frame #0: 0x0000fffff798ae48 libcoreclr.so`___lldb_unnamed_symbol15329$$libcoreclr.so
    libcoreclr.so`___lldb_unnamed_symbol15329$$libcoreclr.so:
    ->  0xfffff798ae48 <+0>: brk    #0
        0xfffff798ae4c <+4>: ret    

    libcoreclr.so`___lldb_unnamed_symbol15330$$libcoreclr.so:
        0xfffff798ae50 <+0>: stp    x29, x30, [sp, #-0x10]!
        0xfffff798ae54 <+4>: ldp    x19, x20, [x0, #0xa0]
    ```

1. Now that the `ints` list is loaded into memory, let's take a look at its contents:

    ```bash
    clrstack -i ints
    ```

    ```bash
    Dumping managed stack and managed variables using ICorDebug.
    =============================================================================
    Child SP         IP               Call Site
    0000FFFFFFFFE890 0000fffff798ae48 [NativeStackFrame]
    0000FFFFFFFFE8D8 (null) [Internal call: 0000FFFFFFFFE8D8]
    0000FFFFFFFFEA50 0000ffff7dc88ea8 [DEFAULT] Void System.Diagnostics.Debugger.Break() (/root/.dotnet/shared/Microsoft.NETCore.App/6.0.1/System.Private.CoreLib.dll)

    PARAMETERS: (none)

    LOCALS: (none)

    0000FFFFFFFFEA60 0000ffff7e4dfa04 [DEFAULT] Void console.Program.Main(SZArray String) (/go-generics-the-hard-way/05-internals/dotnet/bin/dotnet.dll)

    PARAMETERS:
    + string[] args   (empty)

    LOCALS:
    + System.Collections.Generic.List`1&lt;int&gt; ints @ 0xffff500085f8
        |- int[] _items   (4 elements)
        |- int _size  = 3
        |- int _version  = 3
        |- int[] s_emptyArray   (empty)
    + System.Collections.Generic.List`1&lt;string&gt; strs @ 0xffff50008670

    0000FFFFFFFFEAA0 0000fffff7807688 [NativeStackFrame]
    Stack walk complete.
    =============================================================================
    ```

    Record the memory address of the variable, ex. `0xffff500085f8`.

1. Dump the object at the recorded address:

    ```bash
    dumpobj 0xffff500085f8
    ```

    ```bash
    Name:        System.Collections.Generic.List`1[[System.Int32, System.Private.CoreLib]]
    MethodTable: 0000ffff7e594000
    EEClass:     0000ffff7e623598
    Tracked Type: false
    Size:        32(0x20) bytes
    File:        /root/.dotnet/shared/Microsoft.NETCore.App/6.0.1/System.Private.CoreLib.dll
    Fields:
                MT    Field   Offset                 Type VT     Attr            Value Name
    0000ffff7e548080  4001fa5        8       System.Int32[]  0 instance 0000ffff50008648 _items
    0000ffff7e539018  4001fa6       10         System.Int32  1 instance                3 _size
    0000ffff7e539018  4001fa7       14         System.Int32  1 instance                3 _version
    0000ffff7e548080  4001fa8        8       System.Int32[]  0   static dynamic statics NYI                 s_emptyArray
    ```

    Record the address of the object's method table, ex. `0000ffff7e594000`.

1. Dump the method table for the object:

    ```bash
    dumpmt -MD 0000ffff7e594000
    ```

    This will emit a _lot_ of information, but it is this line in particular that is of interest:

    ```
    0000FFFF7DD24A40 0000FFFF7E593BB0 PreJIT System.Collections.Generic.List`1[[System.Int32, System.Private.CoreLib]].Add(Int32)
    ```

    The method `Add` for the `ints` object takes a single argument, and it is type `Int32`, supporting the idea that .NET is not going to allow a value of a type other than `Int32` to be added to a `List<Int32>`. It would be nice to verify this by calling the `Add` function, but calling managed code using the `lldb` debugger is quite difficult, and I have not yet figured out how to do so.

1. Detach from the process:

    ```
    process detach
    ```

1. Type `quit` to exit the debugger.

1. Type `exit` to stop and remove the container.

In conclusion, generics in .NET **do** enforce runtime type safety. Is the same true for Golang?

---

Next: [Golang](./03-golang.md)
