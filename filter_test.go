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
	"path"
	"testing"

	. "github.com/abc-inc/go-fuzzaldrin-plus"
	. "github.com/stretchr/testify/assert"
)

var sep = string(PathSeparator)

func buildPath(seg ...string) string {
	p := string(PathSeparator)
	for _, s := range seg {
		p = path.Join(p, s)
	}
	return p
}

func testBestMatch(items []string, query, exp string, t *testing.T) {
	results := Filter(items, query, ident[string])[0]
	Equal(t, exp, results)
}

func TestFilter_General(t *testing.T) {
	its := []string{"Gruntfile", "filter", "bile", ""}
	res := Filter(its, "file", ident[string])
	Equal(t, []string{"filter", "Gruntfile"}, res)

	res = Filter(its, "file", ident[string])[:1]
	Equal(t, []string{"filter"}, res)
}

func TestFilter_Path(t *testing.T) {
	its := []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
	}
	testBestMatch(its, "bar", its[1], t)

	its = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar", sep, sep, sep, sep, sep),
	}
	testBestMatch(its, "bar", its[1], t)

	its = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
		"bar",
	}
	testBestMatch(its, "bar", its[2], t)

	its = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
		buildPath("bar"),
	}
	testBestMatch(its, "bar", its[2], t)

	its = []string{
		buildPath("bar", "foo"),
		fmt.Sprintf("bar%s%s%s%s%s", sep, sep, sep, sep, sep),
	}
	testBestMatch(its, "bar", its[1], t)

	its = []string{
		path.Join("f", "o", "1_a_z"),
		path.Join("f", "o", "a_z"),
	}
	testBestMatch(its, "az", its[1], t)

	its = []string{
		path.Join("f", "1_a_z"),
		path.Join("f", "o", "a_z"),
	}
	testBestMatch(its, "az", its[1], t)
}

func TestFilter_OnlySep(t *testing.T) {
	Empty(t, Filter([]string{sep}, "bar", ident[string]))
}

func TestFilter_WithSpace(t *testing.T) {
	its := []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar"),
	}
	testBestMatch(its, "br f", its[0], t)

	its = []string{
		buildPath("bar", "foo"),
		buildPath("foo", "bar foo"),
	}
	testBestMatch(its, "br f", its[1], t)

	its = []string{
		buildPath("barfoo", "foo"),
		buildPath("foo", "barfoo"),
	}
	testBestMatch(its, "br f", its[1], t)

	its = []string{
		buildPath("lib", "exportable.rb"),
		buildPath("app", "models", "table.rb"),
	}
	testBestMatch(its, "table", its[1], t)
}

func TestFilter_MixedCase(t *testing.T) {
	its := []string{"statusurl", "StatusUrl"}

	testBestMatch(its, "Status", "StatusUrl", t)
	testBestMatch(its, "SU", "StatusUrl", t)
	testBestMatch(its, "status", "statusurl", t)
	testBestMatch(its, "su", "statusurl", t)
	testBestMatch(its, "statusurl", "statusurl", t)
	testBestMatch(its, "StatusUrl", "StatusUrl", t)
}

func TestFilter_WithSign(t *testing.T) {
	its := []string{"sub-zero", "sub zero", "sub_zero"}
	testBestMatch(its, "sz", its[0], t)

	its = []string{"sub zero", "sub_zero", "sub-zero"}
	testBestMatch(its, "sz", its[0], t)

	its = []string{"sub_zero", "sub-zero", "sub zero"}
	testBestMatch(its, "sz", its[0], t)

	its = []string{"a_b_c", "a_b"}
	testBestMatch(its, "ab", its[1], t)

	its = []string{"z_a_b", "a_b"}
	testBestMatch(its, "ab", its[1], t)

	its = []string{"a_b_c", "c_a_b"}
	testBestMatch(its, "ab", its[0], t)
}

func TestFilter_DirectoryDepth(t *testing.T) {
	its := []string{
		buildPath("app", "models", "sutomotive", "car.rb"),
		buildPath("spec", "factories", "cars.rb"),
	}
	testBestMatch(its, "car.rb", its[0], t)

	its = []string{
		buildPath("app", "models", "sutomotive", "car.rb"),
		"car.rb",
	}
	testBestMatch(its, "car.rb", its[1], t)

	its = []string{
		"car.rb",
		buildPath("app", "models", "sutomotive", "car.rb"),
	}
	testBestMatch(its, "car.rb", its[0], t)

	its = []string{
		buildPath("app", "models", "cars", "car.rb"),
		buildPath("spec", "cars.rb"),
	}
	testBestMatch(its, "car.rb", its[0], t)
}

// ident always returns its input argument.
func ident[T any](t T) T {
	return t
}
