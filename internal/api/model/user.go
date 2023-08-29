package model

type User struct {
	ID        int     `json:"id"`
	Phone     string  `json:"phone"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Birthday  *string `json:"birthday"`
	Code      *string `json:"-"`
}
