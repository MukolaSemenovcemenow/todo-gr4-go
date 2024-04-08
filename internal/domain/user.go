package domain

import (
	"time"
)

type User struct {
	Id          uint64
	Phone       string
	FirstName   string
	SecondName  string
	Email       string
	Role        Role
	CreatedDate time.Time
	UpdatedDate time.Time
	DeletedDate *time.Time
}

type Role string

const (
	AdminRole    Role = "ADMIN"
	CustomerRole Role = "CUSTOMER"
)

type ChangePassword struct {
	OldPassword string
	NewPassword string
}

func (u User) GetUserId() uint64 {
	return u.Id
}
