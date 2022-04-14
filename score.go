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
	"strings"
	"unicode"
)

func basenameScore(str, query string, sc float64) float64 {
	idx := len(str) - 1
	for idx >= 0 && os.IsPathSeparator(str[idx]) {
		idx--
	}

	var base string
	slashCnt := 0
	lastChar := idx
	for idx >= 0 {
		if os.IsPathSeparator(str[idx]) {
			slashCnt++
			if base == "" {
				base = str[idx+1 : lastChar+1]
			}
		} else if idx == 0 {
			if lastChar < len(str)-1 {
				if base == "" {
					base = str[0 : lastChar+1]
				}
			} else {
				if base == "" {
					base = str
				}
			}
		}
		idx--
	}

	if base == str {
		sc *= 2
	} else if base != "" {
		sc += score(base, query)
	}
	segCnt := slashCnt + 1
	depth := max(1, 10-segCnt)
	sc *= float64(depth) * 0.01
	return sc
}

func score(str, query string) float64 {
	if str == query {
		return 1
	}
	if queryIsLastPathSegment(str, query) {
		return 1
	}

	totSc := 0.0
	strLen := len(str)
	idxInStr := 0
	for idxInQuery := 0; idxInQuery < len(query); idxInQuery++ {
		c := rune(query[idxInQuery])
		lcIdx := strings.IndexRune(str, unicode.ToLower(c))
		ucIdx := strings.IndexRune(str, unicode.ToUpper(c))
		minIdx := min(lcIdx, ucIdx)
		if minIdx == -1 {
			minIdx = max(lcIdx, ucIdx)
		}
		idxInStr = minIdx
		if idxInStr == -1 {
			return 0
		}

		cSc := 0.1
		if str[idxInStr] == byte(c) {
			cSc += 0.1
		}
		if idxInStr == 0 || os.IsPathSeparator(str[idxInStr-1]) {
			cSc += 0.8
		} else if c == '-' || c == '_' || c == ' ' {
			cSc += 0.7
		}
		str = str[idxInStr+1:]
		totSc += cSc
	}
	querySc := totSc / float64(len(query))
	return ((querySc * (float64(len(query)) / float64(strLen))) + querySc) / 2.0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
