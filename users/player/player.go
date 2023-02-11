package player

import (
	"fmt"
	"theGame/users"
)

type Player struct {
	users.User
}

func NewPlayer(
	addId int,
	addNickName, addEmail, addPassword string,
	addTpUser users.TypeUser,	
) *Player{
	return &Player{
		User: users.User{
			Id: addId,
			NickName: addNickName,
			Email: addEmail,
			Password: addPassword,
			TpUser: addTpUser,
		},
	}

}

func (p Player) PlayGame() {
	fmt.Println("playGame")
}

type playerData map[int]Player

var PLAYER_MAP playerData

func MakePlayerMap() {
	PLAYER_MAP = make(playerData)
}