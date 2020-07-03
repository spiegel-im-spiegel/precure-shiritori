package precure

import (
	"fmt"
	"sort"
	"testing"
)

func TestNames(t *testing.T) {
	names := Names()
	sort.Strings(names)
	nominal := "[キュアアクア キュアアムール キュアアンジュ キュアイーグレット キュアウィンディ キュアエコー キュアエトワール キュアエース キュアエール キュアカスタード キュアグレース キュアコスモ キュアサニー キュアサンシャイン キュアショコラ キュアジェラート キュアスカーレット キュアスター キュアスパークル キュアセレーネ キュアソレイユ キュアソード キュアダイヤモンド キュアトゥインクル キュアドリーム キュアハッピー キュアハニー キュアハート キュアパイン キュアパッション キュアパルフェ キュアビューティ キュアビート キュアピース キュアピーチ キュアフェリーチェ キュアフォンテーヌ キュアフォーチュン キュアフローラ キュアブライト キュアブラック キュアブルーム キュアブロッサム キュアプリンセス キュアベリー キュアホイップ キュアホワイト キュアマカロン キュアマシェリ キュアマジカル キュアマリン キュアマーチ キュアマーメイド キュアミューズ キュアミラクル キュアミルキー キュアミント キュアムーンライト キュアメロディ キュアラブリー キュアリズム キュアルージュ キュアレモネード キュアロゼッタ シャイニールミナス ミルキィローズ]"
	res := fmt.Sprintf("%v", names)
	if res != nominal {
		t.Errorf("Names() is \"%v\", want \"%v\"", res, nominal)
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
