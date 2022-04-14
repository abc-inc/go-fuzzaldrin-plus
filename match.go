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
	"math"
	"os"
	"unicode"
)

func basenameMatch(str, query string) []int {
	strRunes := []rune(str)

	idx := len(strRunes) - 1
	for idx >= 0 && os.IsPathSeparator(uint8(strRunes[idx])) {
		idx--
	}
	slashCnt := 0
	lastChar := idx
	var base string
	for idx >= 0 {
		if os.IsPathSeparator(uint8(strRunes[idx])) {
			slashCnt++
			if base == "" {
				base = string(strRunes[idx+1 : lastChar+1])
			}
		} else if idx == 0 {
			if base == "" {
				if lastChar < len(strRunes)-1 {
					base = string(strRunes[0 : lastChar+1])
				} else {
					base = string(strRunes)
				}
			}
		}
		idx--
	}

	return match(base, query, len(str)-len(base))
}

func match(str, query string, offset int) (indices []int) {
	strRunes := []rune(str)

	if str == query {
		for idx := offset; idx < len(strRunes)+offset; idx++ {
			indices = append(indices, idx)
		}
		return
	}

	idxInStr := 0

	for _, ch := range query {
		lcIdx := findIndex(strRunes, unicode.ToLower(ch))
		ucIdx := findIndex(strRunes, unicode.ToUpper(ch))

		minIdx := int(math.Min(float64(lcIdx), float64(ucIdx)))
		if minIdx == -1 {
			minIdx = int(math.Max(float64(lcIdx), float64(ucIdx)))
		}
		if minIdx == -1 {
			return
		}
		indices = append(indices, offset+minIdx)
		idxInStr = minIdx

		offset += idxInStr + 1
		strRunes = strRunes[idxInStr+1:]
	}

	return
}

func findIndex(strRunes []rune, ch rune) int {
	for i, c := range strRunes {
		if c == ch {
			return i
		}
	}
	return -1
}
