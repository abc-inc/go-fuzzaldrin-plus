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
	"os"
	"sort"
	"strings"
)

var PathSeparator = os.PathSeparator

// Filter sorts and filters the given items by matching them against the query.
// It returns a slice of items sorted by best match against the query.
func Filter[T any](items []T, query string, strFunc func(T) string) (res []T) {
	qHasSlash := strings.ContainsRune(query, PathSeparator)
	query = strings.ReplaceAll(query, " ", "")
	return filter(items, query, qHasSlash, strFunc)
}

// Score calculates the score of the given string against the query.
func Score(str, query string) float64 {
	if str == "" || query == "" {
		return 0
	} else if str == query {
		return 2
	}

	qHasSlash := strings.ContainsRune(query, os.PathSeparator)
	query = strings.ReplaceAll(query, " ", "")
	sc := score(str, query)
	if !qHasSlash {
		sc = basenameScore(str, query, sc)
	}
	return sc
}

// Match returns the indices of any query matches in the given string.
func Match(str, query string) (indices []int) {
	if str == "" || query == "" {
		return
	}

	strRunes := []rune(str)
	if str == query {
		for i := 0; i < len(strRunes); i++ {
			indices = append(indices, i)
		}
		return
	}

	qHasSlash := strings.ContainsRune(query, os.PathSeparator)
	query = strings.ReplaceAll(query, " ", "")
	indices = match(str, query, 0)
	if !qHasSlash {
		baseIdx := basenameMatch(str, query)
		for _, idx := range baseIdx {
			if contains(indices, idx) {
				continue
			}
			indices = append(indices, idx)
		}
		sort.Ints(indices)
	}
	return
}

func queryIsLastPathSegment(str, query string) bool {
	if len(str) > len(query) && os.IsPathSeparator(str[len(str)-len(query)-1]) {
		return strings.LastIndex(str, query) == len(str)-len(query)
	}
	return false
}

func contains(is []int, search int) bool {
	for _, i := range is {
		if i == search {
			return true
		}
	}
	return false
}
