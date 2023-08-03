package router

import "github.com/KingFarGrace/CollabSearch/server/entity"

type AccountResponse struct {
	Response
	user entity.User
}

func (receiver *AccountResponse) SetReturnObj(user entity.User) {
	receiver.user = user
}
