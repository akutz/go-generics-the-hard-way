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
  go test -tags benchmarks -bench Boxing -run Boxing -benchmem -count 5 -v ./06-benchmarks
```

---

:warning: **Please note** that Docker may hang if not provided enough resources. The performance improvements are significant enough that the container may be starved for CPU and killed.

---


**Run via docker with 14cpu and 6GiB**

```bash
goos: linux
goarch: amd64
pkg: go-generics-the-hard-way/06-benchmarks
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkBoxing
BenchmarkBoxing/BoxedList
BenchmarkBoxing/BoxedList-14         	 8863634	       145.5 ns/op	     101 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-14         	12885366	       112.7 ns/op	      88 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-14         	 8696394	       132.6 ns/op	     103 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-14         	11669534	       131.5 ns/op	      96 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-14         	10448851	       141.4 ns/op	     106 B/op	       0 allocs/op
BenchmarkBoxing/GenericList
BenchmarkBoxing/GenericList-14       	178167445	        27.99 ns/op	      40 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-14       	152708895	        12.67 ns/op	      46 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-14       	180382786	        25.56 ns/op	      49 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-14       	78733688	        21.31 ns/op	      46 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-14       	170434779	         6.577 ns/op	      42 B/op	       0 allocs/op
BenchmarkBoxing/TypedList
BenchmarkBoxing/TypedList-14         	200555131	        14.78 ns/op	      44 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-14         	197507965	        27.01 ns/op	      45 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-14         	 8637734	       132.5 ns/op	      45 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-14         	100000000	        32.71 ns/op	      45 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-14         	164098374	        13.31 ns/op	      43 B/op	       0 allocs/op
PASS
ok  	go-generics-the-hard-way/06-benchmarks	50.271s
```

**Run natively**

```bash
goos: darwin
goarch: amd64
pkg: go-generics-the-hard-way/06-benchmarks
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkBoxing
BenchmarkBoxing/BoxedList
BenchmarkBoxing/BoxedList-16         	11618419	        95.01 ns/op	      96 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-16         	13484241	        91.64 ns/op	     103 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-16         	13518324	        93.74 ns/op	     103 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-16         	13723294	        95.55 ns/op	     102 B/op	       0 allocs/op
BenchmarkBoxing/BoxedList-16         	14464362	        95.60 ns/op	      97 B/op	       0 allocs/op
BenchmarkBoxing/GenericList
BenchmarkBoxing/GenericList-16       	100000000	        40.08 ns/op	      45 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-16       	135503811	         7.548 ns/op	      42 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-16       	292832476	        14.83 ns/op	      47 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-16       	268258267	         5.164 ns/op	      41 B/op	       0 allocs/op
BenchmarkBoxing/GenericList-16       	275377506	         3.793 ns/op	      40 B/op	       0 allocs/op
BenchmarkBoxing/TypedList
BenchmarkBoxing/TypedList-16         	298521306	         9.057 ns/op	      46 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-16         	286382421	         4.757 ns/op	      48 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-16         	309414636	         6.279 ns/op	      45 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-16         	278727175	         3.962 ns/op	      40 B/op	       0 allocs/op
BenchmarkBoxing/TypedList-16         	299711601	         4.025 ns/op	      46 B/op	       0 allocs/op
PASS
ok  	go-generics-the-hard-way/06-benchmarks	33.954s
```

## Key takeaways

A few, key takeaways:

* On average the implementation of `List[T any]` was more performant than the boxed list:
  * operations were 10x faster
  * consumed half the memory
* The performance improvements were the result of removing the need to box the integer values

---

Next: [Build times](./02-build-times.md)
