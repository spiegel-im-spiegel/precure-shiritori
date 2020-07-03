package shiritori

func Do(strs ...string) []*List {
	ws := NewWords(NewWordSlice(strs...))
	if ws.Len() == 0 {
		return []*List{}
	}

	slst := []*List{}
	w := ws.Begin()
	for {
		if w == nil {
			break
		}
		l := do(w, ws.Remove(w))
		slst = append(slst, l...)
		w = ws.Next()
	}
	return slst
}

func do(first *Word, rest *Words) []*List {
	if rest.Len() == 0 {
		return []*List{NewList(first.String())}
	}
	tail := first.Tail()
	if tail == 'ん' || tail == 'ン' {
		return []*List{NewList(first.String())}
	}
	wlst := rest.Filter(tail)
	if wlst.Len() == 0 {
		return []*List{NewList(first.String())}
	}

	slst := []*List{}
	next := wlst.Begin()
	for {
		if next == nil {
			break
		}
		ll := do(next, rest.Remove(next))
		for _, l := range ll {
			slst = append(slst, NewList(first.String()).AddList(l))
		}
		next = wlst.Next()
	}
	return slst
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
