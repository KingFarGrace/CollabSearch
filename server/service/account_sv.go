package service

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	"github.com/KingFarGrace/CollabSearch/server/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"strings"
)

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
		Avatar:   "",
	}
	if success := mapper.InsertUser(user); success {
		return getSuccessAccountResp()
	} else {
		return getFailedAccountResp(2, "Failed to sign up, please try later.")
	}
}

// CheckStatus should always be called before LoginWithoutToken.
// If user has valid token, then success and return user data.
func CheckStatus(claims entity.TokenClaims) *response.AccountResponse {
	if claims.Email == "" {
		return getFailedAccountResp(4, "No such user. Email: ")
	}
	user := mapper.SelectUserByEmail(claims.Email)
	if user == nil {
		errMsg := util.Concat("No such user. Email: ", claims.Email)
		return getFailedAccountResp(4, errMsg)
	}
	return getSuccessAccountResp(*user)
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

// UpdateUser to update user data.
func UpdateUser(user entity.User) *response.AccountResponse {
	if mapper.UpdateUser(user) {
		return getSuccessAccountResp()
	}
	return getFailedAccountResp(5, "Failed to update user data.")
}

// UpdateAvatar to update user avatar.
// This is special since the request Content-Type will be different.
//func UpdateAvatar(uid int64, avatar []uint8) *response.AccountResponse {
//	if mapper.UpdateAvatar(uid, avatar) {
//		return getSuccessAccountResp()
//	}
//	return getFailedAccountResp(5, "Failed to update user avatar.")
//}

// SearchUser to get user whose email contains key.
func SearchUser(key string) *response.AccountResponse {
	users := mapper.SelectUserByEmailLike(key)
	if users == nil {
		return getFailedAccountResp(4, "No users found.")
	}
	return getSuccessAccountResp(users...)
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

// passAuth: password authentication.
func passAuth(pwdInput, pwd string) bool {
	if pwdInput != pwd {
		return false
	}
	return true
}
