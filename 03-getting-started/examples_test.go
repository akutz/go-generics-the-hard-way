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
	"os"
	"os/exec"
)

func ExampleWhatIsAGeneric_SumInt() {
	fmt.Println(SumInt(1, 2, 3))
	// Output: 6
}

func ExampleWhatIsAGeneric_EmptyInterface() {
	print := func(i interface{}, err error) {
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i)
		}
	}
	print(SumInterface(int64(1), int64(2), int64(3)))
	print(SumInterface(uint32(1), int(2), uint32(3)))
	print(SumInterface(float64(1), int(2), uint32(3)))
	print(SumInterface(uint32(1), uint32(2), uint32(3)))
	// Output:
	// 6
	// int is not supported
	// float64 is not supported
	// 6
}

func ExampleWhatIsAGeneric_MultipleFunctions() {
	fmt.Println(SumInt64(1, 2, 3))
	fmt.Println(SumUint32(1, 2, 3))
	// Output:
	// 6
	// 6
}

func ExampleSyntax() {
	fmt.Println(SumTint(1, 2, 3))
	// Output: 6
}

func ExampleConstraints() {
	fmt.Println(SumTintOrint64(1, 2, 3))
	fmt.Println(SumTintOrint64(int64(1), 2, 3))
	// Output:
	// 6
	// 6
}

func ExampleTheAnyConstraint() {
	cmd := exec.Command(
		"go",
		"build",
		/*
			The `go build` command below uses the flag `-tags invalid`. This
			flag instructs Go to include the file
			`./examples/04-the-any-constraint/main.go` in the compilation.
			Normally the file has a build constraint that prevents it from being
			considered to prevent:

			  * Dev tools (ex. an IDE such as VS Code) from constantly warning
			    this file is in error
			  * a failed `go test ./...` command from the root of this
			    repository

			However, the entire point of this example is to demonstrate that
			failure, so the file needs to be "activated."
		*/
		"-tags", "invalid",
		"./examples/04-the-any-constraint/",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	fmt.Println(cmd.Run())

	// Output:
	// # go-generics-the-hard-way/03-getting-started/examples/04-the-any-constraint
	// examples/04-the-any-constraint/main.go:30:3: invalid operation: operator + not defined on sum (variable of type T constrained by any)
	// exit status 2
}

// SumInt returns the sum of the provided arguments.
func SumInt(args ...int) int {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// SumInterface returns the sum of the provided arguments.
//
// An error is returned if the arguments are not of a supported
// type or of mixed types.
func SumInterface(args ...interface{}) (interface{}, error) {

	var sum interface{}

	for i := 0; i < len(args); i++ {
		switch a := args[i].(type) {
		case int64:
			if sum == nil {
				sum = int64(0)
			}

			tsum, ok := sum.(int64)
			if !ok {
				return nil, fmt.Errorf("previous arg was not an int64")
			}

			tsum += a
			sum = tsum

		case uint32:
			if sum == nil {
				sum = uint32(0)
			}

			tsum, ok := sum.(uint32)
			if !ok {
				return nil, fmt.Errorf("previous arg was not an uint32")
			}

			tsum += a
			sum = tsum
		default:
			return nil, fmt.Errorf("%T is not supported", args[i])
		}
	}
	return sum, nil
}

// SumInt64 returns the sum of the provided int64 values.
func SumInt64(args ...int64) int64 {
	var sum int64
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// SumUint32 returns the sum of the provided uint32 values.
func SumUint32(args ...uint32) uint32 {
	var sum uint32
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// SumTint returns the sum of the provided arguments.
func SumTint[T int](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// SumTintOrint64 returns the sum of the provided arguments.
func SumTintOrint64[T int | int64](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
