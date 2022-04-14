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
	"os"
	"path"
	"testing"

	. "github.com/abc-inc/go-fuzzaldrin-plus"
	. "github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	Equal(t, []int{0, 1}, Match("Hello World", "he"))
	Nil(t, Match("", ""))
	Equal(t, []int{6, 7, 8}, Match("Hello World", "wor"))
	Equal(t, []int{10}, Match("Hello World", "d"))
	Equal(t, []int{1, 2, 6, 7, 8}, Match("Hello World", "elwor"))
	Equal(t, []int{1, 8}, Match("Hello World", "er"))
	Nil(t, Match("Hello World", ""))
	Nil(t, Match("", "abc"))
}

func TestMatch_Path(t *testing.T) {
	Equal(t, []int{0, 1, 2}, Match(path.Join("X", "Y"), path.Join("X", "Y")))
	Equal(t, []int{0, 2}, Match(path.Join("X", "X-x"), "X"))
	Equal(t, []int{0, 2}, Match(path.Join("X", "Y"), "XY"))
	Equal(t, []int{2}, Match(path.Join("-", "X"), "X"))
	Equal(t, []int{0, 2}, Match(path.Join("X-", "-"), fmt.Sprintf("X%s", string(os.PathSeparator))))
}

func TestMatch_DoubleMatch(t *testing.T) {
	Equal(t, []int{0, 1, 3, 4}, Match(path.Join("XY", "XY"), "XY"))
	Equal(t, []int{2, 4, 8, 11}, Match(path.Join("--X-Y-", "-X--Y"), "XY"))
}

func TestMatch_Something(t *testing.T) {
	Equal(t, []int{0}, Match("/", "/"))
}
