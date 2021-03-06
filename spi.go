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

type Case interface {
	Name() (name string)
	Format(atoms []string) (v string)
	Parse(name string) (atoms []string, err error)
}

func Vars() (v map[string]Case) {
	v = make(map[string]Case)
	v[Camel().Name()] = Camel()
	v[LowerCamel().Name()] = LowerCamel()
	v[Snake().Name()] = Snake()
	v[Kebab().Name()] = Kebab()
	v[Qualified().Name()] = Qualified()
	return
}
