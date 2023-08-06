package router

import "github.com/KingFarGrace/CollabSearch/server/entity"

type AccountResponse struct {
	Response
	User entity.User
}

func (receiver *AccountResponse) SetReturnObj(user entity.User) {
	receiver.User = user
}
