package precure

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/rakyll/statik/fs"
)

//Names returns list of Precure names
func Names() []string {
	nmMap := map[string]string{}
	// Get filesystem
	statikFS, err := fs.New()
	if err != nil {
		return []string{}
	}
	//Get precure names in files
	for _, path := range getPaths(statikFS) {
		for _, nm := range getNames(statikFS, path) {
			if len(nm) > 0 && !strings.HasSuffix(nm, "）") {
				nmMap[nm] = nm
			}
		}
	}
	//Output name list
	names := []string{}
	for k, _ := range nmMap {
		names = append(names, k)
	}
	return names
}

func getPaths(statikFS http.FileSystem) []string {
	paths := []string{}
	err := fs.Walk(statikFS, "/", func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		paths = append(paths, path)
		return nil
	})
	if err != nil {
		return []string{}
	}
	return paths
}

func getNames(statikFS http.FileSystem, path string) []string {
	names := []string{}
	r, err := statikFS.Open(path)
	if err != nil {
		return names
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(s, "precure_name:") {
			names = append(names, strings.TrimSpace(strings.TrimPrefix(s, "precure_name:")))
		}
	}
	if err := scanner.Err(); err != nil {
		return []string{}
	}
	return names
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
