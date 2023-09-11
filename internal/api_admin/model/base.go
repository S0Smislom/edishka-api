package model

import "time"

type Base struct {
	Id int `json:"id"`
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
