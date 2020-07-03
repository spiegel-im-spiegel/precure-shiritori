package shiritori

import "testing"

func TestNewWord(t *testing.T) {
	testCases := []struct {
		str string
		w   *Word
	}{
		{str: "キュアイェーー", w: &Word{word: "キュアイェーー", top: 'イ', tail: 'エ'}},
		{str: "キュアイェー。", w: &Word{word: "キュアイェー。", top: 'イ', tail: 'エ'}},
		{str: "キュアイェ々", w: &Word{word: "キュアイェ々", top: 'イ', tail: 'エ'}},
		{str: "㈱", w: &Word{word: "㈱", top: '株', tail: '株'}},
		{str: "㍻", w: &Word{word: "㍻", top: '平', tail: '成'}},
	}

	for _, tc := range testCases {
		w := NewWord(tc.str)
		if w.String() != tc.w.String() {
			t.Errorf("%#v is not %#v", w, tc.w)
		}
		if w.Tail() != tc.w.Tail() {
			t.Errorf("%#v is not %#v", w, tc.w)
		}
		if !w.TopEqual(tc.w.Top()) {
			t.Errorf("%#v is not %#v", w, tc.w)
		}
	}
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
