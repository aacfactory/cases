package cases

import (
	"fmt"
	"testing"
)

func TestSnake(t *testing.T) {

	c := Snake()
	fmt.Println(c.Format([]string{"foo", "bar"}))
	fmt.Println(c.Parse("_foo_bar"))
}
