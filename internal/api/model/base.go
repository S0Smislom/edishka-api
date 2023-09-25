package model

import "time"

type Base struct {
	Id int `json:"id"`
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Creator struct {
	CreatedById int  `json:"created_by_id"`
	UpdatedById *int `json:"updated_by_id"`
}
