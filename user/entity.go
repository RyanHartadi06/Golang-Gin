package user

import "time"

type User struct {
	ID        int
	Email     string
	Password  string
	Name      string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
