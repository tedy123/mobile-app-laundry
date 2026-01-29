package model

import "time"

type Role string

const (
	Owner   Role = "Owner"
	Admin   Role = "Admin"
	Cashier Role = "Cashier"
)

type User struct {
	ID        string
	Email     string
	Password  *string
	GoogleUID *string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
