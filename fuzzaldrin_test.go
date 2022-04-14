// Copyright 2022 The go-fuzzaldrin-plus authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fuzzaldrin_test

import (
	"fmt"
	"strings"

	"github.com/abc-inc/go-fuzzaldrin-plus"
)

type entry struct {
	name string
	id   int
}

func ExampleFilter_string() {
	ss := []string{"Call", "Me", "Maybe"}
	res := fuzzaldrin.Filter(ss, "me", ident[string])
	fmt.Println(strings.Join(res, "\n"))
	// Output:
	// Me
	// Maybe
}

func ExampleFilter_struct() {
	es := []entry{
		{name: "Call", id: 1},
		{name: "Me", id: 2},
		{name: "Maybe", id: 3},
	}
	res := fuzzaldrin.Filter[entry](es, "me", func(t entry) string { return t.name })
	for _, i := range res {
		fmt.Println(i)
	}
	// Output:
	// {Me 2}
	// {Maybe 3}
}

func ExampleMatch() {
	fmt.Println(fuzzaldrin.Match("Call Me Maybe", "m m"))
	fmt.Println(fuzzaldrin.Match("Call Me Maybe", "eY"))
	// Output:
	// [5 8]
	// [6 10]
}

func ExampleScore() {
	fmt.Println(fuzzaldrin.Score("Me", "me"))
	fmt.Println(fuzzaldrin.Score("Maybe", "me"))
	// Output:
	// 0.17099999999999999
	// 0.0693
}
