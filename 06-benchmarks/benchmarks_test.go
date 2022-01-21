//go:build benchmarks
// +build benchmarks

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
	"strings"
	"testing"

	blist "go-generics-the-hard-way/06-benchmarks/lists/boxed-list"
	glist "go-generics-the-hard-way/06-benchmarks/lists/generic-list"
	tlist "go-generics-the-hard-way/06-benchmarks/lists/typed-list"
)

func BenchmarkBoxing(b *testing.B) {
	b.Run("BoxedList", func(b *testing.B) {
		var list blist.List
		for i := 0; i < b.N; i++ {
			list = append(list, i)
		}
	})
	b.Run("GenericList", func(b *testing.B) {
		var list glist.List[int]
		for i := 0; i < b.N; i++ {
			list = append(list, i)
		}
	})
	b.Run("TypedList", func(b *testing.B) {
		var list tlist.IntList
		for i := 0; i < b.N; i++ {
			list = append(list, i)
		}
	})
}

func BenchmarkBuildTimes(b *testing.B) {

	args := func(listType string, types ...string) []string {
		if len(types) == 0 {
			types = []string{
				"int", "int8", "int16", "int32", "int64",
			}
		}
		return []string{
			"build",
			"-a",
			"-tags",
			strings.Join(types, ","),
			"./lists/" + listType + "-list",
		}
	}

	run := func(b *testing.B, listType string, types ...string) {
		for i := 0; i < b.N; i++ {
			err := exec.Command("go", args(listType, types...)...).Run()
			if err != nil {
				b.Error(err)
			}
		}
	}

	runTypedOrGeneric := func(b *testing.B, listType string) {
		b.Run("1-type", func(b *testing.B) {
			run(b, listType, "int")
		})
		b.Run("2-types", func(b *testing.B) {
			run(b, listType, "int", "int8")
		})
		b.Run("3-types", func(b *testing.B) {
			run(b, listType, "int", "int8", "int16")
		})
		b.Run("4-types", func(b *testing.B) {
			run(b, listType, "int", "int8", "int16", "int32")
		})
		b.Run("5-types", func(b *testing.B) {
			run(b, listType, "int", "int8", "int16", "int32", "int64")
		})
	}

	b.Run("BoxedList", func(b *testing.B) {
		run(b, "boxed")
	})

	b.Run("GenericList", func(b *testing.B) {
		runTypedOrGeneric(b, "generic")
	})
	b.Run("TypedList", func(b *testing.B) {
		runTypedOrGeneric(b, "typed")
	})
}

func TestFileSizes(t *testing.T) {

	args := func(listType string, types ...string) []string {
		if len(types) == 0 {
			types = []string{
				"int", "int8", "int16", "int32", "int64",
			}
		}
		return []string{
			"build",
			"-a",
			"-o",
			listType + ".a",
			"-tags",
			strings.Join(types, ","),
			"./lists/" + listType + "-list",
		}
	}

	build := func(t *testing.T, listType string, types ...string) {
		err := exec.Command("go", args(listType, types...)...).Run()
		if err != nil {
			t.Fatal(err)
		}
		info, err := os.Stat(listType + ".a")
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("file size: %d", info.Size())
	}

	buildTypedOrGeneric := func(t *testing.T, listType string) {
		t.Run("1-type", func(t *testing.T) {
			build(t, listType, "int")
		})
		t.Run("2-types", func(t *testing.T) {
			build(t, listType, "int", "int8")
		})
		t.Run("3-types", func(t *testing.T) {
			build(t, listType, "int", "int8", "int16")
		})
		t.Run("4-types", func(t *testing.T) {
			build(t, listType, "int", "int8", "int16", "int32")
		})
		t.Run("5-types", func(t *testing.T) {
			build(t, listType, "int", "int8", "int16", "int32", "int64")
		})
	}

	t.Run("Boxed", func(t *testing.T) {
		build(t, "boxed")
	})
	t.Run("Generic", func(t *testing.T) {
		buildTypedOrGeneric(t, "generic")
	})
	t.Run("Typed", func(t *testing.T) {
		buildTypedOrGeneric(t, "typed")
	})
}
