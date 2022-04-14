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

package fuzzaldrin

import (
	"sort"
)

// candidate holds the item along with its score.
type candidate[T any] struct {
	item  T
	score float64
}

func filter[T any](items []T, query string, qHasSlash bool, strFunc func(T) string) (candidates []T) {
	var scored []candidate[T]
	for _, it := range items {
		str := strFunc(it)
		if str == "" {
			continue
		}

		sc := score(str, query)
		if !qHasSlash {
			sc = basenameScore(str, query, sc)
		}
		if sc > 0 {
			scored = append(scored, candidate[T]{it, sc})
		}
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})
	for _, sc := range scored {
		candidates = append(candidates, sc.item)
	}
	return
}
