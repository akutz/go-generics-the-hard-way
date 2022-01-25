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

package benchmarks_test

import (
	"os"
	"os/exec"
	"testing"

	blist "go-generics-the-hard-way/06-benchmarks/lists/boxed"
	glist "go-generics-the-hard-way/06-benchmarks/lists/generic"
	tlist "go-generics-the-hard-way/06-benchmarks/lists/typed"
)

func BenchmarkBoxing(b *testing.B) {
	b.Run("boxed", func(b *testing.B) {
		var list blist.List
		for i := 0; i < b.N; i++ {
			list = append(list, i)
		}
	})
	b.Run("generic", func(b *testing.B) {
		var list glist.List[int]
		for i := 0; i < b.N; i++ {
			list = append(list, i)
		}
	})
	b.Run("typed", func(b *testing.B) {
		var list tlist.IntList
		for i := 0; i < b.N; i++ {
			list = append(list, i)
		}
	})
}

var (
	tags0Types = "no_int"
	tags1Types = "int"
	tags2Types = "int,int8"
	tags3Types = "int,int8,int16"
	tags4Types = "int,int8,int16,int32"
	tags5Types = "int,int8,int16,int32,int64"
)

type gobuildTestCase struct {
	name     string
	listType string
	testGrps []gobuildSubTestGroup
}

type gobuildSubTestGroup struct {
	name     string
	fileType string
	subTests []gobuildSubTestCase
}

type gobuildSubTestCase struct {
	name     string
	args     []string
	tags     []string
	filePath string
	fileSize int64
}

