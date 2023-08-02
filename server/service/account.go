package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	response "github.com/KingFarGrace/CollabSearch/server/router/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

func Register(form entity.User) *response.AccountResponse {
	resp := new(response.AccountResponse)
	resp.New(response.AccountGroupCode, 0, "Success.")
	if mapper.ExistEmail(form.Email) {
		resp.SetIndex(1)
		resp.SetMsg("Duplicated user email: ", form.Email)
		return resp
	}
	uid := util.GetSnowflakeID(1)
	form.Uid = uid.Int64()
	mapper.InsertUser(form)
	return resp
}
