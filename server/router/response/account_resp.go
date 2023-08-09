package router

import "github.com/KingFarGrace/CollabSearch/server/entity"

type AccountResponse struct {
	Response
	Users []entity.User
}

func (receiver *AccountResponse) SetReturnObjs(users interface{}) {
	receiver.Users = users.([]entity.User)
}

func (receiver *AccountResponse) SetReturnObj(user interface{}) {
	receiver.Users = make([]entity.User, 0)
	receiver.Users = append(receiver.Users, user.(entity.User))
}