func BenchmarkGoBuild(b *testing.B) {
	testCases := []gobuildTestCase{
		{
			name:     "boxed",
			listType: "boxed",
			testGrps: []gobuildSubTestGroup{
				{
					name:     "bin",
					fileType: ".bin",
					subTests: []gobuildSubTestCase{
						{
							name:     "empty interface",
							args:     []string{"build", "-a", "-o", "boxed.bin", "./lists/boxed/cmd"},
							filePath: "boxed.bin",
						},
					},
				},
				{
					name:     "pkg",
					fileType: ".a",
					subTests: []gobuildSubTestCase{
						{
							name:     "empty interface",
							args:     []string{"build", "-a", "-o", "boxed.a", "./lists/boxed/"},
							filePath: "boxed.a",
						},
					},
				},
			},
		},
		{
			name:     "generic",
			listType: "generic",
			testGrps: []gobuildSubTestGroup{
				{
					name:     "bin",
					fileType: ".bin",
					subTests: []gobuildSubTestCase{
						{
							name:     "0-types",
							args:     []string{"build", "-a", "-tags", tags0Types, "-o", "generic-0-types.bin", "./lists/generic/cmd/"},
							filePath: "generic-0-types.bin",
						},
						{
							name:     "1-types",
							args:     []string{"build", "-a", "-tags", tags1Types, "-o", "generic-1-types.bin", "./lists/generic/cmd/"},
							filePath: "generic-1-types.bin",
						},
						{
							name:     "2-types",
							args:     []string{"build", "-a", "-tags", tags2Types, "-o", "generic-2-types.bin", "./lists/generic/cmd/"},
							filePath: "generic-2-types.bin",
						},
						{
							name:     "3-types",
							args:     []string{"build", "-a", "-tags", tags3Types, "-o", "generic-3-types.bin", "./lists/generic/cmd/"},
							filePath: "generic-3-types.bin",
						},
						{
							name:     "4-types",
							args:     []string{"build", "-a", "-tags", tags4Types, "-o", "generic-4-types.bin", "./lists/generic/cmd/"},
							filePath: "generic-4-types.bin",
						},
						{
							name:     "5-types",
							args:     []string{"build", "-a", "-tags", tags5Types, "-o", "generic-5-types.bin", "./lists/generic/cmd/"},
							filePath: "generic-5-types.bin",
						},
					},
				},
				{
					name:     "pkg",
					fileType: ".a",
					subTests: []gobuildSubTestCase{
						{
							name:     "0-types",
							args:     []string{"build", "-a", "-tags", tags0Types, "-o", "generic-0-types.a", "./lists/generic/"},
							filePath: "generic-0-types.a",
						},
						{
							name:     "1-types",
							args:     []string{"build", "-a", "-tags", tags1Types, "-o", "generic-1-types.a", "./lists/generic/"},
							filePath: "generic-1-types.a",
						},
						{
							name:     "2-types",
							args:     []string{"build", "-a", "-tags", tags2Types, "-o", "generic-2-types.a", "./lists/generic/"},
							filePath: "generic-2-types.a",
						},
						{
							name:     "3-types",
							args:     []string{"build", "-a", "-tags", tags3Types, "-o", "generic-3-types.a", "./lists/generic/"},
							filePath: "generic-3-types.a",
						},
						{
							name:     "4-types",
							args:     []string{"build", "-a", "-tags", tags4Types, "-o", "generic-4-types.a", "./lists/generic/"},
							filePath: "generic-4-types.a",
						},
						{
							name:     "5-types",
							args:     []string{"build", "-a", "-tags", tags5Types, "-o", "generic-5-types.a", "./lists/generic/"},
							filePath: "generic-5-types.a",
						},
					},
				},
			},
		},
		{
			name:     "typed",
			listType: "typed",
			testGrps: []gobuildSubTestGroup{
				{
					name:     "bin",
					fileType: ".bin",
					subTests: []gobuildSubTestCase{
						{
							name:     "0-types",
							args:     []string{"build", "-a", "-tags", tags0Types, "-o", "typed-0-types.bin", "./lists/typed/cmd/"},
							filePath: "typed-0-types.bin",
						},
						{
							name:     "1-types",
							args:     []string{"build", "-a", "-tags", tags1Types, "-o", "typed-1-types.bin", "./lists/typed/cmd/"},
							filePath: "typed-1-types.bin",
						},
						{
							name:     "2-types",
							args:     []string{"build", "-a", "-tags", tags2Types, "-o", "typed-2-types.bin", "./lists/typed/cmd/"},
							filePath: "typed-2-types.bin",
						},
						{
							name:     "3-types",
							args:     []string{"build", "-a", "-tags", tags3Types, "-o", "typed-3-types.bin", "./lists/typed/cmd/"},
							filePath: "typed-3-types.bin",
						},
						{
							name:     "4-types",
							args:     []string{"build", "-a", "-tags", tags4Types, "-o", "typed-4-types.bin", "./lists/typed/cmd/"},
							filePath: "typed-4-types.bin",
						},
						{
							name:     "5-types",
							args:     []string{"build", "-a", "-tags", tags5Types, "-o", "typed-5-types.bin", "./lists/typed/cmd/"},
							filePath: "typed-5-types.bin",
						},
					},
				},
				{
					name:     "pkg",
					fileType: ".a",
					subTests: []gobuildSubTestCase{
						{
							name:     "0-types",
							args:     []string{"build", "-a", "-tags", tags0Types, "-o", "typed-0-types.a", "./lists/typed/"},
							filePath: "typed-0-types.a",
						},
						{
							name:     "1-types",
							args:     []string{"build", "-a", "-tags", tags1Types, "-o", "typed-1-types.a", "./lists/typed/"},
							filePath: "typed-1-types.a",
						},
						{
							name:     "2-types",
							args:     []string{"build", "-a", "-tags", tags2Types, "-o", "typed-2-types.a", "./lists/typed/"},
							filePath: "typed-2-types.a",
						},
						{
							name:     "3-types",
							args:     []string{"build", "-a", "-tags", tags3Types, "-o", "typed-3-types.a", "./lists/typed/"},
							filePath: "typed-3-types.a",
						},
						{
							name:     "4-types",
							args:     []string{"build", "-a", "-tags", tags4Types, "-o", "typed-4-types.a", "./lists/typed/"},
							filePath: "typed-4-types.a",
						},
						{
							name:     "5-types",
							args:     []string{"build", "-a", "-tags", tags5Types, "-o", "typed-5-types.a", "./lists/typed/"},
							filePath: "typed-5-types.a",
						},
					},
				},
			},
		},
	}

	b.ResetTimer()

	for i := range testCases {

		// Capture the test case.
		tc := testCases[i]

		b.Run(tc.name, func(b *testing.B) {

			for j := range tc.testGrps {

				// Capture the test group.
				tg := tc.testGrps[j]

				b.Run(tg.name, func(b *testing.B) {

					for k := range tg.subTests {

						// Capture the sub test.
						st := tg.subTests[k]

						b.Run(st.name, func(b *testing.B) {

							for i := 0; i < b.N; i++ {
								err := exec.Command("go", st.args...).Run()
								if err != nil {
									b.Error(err)
								}

								info, err := os.Stat(st.filePath)
								if err != nil {
									b.Error(err)
								}
								b.ReportMetric(float64(info.Size()), "filesize/op")
							}
						})
					}
				})
			}
		})
	}

	b.StopTimer()
}
