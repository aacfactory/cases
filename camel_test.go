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
	"fmt"
	"testing"
)

func TestCamel(t *testing.T) {

	c := Camel()
	fmt.Println(c.Parse("_Hoo0Bar0"))
	fmt.Println(c.Format([]string{"foo", "bar"}))

}

func TestLowerCamel(t *testing.T) {
	c := LowerCamel()
	fmt.Println(c.Parse("_hoo0Bar0"))
	fmt.Println(c.Format([]string{"foo", "bar"}))

}
