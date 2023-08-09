package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

type AccountResponse struct {
	Response
	Users []entity.User
}

func (receiver *AccountResponse) SetReturnObjs(objs interface{}) {
	if users, ok := objs.([]entity.User); ok {
		receiver.Users = users
	} else {
		util.WarnLogger("SetReturnObjs()", "Failed to set return objs.")
	}
}
