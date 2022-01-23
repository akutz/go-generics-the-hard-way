# File sizes

---

:warning: **Please note** that it is likely the discrepancy in file sizes is due to [golang/go#50438](https://github.com/golang/go/issues/50438). This page will be updated with more information as soon as possible. Thanks!

---

While generics did not seem to have demonstrable impact on build times based on the limited sampling, the same cannot be said for file sizes. This page describes how to run `TestFileSizes`, a Go-based test that builds the following packages:

* **`./lists/boxed`**: defines `type List []interface{}`
* **`./lists/typed`**: defines `type ListInt []int` and variants for `int8`, `int16`, `int32`, and `int64`
* **`./lists/generic`**: defines `type List[T any] []T`

The `typed` and `generic` packages are also subject to the following build tags:

* **`int`**: activates that package's list of `int`
* **`int8`**: activates that package's list of `int8`
* **`int16`**: activates that package's list of `int16`
* **`int32`**: activates that package's list of `int32`
* **`int64`**: activates that package's list of `int64`

The following tests are defined:

* build the boxed list
* build the typed list:
  * for `int`
  * for `int` and `int8`
  * for `int`, `int8`, and `int16`
  * for `int`, `int8`, `int16`, and `int32`
  * for `int`, `int8`, `int16`, `int32`, and `int64`
* build the generic list:
  * for `int`
  * for `int` and `int8`
  * for `int`, `int8`, and `int16`
  * for `int`, `int8`, `int16`, and `int32`
  * for `int`, `int8`, `int16`, `int32`, and `int64`

At the end of each test the size of the produced, binary artifact is printed.

With this in mind, the benchmark may be run using the following command:

```bash
docker run -it --rm go-generics-the-hard-way \
  go test -tags benchmarks -run FileSizes -count 1 -v ./06-benchmarks
```

The output _should_ looks similar to the following:

```bash
=== RUN   TestFileSizes
=== RUN   TestFileSizes/Boxed
    benchmarks_test.go:135: file size: 6806
=== RUN   TestFileSizes/Generic
=== RUN   TestFileSizes/Generic/1-type
    benchmarks_test.go:135: file size: 11334
=== RUN   TestFileSizes/Generic/2-types
    benchmarks_test.go:135: file size: 19818
=== RUN   TestFileSizes/Generic/3-types
    benchmarks_test.go:135: file size: 28338
=== RUN   TestFileSizes/Generic/4-types
    benchmarks_test.go:135: file size: 36858
=== RUN   TestFileSizes/Generic/5-types
    benchmarks_test.go:135: file size: 45294
=== RUN   TestFileSizes/Typed
=== RUN   TestFileSizes/Typed/1-type
    benchmarks_test.go:135: file size: 6716
=== RUN   TestFileSizes/Typed/2-types
    benchmarks_test.go:135: file size: 11134
=== RUN   TestFileSizes/Typed/3-types
    benchmarks_test.go:135: file size: 15588
=== RUN   TestFileSizes/Typed/4-types
    benchmarks_test.go:135: file size: 20042
=== RUN   TestFileSizes/Typed/5-types
    benchmarks_test.go:135: file size: 24496
--- PASS: TestFileSizes (0.26s)
    --- PASS: TestFileSizes/Boxed (0.04s)
    --- PASS: TestFileSizes/Generic (0.12s)
        --- PASS: TestFileSizes/Generic/1-type (0.02s)
        --- PASS: TestFileSizes/Generic/2-types (0.02s)
        --- PASS: TestFileSizes/Generic/3-types (0.03s)
        --- PASS: TestFileSizes/Generic/4-types (0.03s)
        --- PASS: TestFileSizes/Generic/5-types (0.02s)
    --- PASS: TestFileSizes/Typed (0.10s)
        --- PASS: TestFileSizes/Typed/1-type (0.02s)
        --- PASS: TestFileSizes/Typed/2-types (0.02s)
        --- PASS: TestFileSizes/Typed/3-types (0.02s)
        --- PASS: TestFileSizes/Typed/4-types (0.02s)
        --- PASS: TestFileSizes/Typed/5-types (0.02s)
PASS
ok  	go-generics-the-hard-way/06-benchmarks	0.261s
```

It appears that binary artifacts built from generic code are consistently close to twice as large versus directly implementing type-variants for common patterns (such as variants of a `List` type):

| Number of types | Size (bytes) of artifact w typed lists | Size (bytes) of artifact w generic lists | Difference (bytes) | Difference (%) |
|:---:|:---:|:---:|:---:|:---:|
| 1 | 6716 | 11334 | 4618 | 59.25% |
| 2 | 11134 | 19818 | 8684 | 56.18% |
| 3 | 15588 | 28338 | 12750 | 55.00% |
| 4 | 20042 | 36858 | 16816 | 54.37% |
| 5 | 24496 | 45294 | 20798 | 54.08% |

---

Next: _That's it, thank you for reading!_
