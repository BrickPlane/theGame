package support

import (
	"errors"
	"fmt"
	"theGame/users"
)

type Support struct {
	users.User
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
func (p *Support) setID(id int) error {
	if id <= 0 {
		return errors.New("Cant be less 1")
	}

	p.Id = id

	return nil
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

