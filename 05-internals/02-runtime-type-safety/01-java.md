# Java

So we now know that generic types in Java do not preserve that type information at runtime. What is the consequence of this? Let's once again consider the following program:

```java
import java.util.ArrayList;

class Main {
    public static void printLen(ArrayList<?> list)  {
        System.out.println(list.size());
    }

    public static void main(String[] args) {
        ArrayList<Integer>  ints = new ArrayList<Integer>();
        ints.add(1);
        ints.add(2);
        ints.add(3);

        ArrayList<String>   strs = new ArrayList<String>();
        strs.add("Hello");
        strs.add("world");

        printLen(ints);
        printLen(strs);
    }
}
```

The above program defines two variables using the generic list type, `ArrayList`:

* `ints`: a list of `Integer` values
* `strs`: a list of `String` values

We know that at runtime both `ints` and `strs` lost their type information and each become `ArrayList<Object>`. Does this mean it is possible to add a `String` value to the `ints` list at runtime? Follow the instructions below to find out:

1. Launch the container:

    ```bash
    docker run -it --rm go-generics-the-hard-way
    ```

1. Compile the above program:

    ```bash
    javac -g ./05-internals/java/main.java
    ```

1. Load the above program into the Java debugger:

    ```bash
    jdb -sourcepath ./05-internals/java -classpath ./05-internals/java Main
    ```

1. Set a breakpoint so the debugger will pause execution after the `ints` list has been defined and values added to it:

    ```bash
    stop at Main:30
    ```

1. Run the program:

    ```bash
    run
    ```

    which should stop when the above breakpoint is hit:

    ```bash
    Breakpoint hit: "thread=main", Main.main(), line=30 bci=35
    30            ArrayList<String>   strs = new ArrayList<String>();
    ```

1. Now that it is loaded into memory, let's print the contents of the `ints` list:

    ```bash
    print ints
    ```

    ```
     ints = "[1, 2, 3]"
    ```

1. Try adding the string "Hello" to the `ints` list:

    ```bash
    eval ints.add("Hello")
    ```

    ```bash
     ints.add("Hello") = true
    ```

    It looks like it worked!

1. Print the contents of `ints` again to verify it now contains three numbers and a string:

    ```bash
    print ints
    ```

    ```bash
     ints = "[1, 2, 3, Hello]"
    ```

    This is the consequence of type erasure. At runtime Java provides no type safety for its generic types.

1. Type `quit` to exit the debugger.

1. Type `exit` to stop and remove the container.

In other words, generics in Java are purely a compile-time convenience feature. How does .NET stack up? Keep reading and find out.

---

Next: [.NET](./02-dotnet.md)
