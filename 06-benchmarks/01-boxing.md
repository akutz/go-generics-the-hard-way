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

The results _should_ look something like this:

```bash
goos: linux
goarch: amd64
pkg: go-generics-the-hard-way/06-benchmarks/boxing
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkBoxedList
BenchmarkBoxedList-8   	 7492065	       144.1 ns/op	      96 B/op	       1 allocs/op
BenchmarkBoxedList-8   	 8258203	       130.9 ns/op	      88 B/op	       1 allocs/op
BenchmarkBoxedList-8   	 9099303	       150.1 ns/op	      98 B/op	       1 allocs/op
BenchmarkBoxedList-8   	 9509346	       138.7 ns/op	      94 B/op	       1 allocs/op
BenchmarkBoxedList-8   	 9163674	       120.0 ns/op	      98 B/op	       1 allocs/op
BenchmarkList
BenchmarkList-8        	44022606	        39.95 ns/op	      42 B/op	       0 allocs/op
BenchmarkList-8        	34879302	        49.81 ns/op	      43 B/op	       0 allocs/op
BenchmarkList-8        	50451991	        49.79 ns/op	      46 B/op	       0 allocs/op
BenchmarkList-8        	41202624	        34.65 ns/op	      45 B/op	       0 allocs/op
BenchmarkList-8        	42310273	        28.94 ns/op	      44 B/op	       0 allocs/op
PASS
ok  	go-generics-the-hard-way/06-benchmarks/boxing	16.969s
```

## Key takeaways

A few, key takeaways:

* On average the implementation of `List[T any]` was more performant with respect to CPU _and_ memory by 300% and 200%, respectively.
* Because no boxing was needed to store an `int` in `List[int]`, there was no need to allocate anything to box the integer added to the `List[int]`.

---

Next: [Return to beginning](../README.md)
