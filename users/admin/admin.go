package admin

import (
	"fmt"
	"theGame/users"
)

type Admin struct {
	users.User
	users.Developer
}

func NewAdmin(
	addId int,
	addNickName, addEmail, addPassword string,
	addTpUser users.TypeUser,
) *Admin{
	return &Admin{
		User: users.User{
			Id: addId,
			NickName: addNickName,
			Email: addEmail,
			Password: addPassword,
			TpUser: addTpUser,
		},
	}
}

func (p Admin) PlayGame() {
	fmt.Println("playGame")
}

func (s Admin) CreateGame() {
	fmt.Println("createGame")
}

type adminData map[int]Admin

var ADMIN_MAP adminData

func MakeAdminMap() {
	ADMIN_MAP = make(adminData)
}

