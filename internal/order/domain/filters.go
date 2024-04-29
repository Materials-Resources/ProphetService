package domain

type Filters struct {
	Limit     int32
	Cursor    int32
	Direction PageDirection
}
type PageDirection int32

const (
	Next PageDirection = iota
	Previous
)
