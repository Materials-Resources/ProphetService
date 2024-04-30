package domain

type Contact struct {
	Id        string
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Branches  []*CustomerBranch
}
