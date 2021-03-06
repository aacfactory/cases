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

import "strings"

var _qualified = &qualified{}

func Qualified() Case {
	return _qualified
}

type qualified struct {
}

func (q *qualified) Name() (name string) {
	name = "QUALIFIED"
	return
}

func (q *qualified) Format(atoms []string) (v string) {
	for _, atom := range atoms {
		v = v + "." + strings.TrimSpace(atom)
	}
	if v == "" {
		return
	}
	v = strings.ToLower(v[1:])
	return
}

func (q *qualified) Parse(name string) (atoms []string, err error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	}
	v := strings.Split(name, ".")
	if v == nil || len(v) == 0 {
		return
	}
	atoms = make([]string, 0, 1)
	for _, e := range v {
		e = strings.TrimSpace(e)
		if e == "" {
			continue
		}
		atoms = append(atoms, e)
	}
	return
}
