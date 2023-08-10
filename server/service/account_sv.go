package service

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	"github.com/KingFarGrace/CollabSearch/server/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"strings"
)

// TODO: Give an accurate path.
var defaultAvatarPath = ""

// Register a new user.
// Receive a RegisterJSON then convert it to a User instance,
// then register the User through mapper.
func Register(registerJSON entity.RegisterJSON) *response.AccountResponse {
	if mapper.ExistEmail(registerJSON.Email) {
		msg := util.Concat("Duplicated user email: ", registerJSON.Email)
		return getFailedAccountResp(1, msg)
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
		return getSuccessAccountResp()
	} else {
		return getFailedAccountResp(2, "Failed to sign up, please try later.")
	}
}

// LoginWithoutToken is a function for no-token user to login.
// no-token user: new user or whose token was expired.
// When someone sign in the system through this function,
// they should cache the user data in localstorage.
func LoginWithoutToken(loginJSON entity.LoginJSON) (*response.AccountResponse, string) {
	var user = mapper.SelectUserByEmail(strings.ToLower(loginJSON.Email))
	if user == nil {
		msg := util.Concat("No such user. Email: ", loginJSON.Email)
		return getFailedAccountResp(4, msg), ""
	}
	if !passAuth(loginJSON.Password, user.Password) {
		return getFailedAccountResp(3, "Wrong password."), ""
	}
	token := util.GenerateJWT(strings.ToLower(loginJSON.Email), conf.Salt)
	user.Password = "" // Block user privacy data.
	return getSuccessAccountResp(*user), token.String()
}

func UpdateUser(user entity.User) *response.AccountResponse {
	if mapper.UpdateUser(user) {
		return getSuccessAccountResp()
	}
	return getFailedAccountResp(5, "Failed to update user data.")
}

func getSuccessAccountResp(users ...entity.User) *response.AccountResponse {
	resp := new(response.AccountResponse)
	resp.New(response.AccountGroupCode, 0, "Success.")
	if len(users) == 0 {
		return resp
	}
	resp.SetReturnObjs(users)
	return resp
}

func getFailedAccountResp(failureIdx int, failureMsg string) *response.AccountResponse {
	resp := new(response.AccountResponse)
	resp.New(response.AccountGroupCode, failureIdx, failureMsg)
	return resp
}

func passAuth(pwdInput, pwd string) bool {
	if pwdInput != pwd {
		return false
	}
	return true
}
