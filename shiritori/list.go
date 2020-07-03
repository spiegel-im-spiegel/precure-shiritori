package shiritori

import "strings"

//Word is word info for shiritori
type List struct {
	list []string
}

func NewList(strs ...string) *List {
	if len(strs) == 0 {
		return &List{list: []string{}}
	}
	lst := make([]string, len(strs), cap(strs))
	copy(lst, strs)
	return &List{list: lst}
}

func (l *List) String() string {
	if l.Len() == 0 {
		return ""
	}
	return strings.Join(l.list, ", ")
}

func (l *List) AddStrings(strs ...string) *List {
	if l == nil {
		return NewList(strs...)
	}
	l.list = append(l.list, strs...)
	return l
}

func (l *List) AddList(add *List) *List {
	if add.Len() == 0 {
		return l
	}
	return l.AddStrings(add.list...)
}

func (l *List) Len() int {
	if l == nil {
		return 0
	}
	return len(l.list)
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
