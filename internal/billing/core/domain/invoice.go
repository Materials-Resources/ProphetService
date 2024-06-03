package domain

import "time"

type Invoice struct {
	Id        string
	OrderID   string
	Total     float64
	CreatedAt time.Time
}
