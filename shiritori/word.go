package shiritori

import (
	"strings"
	"unicode"

	"github.com/spiegel-im-spiegel/precure-shiritori/kana"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
)

//Word is word info for shiritori
type Word struct {
	word string
	top  rune
	tail rune
}

//NewWord returns new Word instance
func NewWord(s string) *Word {
	if len(s) == 0 {
		return nil
	}
	ss := width.Fold.String(norm.NFKC.String(s))
	ss = strings.TrimPrefix(ss, "キュア") // special rule for precure shiritori
	//rs := []rune(strings.TrimRightFunc(ss, func(r rune) bool { return unicode.Is(unicode.Lm, r) || unicode.Is(unicode.P, r) }))
	rs := []rune(strings.TrimFunc(ss, func(r rune) bool { return unicode.Is(unicode.Lm, r) || unicode.Is(unicode.P, r) }))
	return &Word{word: s, top: kana.Norm(rs[0]), tail: kana.Norm(rs[len(rs)-1])}
}

func NewWordSlice(strs ...string) []*Word {
	if len(strs) == 0 {
		return []*Word{}
	}
	ws := make([]*Word, 0, len(strs))
	for _, s := range strs {
		if w := NewWord(s); w != nil {
			ws = append(ws, w)
		}
	}
	return ws
}

//String is Stringer method
func (w *Word) String() string {
	if w == nil {
		return ""
	}
	return w.word
}

//Equal returns true if word property equals
func (left *Word) Equal(right *Word) bool {
	return strings.EqualFold(left.String(), right.String())
}

//Top returns top character of shiritori
func (w *Word) Top() rune {
	if w == nil {
		return rune(0)
	}
	return w.top
}

//Tail returns tail character of shiritori
func (w *Word) Tail() rune {
	if w == nil {
		return rune(0)
	}
	return w.tail
}

//TopEqual returns true if character in argument equals top character
func (w *Word) TopEqual(c rune) bool {
	if w == nil {
		return false
	}
	return w.Top() == kana.Norm(c)
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
