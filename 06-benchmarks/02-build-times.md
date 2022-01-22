# Build times

The Go compiler essentially transforms every instance of a generic into a defined, named type, and that additional work _must_ impact build times, right? This page describes how to run `BenchmarkBuildTimes`, a Go-based benchmark that builds the following packages:

* **`./lists/boxed`**: defines `type List []interface{}`
* **`./lists/typed`**: defines `type ListInt []int` and variants for `int8`, `int16`, `int32`, and `int64`
* **`./lists/generic`**: defines `type List[T any] []T`

The `typed` and `generic` packages are also subject to the following build tags:

* **`int`**: activates that package's list of `int`
* **`int8`**: activates that package's list of `int8`
* **`int16`**: activates that package's list of `int16`
* **`int32`**: activates that package's list of `int32`
* **`int64`**: activates that package's list of `int64`

The following benchmarks are defined:

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

With this in mind, the benchmark may be run using the following command:

```bash
docker run -it --rm go-generics-the-hard-way \
  go test -tags benchmarks -bench BuildTimes -run BuildTimes -count 5 -v ./06-benchmarks
```

The output _should_ looks similar to the following:

```bash
goos: linux
goarch: amd64
pkg: go-generics-the-hard-way/06-benchmarks
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkBuildTimes
BenchmarkBuildTimes/BoxedList
BenchmarkBuildTimes/BoxedList-8         	      72	  16941312 ns/op
BenchmarkBuildTimes/BoxedList-8         	      55	  18869829 ns/op
BenchmarkBuildTimes/BoxedList-8         	      73	  17516405 ns/op
BenchmarkBuildTimes/BoxedList-8         	      76	  16455497 ns/op
BenchmarkBuildTimes/BoxedList-8         	      66	  17547718 ns/op
BenchmarkBuildTimes/GenericList
BenchmarkBuildTimes/GenericList/1-type
BenchmarkBuildTimes/GenericList/1-type-8         	      58	  19648051 ns/op
BenchmarkBuildTimes/GenericList/1-type-8         	      63	  19293062 ns/op
BenchmarkBuildTimes/GenericList/1-type-8         	      63	  19299799 ns/op
BenchmarkBuildTimes/GenericList/1-type-8         	      57	  19215178 ns/op
BenchmarkBuildTimes/GenericList/1-type-8         	      63	  19148275 ns/op
BenchmarkBuildTimes/GenericList/2-types
BenchmarkBuildTimes/GenericList/2-types-8        	      55	  20769417 ns/op
BenchmarkBuildTimes/GenericList/2-types-8        	      58	  21225023 ns/op
BenchmarkBuildTimes/GenericList/2-types-8        	      62	  20908756 ns/op
BenchmarkBuildTimes/GenericList/2-types-8        	      61	  20618478 ns/op
BenchmarkBuildTimes/GenericList/2-types-8        	      49	  20609018 ns/op
BenchmarkBuildTimes/GenericList/3-types
BenchmarkBuildTimes/GenericList/3-types-8        	      58	  22249882 ns/op
BenchmarkBuildTimes/GenericList/3-types-8        	      50	  21661469 ns/op
BenchmarkBuildTimes/GenericList/3-types-8        	      60	  20845732 ns/op
BenchmarkBuildTimes/GenericList/3-types-8        	      57	  21294356 ns/op
BenchmarkBuildTimes/GenericList/3-types-8        	      55	  21751211 ns/op
BenchmarkBuildTimes/GenericList/4-types
BenchmarkBuildTimes/GenericList/4-types-8        	      52	  21766629 ns/op
BenchmarkBuildTimes/GenericList/4-types-8        	      54	  21698073 ns/op
BenchmarkBuildTimes/GenericList/4-types-8        	      51	  22035912 ns/op
BenchmarkBuildTimes/GenericList/4-types-8        	      57	  21505650 ns/op
BenchmarkBuildTimes/GenericList/4-types-8        	      55	  21903746 ns/op
BenchmarkBuildTimes/GenericList/5-types
BenchmarkBuildTimes/GenericList/5-types-8        	      52	  22943731 ns/op
BenchmarkBuildTimes/GenericList/5-types-8        	      55	  22874752 ns/op
BenchmarkBuildTimes/GenericList/5-types-8        	      55	  22970808 ns/op
BenchmarkBuildTimes/GenericList/5-types-8        	      52	  23217451 ns/op
BenchmarkBuildTimes/GenericList/5-types-8        	      49	  22741901 ns/op
BenchmarkBuildTimes/TypedList
BenchmarkBuildTimes/TypedList/1-type
BenchmarkBuildTimes/TypedList/1-type-8           	      68	  17433482 ns/op
BenchmarkBuildTimes/TypedList/1-type-8           	      67	  17455819 ns/op
BenchmarkBuildTimes/TypedList/1-type-8           	      69	  17374573 ns/op
BenchmarkBuildTimes/TypedList/1-type-8           	      70	  18940335 ns/op
BenchmarkBuildTimes/TypedList/1-type-8           	      61	  19109838 ns/op
BenchmarkBuildTimes/TypedList/2-types
BenchmarkBuildTimes/TypedList/2-types-8          	      57	  20981557 ns/op
BenchmarkBuildTimes/TypedList/2-types-8          	      58	  19088787 ns/op
BenchmarkBuildTimes/TypedList/2-types-8          	      66	  19814681 ns/op
BenchmarkBuildTimes/TypedList/2-types-8          	      63	  19388103 ns/op
BenchmarkBuildTimes/TypedList/2-types-8          	      60	  19494429 ns/op
BenchmarkBuildTimes/TypedList/3-types
BenchmarkBuildTimes/TypedList/3-types-8          	      54	  19548919 ns/op
BenchmarkBuildTimes/TypedList/3-types-8          	      61	  19365266 ns/op
BenchmarkBuildTimes/TypedList/3-types-8          	      63	  20083666 ns/op
BenchmarkBuildTimes/TypedList/3-types-8          	      58	  20796191 ns/op
BenchmarkBuildTimes/TypedList/3-types-8          	      55	  20319472 ns/op
BenchmarkBuildTimes/TypedList/4-types
BenchmarkBuildTimes/TypedList/4-types-8          	      54	  21392083 ns/op
BenchmarkBuildTimes/TypedList/4-types-8          	      54	  20806404 ns/op
BenchmarkBuildTimes/TypedList/4-types-8          	      62	  20343034 ns/op
BenchmarkBuildTimes/TypedList/4-types-8          	      61	  20599776 ns/op
BenchmarkBuildTimes/TypedList/4-types-8          	      58	  20353588 ns/op
BenchmarkBuildTimes/TypedList/5-types
BenchmarkBuildTimes/TypedList/5-types-8          	      52	  20586529 ns/op
BenchmarkBuildTimes/TypedList/5-types-8          	      61	  21168941 ns/op
BenchmarkBuildTimes/TypedList/5-types-8          	      51	  21378163 ns/op
BenchmarkBuildTimes/TypedList/5-types-8          	      55	  21403382 ns/op
BenchmarkBuildTimes/TypedList/5-types-8          	      62	  20776464 ns/op
PASS
ok  	go-generics-the-hard-way/06-benchmarks	67.411s
```

The Go compiler appears to be incredibly efficient at stencling the generic types together into concrete types as the build times for the generic lists do not demonstrate any distinguishable difference from the typed lists.

---

Next: [File sizes](./03-file-sizes.md)
