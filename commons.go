package cases

import (
	"regexp"
	"strings"
)

func MapTo(src Case, dest Case, name string) (v string, err error) {
	atoms, parseErr := src.Parse(name)
	if parseErr != nil {
		err = parseErr
		return
	}
	v, err = dest.Format(atoms)
	return
}

func split(s string, sep string) (atoms []string, err error) {
	reg, regErr := regexp.Compile(sep)
	if regErr != nil {
		err = regErr
		return
	}
	matches := reg.FindAllStringIndex(s, -1)
	atoms = make([]string, 0, len(matches))
	for _, match := range matches {
		atoms = append(atoms, strings.ToLower(s[match[0]:match[1]]))
	}
	return
}
