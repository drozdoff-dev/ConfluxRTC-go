package domain

import "time"

type Room struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
