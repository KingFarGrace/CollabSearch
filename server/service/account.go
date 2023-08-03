package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	response "github.com/KingFarGrace/CollabSearch/server/router/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

// TODO: Give an accurate path.
var defaultAvatarPath = ""

// Register a new user.
// Receive a RegisterJSON then convert it to a User instance,
// then register the User through mapper.
func Register(registerJSON entity.RegisterJSON) *response.AccountResponse {
	resp := new(response.AccountResponse)
	if mapper.ExistEmail(registerJSON.Email) {
		msg := util.Concat("Duplicated user email: ", registerJSON.Email)
		resp.New(response.AccountGroupCode, 2, msg)
		return resp
	}
	uid := util.GetSnowflakeID(1).Int64()
	user := entity.User{
		Uid:      uid,
		Email:    registerJSON.Email,
		Username: registerJSON.Username,
		Password: registerJSON.Password,
		Profile:  "Introduce yourself to others.",
		Avatar:   defaultAvatarPath,
	}
	if success := mapper.InsertUser(user); success {
		resp.New(response.AccountGroupCode, 0, "Success.")
		return resp
	} else {
		resp.New(response.AccountGroupCode, 3, "Success")
		return resp
	}
}
