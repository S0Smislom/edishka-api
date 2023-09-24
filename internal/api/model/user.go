package model

type User struct {
	ID        int     `json:"id"`
	Phone     string  `json:"phone"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Birthday  *string `json:"birthday"`
	Photo     *string `json:"photo"`
	Code      *string `json:"-"`
}

type UpdateUser struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Birthday  *string `json:"birthday"`
}
