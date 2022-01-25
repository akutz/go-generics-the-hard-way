//go:build int && int8 && !int16 && !int32 && !int64
// +build int,int8,!int16,!int32,!int64

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
	list "go-generics-the-hard-way/06-benchmarks/lists/typed"
)

func main() {
	var intList list.IntList
	intList.Add(1)

	var int8List list.Int8List
	int8List.Add(1)
}
