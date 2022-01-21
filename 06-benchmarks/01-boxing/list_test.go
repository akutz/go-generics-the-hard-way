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

package boxing_test

import (
	"testing"

	boxing "go-generics-the-hard-way/06-benchmarks/01-boxing"
)

func BenchmarkBoxedList(b *testing.B) {
	var list boxing.BoxedList
	for i := 0; i < b.N; i++ {
		list = append(list, i)
	}
}

func BenchmarkList(b *testing.B) {
	var list boxing.List[int]
	for i := 0; i < b.N; i++ {
		list = append(list, i)
	}
}
