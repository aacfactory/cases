package cases

import (
	"fmt"
	"testing"
)

func TestCamel(t *testing.T) {

	c := Camel()
	fmt.Println(c.Parse("fooBar"))
	fmt.Println(c.Format([]string{"foo", "bar"}))

}
