package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/response"
)

func getSuccessResultResp(results ...entity.Result) *response.ResultResponse {
	resp := new(response.ResultResponse)
	resp.New(response.ResultGroupCode, 0, "Success.")
	if len(results) == 0 {
		return resp
	}
	resp.SetReturnObjs(results)
	return resp
}

func getFailedResultResp(failureIdx int, failureMsg string) *response.ResultResponse {
	resp := new(response.ResultResponse)
	resp.New(response.ResultGroupCode, failureIdx, failureMsg)
	return resp
}
