package shiritori

type Words struct {
	words  []*Word
	offset int
}

//NewWords returns new Words instance
func NewWords(words []*Word) *Words {
	return &Words{words: words, offset: -1}
}

func (ws *Words) Begin() *Word {
	if ws.Len() == 0 {
		return nil
	}
	ws.offset = 0
	return ws.words[ws.offset]
}

func (ws *Words) Next() *Word {
	if ws.Len() == 0 {
		return nil
	}
	if ws.offset < 0 || !(ws.offset < len(ws.words)-1) {
		return nil
	}
	ws.offset++
	return ws.words[ws.offset]
}

func (ws *Words) Filter(c rune) *Words {
	if ws.Len() == 0 {
		return nil
	}
	wss := make([]*Word, 0)
	for _, w := range ws.words {
		if w.TopEqual(c) {
			wss = append(wss, w)
		}
	}
	return NewWords(wss)
}

func (ws *Words) Remove(word *Word) *Words {
	if !ws.In(word) {
		return ws.Clone()
	}
	wss := make([]*Word, 0, cap(ws.words)-1)
	for _, w := range ws.words {
		if !w.Equal(word) {
			wss = append(wss, w)
		}
	}
	return NewWords(wss)
}

//Clone returns copy instance of mine
func (ws *Words) Clone() *Words {
	if ws.Len() == 0 {
		return nil
	}
	wss := make([]*Word, len(ws.words), cap(ws.words))
	copy(wss, ws.words)
	return NewWords(wss)
}

func (ws *Words) In(word *Word) bool {
	if ws.Len() == 0 {
		return false
	}
	for _, w := range ws.words {
		if w.Equal(word) {
			return true
		}
	}
	return false
}

func (ws *Words) Len() int {
	if ws == nil {
		return 0
	}
	return len(ws.words)
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
