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
	"bytes"
	"fmt"
)

var _camel = &camel{}

func Camel() Case {
	return _camel
}

type camel struct {
}

func (c *camel) Name() (name string) {
	name = "CAMEL"
	return
}

func (c *camel) Format(atoms []string) (v string) {
	buf := bytes.NewBufferString("")
	for _, atom := range atoms {
		if len(atom) > 0 {
			ch := atom[0]
			if 'a' <= ch && ch <= 'z' {
				ch = bytes.ToUpper([]byte(string(ch)))[0]
				buf.WriteByte(ch)
				buf.WriteString(atom[1:])
			}
		} else {
			buf.WriteString(atom)
		}
	}
	v = buf.String()
	return
}

func (c *camel) Parse(name string) (atoms []string, err error) {
	atoms, err = split(name, "([A-Z]|[a-z])([a-z0-9]+)")
	if err != nil {
		err = fmt.Errorf("parse camel failed, %v", err)
		return
	}
	return
}
