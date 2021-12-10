package cases

type Case interface {
	Name() (name string)
	Format(atoms []string) (v string, err error)
	Parse(name string) (atoms []string, err error)
}

func Vars() (v map[string]Case) {
	v = make(map[string]Case)
	// todo

	return
}
