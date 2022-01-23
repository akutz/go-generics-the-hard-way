# Java

Because Java employes type erasure, it is not possible, nor would it be useful, to create new, generic types at runtime. Take the following program:

```java
import java.lang.Class;
import java.util.ArrayList;
import java.lang.reflect.Constructor;

class Main {
    public static void printLen(ArrayList<?> list)  {
        System.out.println(list.size());
    }

    public static void main(String[] args) throws Exception {
        // We can get the Class based on a string value -- a useful feature for
        // some meta programming.
        Class<?> intClass = Class.forName("java.lang.Integer");

        // Get the constructor for instantiating an ArrayList.
        Constructor<ArrayList> listCtor = ArrayList.class.getDeclaredConstructor();

        // But there is no way to instantiate ArrayList using the intClass!
        // Instead all we can do is create the ArrayList the way it would look
        // at runtime anyway.
        ArrayList ints = listCtor.newInstance();

        // Add some integers to the list.
        ints.add(1);
        ints.add(2);
        ints.add(3);

        // And because there is no type safety for lists created at runtime,
        // and we used reflection to mimic that here, there's no compile-time
        // type safety either!
        ints.add("Hello");

        // That said, we *can* assert some compile-time type safety back into
        // the mix by randomly declaring "strs" to be an ArrayList<String>.
        //
        // But remember, we did *nothing* to ensure the underlying storage of
        // strs is bound to the String class. Again, this is purely a compile-
        // time conveience feature.
        ArrayList<String> strs = listCtor.newInstance();
        strs.add("Hello");
        strs.add("world");

        // By asserting that strs is "ArrayList<String>", the following line
        // would cause a compile error as expected.
        // strs.add(1);

        printLen(ints);
        printLen(strs);
    }
}
```

The code comments in the above program demonstrate that Java does not allow runtime instantiation from generic types. Which makes sense, as generics are strictly a compile-time convenience feature in Java.

To run the above program, please follow these instructions:

1. Launch the container:

    ```bash
    docker run -it --rm go-generics-the-hard-way
    ```

1. Compile the above program:

    ```bash
    javac -g ./05-internals/03-runtime-instantiation/java/main.java
    ```

1. Run the program:

    ```bash
    java -classpath ./05-internals/03-runtime-instantiation/java Main
    ```

    ```bash
    4
    2
    ```

1. Type `exit` to stop and remove the container.

It bears saying once more -- generics in Java are purely a compile-time, convenience feature. What about .NET and Go?

---

Next: [.NET](./02-dotnet.md)
