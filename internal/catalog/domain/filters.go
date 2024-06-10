package domain

type Filter struct {
	Limit  int
	Cursor int
}

type Metadata struct {
	CurrentCursor int
	NextCursor    int
	Limit         int
}
