package cases

import (
	"bytes"
	"fmt"
)

func Camel() Case {
	return &camel{}
}

type camel struct {
}

func (c *camel) Name() (name string) {
	name = "CAMEL"
	return
}

func (c *camel) Format(atoms []string) (v string, err error) {
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
