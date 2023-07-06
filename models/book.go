package models

import "time"

type Book struct {
	ID        int64
	Title     string
	Author    Author
	CreatedAt time.Time
	UpdateAt  time.Time
}
