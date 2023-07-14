package models

import "time"

type Author struct {
	ID        int64     `json:id`
	Name      string    `json:name`
	CreatedAt time.Time `json:createdAt`
	UpdatedAt time.Time `json:updatedAt`
}
