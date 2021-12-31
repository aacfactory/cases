/*
 * Copyright 2021 Wang Min Xiang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package cases

import (
	"regexp"
	"strings"
)

func MapTo(src Case, dest Case, name string) (v string, err error) {
	atoms, parseErr := src.Parse(name)
	if parseErr != nil {
		err = parseErr
		return
	}
	v = dest.Format(atoms)
	return
}

func split(s string, sep string) (atoms []string, err error) {
	reg, regErr := regexp.Compile(sep)
	if regErr != nil {
		err = regErr
		return
	}
	matches := reg.FindAllStringIndex(s, -1)
	atoms = make([]string, 0, len(matches))
	for i := 0; i < len(matches); i++ {
		beg := 0
		if i > 0 {
			beg = matches[i][0]
		}
		end := len(s)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		atoms = append(atoms, strings.ToLower(s[beg:end]))
	}
	return
}
