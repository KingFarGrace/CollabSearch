package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/response"
)

func getSuccessNoteResp(notes ...entity.Note) *response.NoteResponse {
	resp := new(response.NoteResponse)
	resp.New(response.NoteGroupCode, 0, "Success.")
	if len(notes) == 0 {
		return resp
	}
	resp.SetReturnObjs(notes)
	return resp
}

func getFailedNoteResp(failureIdx int, failureMsg string) *response.NoteResponse {
	resp := new(response.NoteResponse)
	resp.New(response.NoteGroupCode, failureIdx, failureMsg)
	return resp
}
