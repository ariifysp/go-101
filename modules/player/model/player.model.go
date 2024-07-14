package model

import "time"

type Player struct {
	ID        string
	Email     string
	Name      string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
