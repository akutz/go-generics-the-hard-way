# Summary

This page is a brief summary of how generics are implemented in Java, .NET, and Golang.

|      | Compile-time type safety | Type erasure | Runtime type safety | Runtime instantiation |
|------|:---:|:---:|:---:|:---:|
| Java |  ✓  |  ✓  |     |     |
| .NET |  ✓  |     |  ✓  |  ✓  |
|  Go  |  ✓  |     |  ✓  |     |

Microsoft .NET by far has the most advanced implementation of generics, but there is much to appreciate about the simpler, more elegant approach adopted by Golang.

So now that we know how Go implements generics, how do they perform?

---

Next: [Benchmarks](../06-benchmarks/)
