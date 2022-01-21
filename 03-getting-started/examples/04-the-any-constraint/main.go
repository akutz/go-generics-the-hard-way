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
)

// Sum returns the sum of the provided arguments.
func Sum[T any](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}...))
	fmt.Println(Sum([]int8{1, 2, 3}...))
	fmt.Println(Sum([]uint32{1, 2, 3}...))
	fmt.Println(Sum([]float64{1.1, 2.2, 3.3}...))
	fmt.Println(Sum([]complex128{1.1i, 2.2i, 3.3i}...))
}
