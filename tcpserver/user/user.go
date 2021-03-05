package user

type User struct {
	Id     int32
	Name   string
	Passwd string
	Addr   string

}

var usersmap  map[int]*User


func GetUsers() map[int]*User {
	return usersmap
}

