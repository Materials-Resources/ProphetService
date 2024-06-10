package domain

type CursorDirection int32

const (
	CursorDirectionUnspecified CursorDirection = iota
	CursorDirectionPrevious
	CursorDirectionNext
)

type Pagination struct {
	Direction CursorDirection
	Cursor    int32
	Size      int64
}

type PaginationMetadata struct {
	NextCursor     int32
	PreviousCursor int32
	Total          int64
}
