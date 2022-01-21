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

package example_test

import (
	"fmt"
)

type request struct {
	host *string
	port *int
}

func print(r request) {
	fmt.Print("request: host=")
	if r.host != nil {
		fmt.Print(*r.host)
	}
	fmt.Print(", port=")
	if r.port != nil {
		fmt.Printf("%d", *r.port)
	}
	fmt.Println()
}

func ExampleTheProblem() {
	print(request{
		host: nil, // needs a *int
		port: nil, // needs a *string
	})

	// Output: request: host=, port=
}

func ExampleLocalVars() {
	// Declare "host" and "port" in order to create pointers to satisfy the
	// fields in the "request" struct.
	host, port := "local", 80

	print(request{
		host: &host,
		port: &port,
	})

	// Output: request: host=local, port=80
}

// PtrInt returns *i.
func PtrInt(i int) *int {
	return &i
}

// PtrStr returns *s.
func PtrStr(s string) *string {
	return &s
}

func ExampleTypedHelpers() {
	// Use the two helper functions that return pointers to their provided
	// values. Remember, this pattern must scale with the number of distinct,
	// defined types that need to be passed by pointer instead of value.
	print(request{
		host: PtrStr("local"),
		port: PtrInt(80),
	})

	// Output: request: host=local, port=80
}

// Ptr returns *value.
func Ptr[T any](value T) *T {
	return &value
}

func ExampleGenericSolution() {
	// No local variables and the typed helper functions can be collapsed into
	// a single, generic function for getting a pointer to a value.
	print(request{
		host: Ptr("local"),
		port: Ptr(80),
	})

	// Output: request: host=local, port=80
}
