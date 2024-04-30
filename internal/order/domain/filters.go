package domain

type Filters struct {
	Limit     int
	Cursor    int
	Direction PageDirection
}
type PageDirection int

const (
	PageDirectionUnknown PageDirection = iota
	PageDirectionNext
	PageDirectionPrevious
)
