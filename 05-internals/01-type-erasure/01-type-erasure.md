# Type erasure

Wikipedia [defines](https://en.wikipedia.org/wiki/Type_erasure) type erasure as the following:

> the load-time process by which explicit type annotations are removed from a program before it is executed at run-time

In other words, let us pretend there are two variables defined per the following pseudo-code:

```
var ints = List<Int32>{1, 2, 3}
var strs = List<String>{"Hello", "world"}
```

The above pseudo-code makes it perfectly clear that `ints` is a list of 32-bit integers and `strs` is a list of string values. Languages that practice type erasure would transform the above pseudo-code to looking something a litle more like this at runtime:

```
var ints = List{1, 2, 3}
var strs = List{"Hello", "world"}
```

At runtime there is nothing to distinguish the type of `ints` from `strs` without looking at the types of the elements in those lists. This is a very common example of type erasure, one which is present in a very popular programming language. Let's see just what impact that has by continuing to the next page...

---

Next: [Java](./02-java.md)
