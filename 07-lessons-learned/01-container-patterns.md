# Container patterns

The type parameter proposal itself mentions the [_container_](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#containers) pattern -- where boilerplate code is repeated for different types:

```golang
type IntList []int

type Int8List []int

type MyTypeList []MyType
```

are replaced by:

```golang
type List[T any] []T
```

Soon stdlib will start to [build support](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#pervasiveness) into well-known packages such as `sort` and new packages for things like `chan` and `slice` that use generics to provide a one-stop-shop for common functionality.

The container pattern is likely the most pervasive way generics will impact Go in the immediate future.

---

Next: [Eliminating boxing](./02-eliminating-boxing.md)
