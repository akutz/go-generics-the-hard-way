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
	"runtime"
)

type List[T any] []T

func (a *List[T]) add(val T) {
	*a = append(*a, val)
}

func printLen[T any](list List[T]) {
	fmt.Println(len(list))
}

func main() {
	var ints List[int]
	ints.add(1)
	ints.add(2)
	ints.add(3)

	var strs List[string]
	strs.add("Hello")
	strs.add("world")

	printLen(ints)
	printLen(strs)

	runtime.Breakpoint()
}
