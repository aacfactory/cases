package cases

import "strings"

var _snake = &snake{}

func Snake() Case {
	return _snake
}

type snake struct{}

func (s *snake) Name() (name string) {
	name = "SNAKE"
	return
}

func (s *snake) Format(atoms []string) (v string, err error) {
	for _, atom := range atoms {
		v = v + "_" + strings.TrimSpace(atom)
	}
	if v == "" {
		return
	}
	v = strings.ToLower(v[1:])
	return
}

func (s *snake) Parse(name string) (atoms []string, err error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return
	}
	v := strings.Split(name, "_")
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
