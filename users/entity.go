package users

import "time"

type User struct {
	ID         int
	Name       string
	Occupation string
	Email      string
	Password   string
	Role       string
	Avatar     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
