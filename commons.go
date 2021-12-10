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
	for i := 0; i < len(matches); i++ {
		beg := 0
		if i > 0 {
			beg = matches[i][0]
		}
		end := len(s)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		atoms = append(atoms, strings.ToLower(s[beg:end]))
	}
	return
}
