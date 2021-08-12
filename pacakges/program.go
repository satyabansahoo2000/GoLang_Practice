package main

type Program struct {
	Name string
	Year string
}

func (p Program) details() string {
	return p.Name + " was founded in " + p.Year
}