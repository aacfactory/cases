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
