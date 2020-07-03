package kana

import (
	"golang.org/x/text/unicode/norm"
)

var kanaLowers = []rune{'ぁ', 'ぃ', 'ぅ', 'ぇ', 'ぉ', 'ァ', 'ィ', 'ゥ', 'ェ', 'ォ', 'ヵ', 'ヶ', 'っ', 'ッ', 'ゃ', 'ゅ', 'ょ', 'ャ', 'ュ', 'ョ', 'ゎ', 'ヮ'}
var kanaUppers = []rune{'あ', 'い', 'う', 'え', 'お', 'ア', 'イ', 'ウ', 'エ', 'オ', 'カ', 'ケ', 'つ', 'ツ', 'や', 'ゆ', 'よ', 'ヤ', 'ユ', 'ヨ', 'わ', 'ワ'}

//Norm returns normalized hiragana-katakana character
func Norm(c rune) rune {
	return kanaToUpper(baseChar(c))
}

//kanaToUpper returns normal character from hiragana-katakana small character
func kanaToUpper(c rune) rune {
	for i, k := range kanaLowers {
		if c == k {
			return kanaUppers[i]
		}
	}
	return c
}

//baseChar returns base character
func baseChar(c rune) rune {
	s := string([]rune{c})
	for _, cc := range norm.NFD.String(s) {
		return cc
	}
	return c
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
