package models

import "time"

type Receipt struct {
	Id        int
	Total     float64
	CreatedAt time.Time
}
