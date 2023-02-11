package support

import (
	"fmt"
	"theGame/users"
)

type Support struct {
	users.User
	users.Developer
}

func NewSupport(
	addId int,
	addNickName, addEmail, addPasswors string,
	addTpUser users.TypeUser,
) *Support{
	return &Support{
		User: users.User{
			Id: addId,
			NickName: addNickName,
			Email: addEmail,
			Password: addPasswors,
			TpUser: addTpUser,
		},
	}
}

func (p Support) PlayGame() {
	fmt.Println("playGame")
}

func (s Support) CreateGame() {
	fmt.Println("createGame")
}

type supportData map[int]Support

var SUPPORT_MAP supportData

func MakeSupportData() {
	SUPPORT_MAP = make(supportData)
}

