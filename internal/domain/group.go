package domain

type Group struct {
	ID       int
	Name     string
	Members  []string
	Subgroup Subgroup
}
