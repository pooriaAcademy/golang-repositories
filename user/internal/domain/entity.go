package domain

import (
	"github.com/thanhpk/randstr"
)

type User struct {
	Name string
	Id string
}

func NewUser(Name string) User {
	return User{Name: Name, Id: randstr.Hex(16)}
}


