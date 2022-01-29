# Careful constructors

The last several pages discussed constraints based on structs:

```golang
// HasID is a structural constraint satisfied by structs with a single field
// called "ID" of type "string".
type HasID interface {
	~struct {
		ID string
	}
}
```

and interfaces:

```golang
// CanGetID is an interface constraint satisfied by a type that has a function
// with the signature "GetID() string".
type CanGetID interface {
	GetID() string
}
```

The question a lot of people are probably asking is "Why would I ever use a structural constraint when only types that match that constraint _exactly_ are permitted," ex.

```golang
// Unique satisfies the structural constraint "HasID" *and* the interface
// constraint "CanGetID."
type Unique struct {
	ID string
}

func (u Unique) GetID() string {
	return u.ID
}

// UniqueName does *not* satisfiy the structural constraint "HasID," because
// while UniqueName has the field "ID string," the type also contains the field
// "Name string."
//
// Structural constraints must match *exactly*.
//
// UniqueName *does* satisfy the interface constraint "CanGetID."
type UniqueName struct {
	Unique
	Name string
}
```

So it seems clear that interface constraints are better, right? Well...99.9% of the time they are more flexible, that is true. Better? That might be a matter of opinion. What is indisputable is that it is **much** simpler to initialize new instances of _structural contraints_ compared to interface constraints that have functions that mutate their receiver. For example, the following program behaves as it should ([Go playground](https://gotipplay.golang.org/p/Ppd5lqMbvBG)):

```golang
// NewHasT returns a new instance of T.
func NewHasT[T HasID]() T {
	// Declare a new instance of T on the stack.
	var t T

	// Return the new T by value.
	return t
}

func main() {
	fmt.Printf("%T\n", NewHasT[Unique]())
}
```

The expected output is emitted:

```bash
main.Unique
```

Just to prove the point, the above example fails to compile if the call is switched to `NewHasT[UniqueName]` ([Go playground](https://gotipplay.golang.org/p/gLi-g2eUsci)):

```bash
./prog.go:53:29: UniqueName does not implement HasID
```

Ah, so _clearly_ the answer is a generic, helper function that can be satisfied by `HasID` _and_ `CanGetID` ([Go playground](https://gotipplay.golang.org/p/gfAc_QgJfQ8)):

```golang
// NewT returns a new instance of T...
//
// ...or it would if this function were not invalid. Composite constraints
// cannot contain unions of concrete types such as Go primitive or struct
// types and interface types.
func NewT[T HasID | CanGetID]() T {
	var t T
	return t
}

func main() {
	fmt.Printf("%T\n", NewT[UniqueName]())
}
```

As the above code comment states, the function `NewT` is in fact invalid. Attempting to compile this example results in the following error:

```bash
./prog.go:48:21: cannot use main.CanGetID in union (main.CanGetID contains methods)
```

A quick way around that is to create a separate helper ([Go playground](https://gotipplay.golang.org/p/wz03yBdHyf5)):

```golang
// NewCanGetT returns a new instance of T.
func NewCanGetT[T CanGetID]() T {
	// Declare a new instance of T on the stack.
	var t T

	// Return the new T by value.
	return t
}

func main() {
	fmt.Printf("%T\n", NewCanGetT[UniqueName]())
}
```

Once again the expected output occurs:

```bash
main.UniqueName
```

Other than requiring two different helper methods for instantiating `T`, there does not seem to be any discernable difference between structural and interface constraints, right? Not so fast. Consider, for a moment, _this_ constraint:

```golang
// CanSetID is an interface constraint satisfied by a type that has a function
// with the signature "SetID(string)".
type CanSetID interface {
	SetID(string)
}
```

Let's go ahead and implement this on `UniqueName`:

```golang
func (u *UniqueName) SetID(s string) {
	u.ID = s
}
```

Note the function receiver is _by address_, not by value as was the case for `Unique.GetID`. This is because 99.99999% of the time a function that mutates something on its receiver intends to mutate the instance of that type, not a copy. So what do you think the outcome of the following example will be ([Go playground](https://gotipplay.golang.org/p/c_dKfiBnYFc))?

```golang
// NewCanSetT returns a new instance of T.
func NewCanSetT[T CanSetID]() T {
	// Declare a new instance of T. Because T is constrained to be a
	// concrete type, it can easily be declared on the stack.
	var t T

	// Return the new T by value.
	return t
}

func main() {
	fmt.Printf("%T\n", NewCanSetT[UniqueName]())
}
```

That's right, a compiler error:

```bash
./prog.go:81:32: UniqueName does not implement CanSetID: wrong method signature
	got  func (*UniqueName).SetID(s string)
	want func (CanSetID).SetID(string)
```

Remember, it is not `UniqueName` that implements `CanSetT`, but rather `*UniqueName` ([Go playground](https://gotipplay.golang.org/p/o6JRltQHNYQ)):

```golang
func main() {
	fmt.Printf("%T\n", NewCanSetT[*UniqueName]())
}
```

Now the expected output occurs:

```bash
*main.UniqueName
```

Hmm, that's odd, a pointer to a `UniqueName` was printed, not `UniqueName`? That's because that is what was provided as `T` -- `*UniqueName`. And not just that, but what happens if the value to which the pointer refers is printed ([Go playground](https://gotipplay.golang.org/p/bvAi8Ax0Q9f)):

```golang
func main() {
	fmt.Printf("%T\n", *(NewCanSetT[*UniqueName]()))
}
```

Unfortunately a nil pointer exception (NPE) occurs:

```bash
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x47eb7c]

goroutine 1 [running]:
main.main()
	/tmp/sandbox3877304423/prog.go:81 +0x1c
```

This is because a new instance of a _pointer_ to a `UniqueName` was declared and initialized, and it is in fact possible to declare and initialize a pointer to a concrete type without actually initializing the underlying concrete type itself.

Hopefully the reason for structural constraints is now a little more clear. Instead of using an interface constraint with `SetID(string)` for assigning a value, a structural constraint would simultaneously enable:

* initializing a new instance of a generic type
* assigning values to the fields of the new instance

Bear in mind this could be considered an edge case, but it is worth demonstrating. For what it is worth, it _is_ possible to figure out `T` from `*T` using the `reflect` package, but without _extreme_ care this can _easily_ lead to the violation of the runtime type safety provided by generics and should be avoided whenever possible. However, for those inquiring minds that want to know... ([Golang playground](https://gotipplay.golang.org/p/xOW18VTorJS)):

```golang
// NewCanSetT returns a new instance of T.
func NewCanSetT[T CanSetID]() T {
	// Declare a new instance of T on the stack.
	//
	// In 99.9999999% of the cases this will be a new pointer that
	// does not address a valid value. This is the common case because
	// setter functions are most often defined as receivers by-address.
	//
	// It is possible to write a *lot* of code to figure out what T
	// actually is, but that is not the purpose of this example, and
	// quite frankly that type of code should be avoided at all costs.
	//
	// At the moment, since we know "NewCanSetT[T CanSetID]()" was invoked
	// where T is "*UniqueName", we know that "t" below is a new instance
	// of a "*UniqueName" that is currently nil.
	var t T

	// Again, we know t is a pointer that is set to nil, and in order
	// to initialize a new value for the pointer to address, we need to
	// figure out what T is in *T (we know it is "UniqueName," but in
	// the wild we will not have this luxury).
	//
	// To figure this out we can use the "reflect" package to get the
	// type of T and then get the type of the element addressed by T
	// (since again, we know T is a pointer).
	typeOfValueAddressedByT := reflect.TypeOf(t).Elem()

	// So now typeOfValueAddressedByT represents the underlying type
	// addressed by the pointer. That means we can use the "reflect.New"
	// function to create a new instance of this type.
	newValueOfTypeOfValueAddressedByT := reflect.New(typeOfValueAddressedByT)

	// If it was not dangerous enough already, here is where it will
	// absolutely result in a panic if things are not as expected.
	// We need to assert the type of newValueOfTypeOfValueAddressedByT
	// in order to assign its address to t.
	t = newValueOfTypeOfValueAddressedByT.Interface().(T)

	// Finally, return t.
	return t
}

func main() {
	fmt.Printf("%T\n", *(NewCanSetT[*UniqueName]()))
}
```

The above hack _does_ print the expected output:

```bash
main.UniqueName
```

FWIW, the above hack can also be simplified ([Golang playground](https://gotipplay.golang.org/p/20yKw_LgiVd)):

```golang
// NewCanSetT returns a new instance of T.
func NewCanSetT[T CanSetID]() T {
	return reflect.New(reflect.TypeOf(*(new(T))).Elem()).Interface().(T)
}

func main() {
	fmt.Printf("%T\n", *(NewCanSetT[*UniqueName]()))
}
```

Still, please do not do this. Instead, if you have made it this far, why not spend your time by delving deeper into generics and Go!

---

Next: [Internals](../05-internals/)
