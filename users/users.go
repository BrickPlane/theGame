package users

type TypeUser byte

const (
	ADMIN TypeUser = iota + 1
	PLAYER
	SUPPORT
)

type User struct {
	Id int
	NickName string
	Email string
	Password string
	TpUser TypeUser 
}

type Developer struct {
	Ban	int
	Unban int
	Delete int
}

