# Java

When generics were introduced to Java 5.0 back in 2004, they were implemented with _type erasure_ to provide backwards-compatibility with existing bytecode. However, this resulted in an implementation of generics that was largely a compile-time convenience feature for developers as the type information is erased at runtime. To see what this looks like, please consider the following program:

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

But what do `ints` and `strs` look like at runtime? Follow the instructions below to find out:

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

1. Set a breakpoint so the debugger will pause execution after both `ints` and `strs` have been defined and values added to them:

    ```bash
    stop at Main:34
    ```

1. Run the program:

    ```bash
    run
    ```

    which should stop when the above breakpoint is hit:

    ```bash
    Breakpoint hit: "thread=main", Main.main(), line=34 bci=57
    34            printLen(ints);
    ```

1. Now that it is loaded into memory, print information about the `ints` variable:

    ```bash
    dump ints
    ```

    ```
     ints = {
        serialVersionUID: 8683452581122892189
        DEFAULT_CAPACITY: 10
        EMPTY_ELEMENTDATA: instance of java.lang.Object[0] (id=444)
        DEFAULTCAPACITY_EMPTY_ELEMENTDATA: instance of java.lang.Object[0] (id=445)
        elementData: instance of java.lang.Object[10] (id=446)
        size: 3
        MAX_ARRAY_SIZE: 2147483639
        java.util.AbstractList.modCount: 3
        java.util.AbstractCollection.MAX_ARRAY_SIZE: 2147483639
    }
    ```

    Despite being defined at compile-time as an `ArrayList<Integer>`, at runtime `ints` has a type of `ArrayList<Object>`.

1. Print information about the `strs` variable:

    ```bash
    dump strs
    ```

    ```bash
     strs = {
        serialVersionUID: 8683452581122892189
        DEFAULT_CAPACITY: 10
        EMPTY_ELEMENTDATA: instance of java.lang.Object[0] (id=444)
        DEFAULTCAPACITY_EMPTY_ELEMENTDATA: instance of java.lang.Object[0] (id=445)
        elementData: instance of java.lang.Object[10] (id=448)
        size: 2
        MAX_ARRAY_SIZE: 2147483639
        java.util.AbstractList.modCount: 2
        java.util.AbstractCollection.MAX_ARRAY_SIZE: 2147483639
    }
    ```

    Despite being defined at compile-time as an `ArrayList<String>`, at runtime `strs` also has a type of `ArrayList<Object>`.

1. Type `quit` to exit the debugger.

1. Type `exit` to stop and remove the container.

In other words, generics in Java do **not** retain their type information at runtime. What about a more modern VM than the JVM such as Microsoft's Common Language Runtime (CLR) , the basis for .NET? Let's find out!

---

Next: [.NET](./03-dotnet.md)
