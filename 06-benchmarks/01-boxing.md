# Boxing

This benchmark demonstrates how Go generics speedup scenarios where boxing can be avoided.

* [**What is boxing?**](#what-is-boxing?): a summary of boxing
* [**The example**](#the-example): an exemplar case
* [**The benchmark**](#the-benchmark): how does it stack up?
* [**Key takeaways**](#key-takeaways): what it all means

## What is boxing?

In general, the term _boxing_ refers to wrapping a value with some container, aka, "putting it inside a box." For the purposes of _this_ page, the term "boxing" will refer to wrapping some value in Golang with the empty interface, ex.:

```go
x := int(0)
i := interface{}(x)
```

Generics can be used to eliminate the need for boxing in many situations, which we will highlight with `List[T]`.

## The example

```golang
// BoxedList is a new type definition for []interface{}.
type BoxedList []interface{}

// Add a new element to the list.
//
// Please note that val will be boxed in order to pass it into the method using
// the empty interface.
func (a *BoxedList) Add(val interface{}) {
	*a = append(*a, val)
}

// List is a new type definition for []T.
type List[T any] []T

// Add a new element to the list.
//
// Please note that val will not be boxed since it will be passed as its actual
// type.
func (a *List[T]) Add(val T) {
	*a = append(*a, val)
}
```

## The benchmark

In order to determine the impact produced by the lack of boxing, we will run some Go benchmarks that add random integers to a `BoxedList` and an `List[int]`:

```bash
docker run -it --rm go-generics-the-hard-way \
  go test -bench . -benchmem -count 5 -v ./06-benchmarks/boxing/
```

---

:warning: **Please note** that Docker may hang if not provided enough resources. The performance improvements are significant enough that the container may be starved for CPU and killed.

---


**Run via docker with 14cpu and 6GiB**

```bash
goos: linux
goarch: amd64
pkg: go-generics-the-hard-way/06-benchmarks/boxing
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkBoxedList
BenchmarkBoxedList-14    	 8322405	       148.7 ns/op	     107 B/op	       0 allocs/op
BenchmarkBoxedList-14    	10976193	       115.0 ns/op	     102 B/op	       0 allocs/op
BenchmarkBoxedList-14    	11569058	       125.3 ns/op	      97 B/op	       0 allocs/op
BenchmarkBoxedList-14    	 9299853	       123.6 ns/op	      96 B/op	       0 allocs/op
BenchmarkBoxedList-14    	 9122316	       132.0 ns/op	      98 B/op	       0 allocs/op
BenchmarkList
BenchmarkList-14         	100000000	        10.52 ns/op	      45 B/op	       0 allocs/op
BenchmarkList-14         	163110246	         8.399 ns/op	      43 B/op	       0 allocs/op
BenchmarkList-14         	202797236	        12.38 ns/op	      44 B/op	       0 allocs/op
BenchmarkList-14         	82267455	        18.61 ns/op	      44 B/op	       0 allocs/op
BenchmarkList-14         	166848986	        12.94 ns/op	      42 B/op	       0 allocs/op
PASS
ok  	go-generics-the-hard-way/06-benchmarks/boxing	19.785s
```

**Run natively**

```bash
goos: darwin
goarch: amd64
pkg: go-generics-the-hard-way/06-benchmarks/boxing
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkBoxedList
BenchmarkBoxedList-16    	16099449	        85.37 ns/op	      88 B/op	       0 allocs/op
BenchmarkBoxedList-16    	13072742	        99.08 ns/op	     106 B/op	       0 allocs/op
BenchmarkBoxedList-16    	14000494	        90.01 ns/op	     100 B/op	       0 allocs/op
BenchmarkBoxedList-16    	16146943	        78.90 ns/op	      88 B/op	       0 allocs/op
BenchmarkBoxedList-16    	13265748	        86.84 ns/op	     105 B/op	       0 allocs/op
BenchmarkList
BenchmarkList-16         	100000000	        15.63 ns/op	      45 B/op	       0 allocs/op
BenchmarkList-16         	297442202	        24.31 ns/op	      47 B/op	       0 allocs/op
BenchmarkList-16         	259559503	        17.17 ns/op	      43 B/op	       0 allocs/op
BenchmarkList-16         	256879438	         5.537 ns/op	      43 B/op	       0 allocs/op
BenchmarkList-16         	276920250	         3.819 ns/op	      40 B/op	       0 allocs/op
PASS
ok  	go-generics-the-hard-way/06-benchmarks/boxing	25.895s
```

## Key takeaways

A few, key takeaways:

* On average the implementation of `List[T any]` was more performant:
  * operations were 10x faster
  * consumed half the memory
* The performance improvements were the result of removing the need to box the integer values

---

Next: [Return to beginning](../README.md)
