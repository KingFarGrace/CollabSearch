package service

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	response "github.com/KingFarGrace/CollabSearch/server/router/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"strings"
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
		Email:    strings.ToLower(registerJSON.Email), // Resolve case sensitive issues.
		Username: registerJSON.Username,
		Password: registerJSON.Password,
		Profile:  "Introduce yourself to others.",
		Avatar:   defaultAvatarPath,
	}
	if success := mapper.InsertUser(user); success {
		resp.New(response.AccountGroupCode, 0, "Success.")
		return resp
	} else {
		resp.New(response.AccountGroupCode, 3, "Failed to sign up, please try later.")
		return resp
	}
}

// LoginWithoutToken is a function for no-token user to login.
// no-token user: new user or whose token was expired.
// When someone sign in the system through this function,
// they should cache the user data in localstorage.
func LoginWithoutToken(loginJSON entity.LoginJSON) (*response.AccountResponse, string) {
	resp := new(response.AccountResponse)
	var user = mapper.GetUserByEmail(strings.ToLower(loginJSON.Email))
	if user == nil {
		msg := util.Concat("No such user. Email: ", loginJSON.Email)
		resp.New(response.AccountGroupCode, 4, msg)
		return resp, ""
	}
	if user.Password != loginJSON.Password {
		resp.New(response.AccountGroupCode, 5, "Wrong password.")
		return resp, ""
	}
	token := util.GenerateJWT(strings.ToLower(loginJSON.Email), conf.Salt)
	resp.New(response.AccountGroupCode, 0, "Success.")
	user.Password = "" // Block user privacy data.
	resp.SetReturnObj(*user)
	return resp, token.String()
}

func Update(user entity.User) {

}
