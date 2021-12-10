package cases

import "strings"

func MapTo(src Case, dest Case, name string) (v string, err error) {
	atoms, parseErr := src.Parse(name)
	if parseErr != nil {
		err = parseErr
		return
	}
	v, err = dest.Format(atoms)
	return
}

func Split(s string, sep string) (v []string) {
	atoms := strings.Split(s, sep)
	if len(atoms) == 1 && atoms[0] == "" {
		v = make([]string, 0, 1)
		return
	}
	v = atoms
	return
}
