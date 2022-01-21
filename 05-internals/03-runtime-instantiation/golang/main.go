//go:build invalid
// +build invalid

/*
Copyright 2022

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"reflect"
)

type List[T any] []T

func (a *List[T]) add(val T) {
	*a = append(*a, val)
}

func printLen[T any](list List[T]) {
	fmt.Println(len(list))
}

func main() {
	// Declare a new List[int] normally and add some values to it.
	var ints List[int]
	ints.add(1)
	ints.add(2)
	ints.add(3)

	// Get the type of the generic List in order to build a new List[string]
	// at runtime using reflection.
	_ = reflect.TypeOf(List)

	// There is no point in going any further as the above statement simply
	// does not compile.
	//
	// Generics in Go do not exist at runtime. Like Java they are a compile-time
	// convenience feature. While we already claimed that generic types in Go
	// do *not* get type erased, that is not entirely true.
	//
	// The generic "template", in this case "List[T any]" is erased at runtime.
	//
	// The only type that exists at runtime is main.List[int], because that is
	// what was instantiated.
	//
	// The Go compiler transforms all instantiated, generic types into concrete
	// types and drops the generic "templates" from the compiled binary.
}
