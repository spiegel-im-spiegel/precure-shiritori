package main

import (
	"fmt"
	"sort"

	"github.com/spiegel-im-spiegel/precure-shiritori/precure"
	"github.com/spiegel-im-spiegel/precure-shiritori/shiritori"
)

//go:generate statik -src=precure/rubicure/config/girls -include=*.yml -p=precure
func main() {
	words := precure.Names()
	sort.Strings(words)
	// fmt.Println("総勢", len(words), "人")
	// fmt.Println(strings.Join(words, ", "))
	lst := shiritori.Do(words...)
	sort.Slice(lst, func(i, j int) bool {
		return lst[i].Len() > lst[j].Len()
	})
	for _, l := range lst {
		fmt.Println(l.Len(), ":", l)
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
