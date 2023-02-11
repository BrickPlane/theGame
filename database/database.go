package database

import (
	"fmt"
	"theGame/users"
	"theGame/users/admin"
	"theGame/users/player"
	"theGame/users/support"
)

type StatAdmin interface {
	CreateGame()
}

func CheckAdminStat(s StatAdmin) {
	s.CreateGame()
}

type StatUser interface {
	PlayGame()
}

func CheckUserStat(p StatUser) {
	p.PlayGame()
}



func initPlayer() {
	admin.MakeAdminMap()
	player.MakePlayerMap()
	support.MakeSupportData()

	adminInst := admin.NewAdmin(1, "cucumber", "admin@mail", "d324ffs", users.ADMIN)
	CheckAdminStat(adminInst)
	CheckUserStat(adminInst)
	playerInst := player.NewPlayer(2, "tomatoes", "player@mail", "l9ukhrt", users.PLAYER)
	CheckUserStat(playerInst)
	supportInst := support.NewSupport(3, "apple", "support@mail", "aweesrfv4", users.SUPPORT)
	CheckAdminStat(supportInst)
	CheckUserStat(supportInst)

	CheckMap(adminInst)
	CheckMap(playerInst)
	CheckMap(supportInst)
}
func CheckMap(mapType any) {

	covertObj, ok := mapType.(users.User)
	if !ok {
		fmt.Println("err")
	}
	switch covertObj.TpUser {
	case users.ADMIN:
		adminInst, _ := mapType.(admin.Admin)
		admin.ADMIN_MAP[adminInst.Id] = adminInst
	case users.PLAYER:
		playerInst, _ := mapType.(player.Player)
		player.PLAYER_MAP[playerInst.Id] = playerInst
	case users.SUPPORT:
		supportInst, _ := mapType.(support.Support)
		support.SUPPORT_MAP[supportInst.Id] = supportInst
		}
	}