package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	"github.com/KingFarGrace/CollabSearch/server/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

// CreateWorkspace can create a new workspace.
func CreateWorkspace(workspace entity.Workspace) *response.WorkspaceResponse {
	var cvt bool
	workspace.Due, cvt = util.String2Time(workspace.DueString)
	if !cvt || !mapper.InsertWorkspace(workspace) {
		return getFailedWorkspaceResp(1, "Failed to create the workspace.")
	}
	return getSuccessWorkspaceResp()
}

// JoinWorkspace allow user to join an existing workspace.
func JoinWorkspace(uw entity.UserWorkspace) *response.WorkspaceResponse {
	if mapper.InsertUserWorkspace(uw) {
		return getSuccessWorkspaceResp()
	}
	return getFailedWorkspaceResp(2, "Failed to join the workspace.")
}

// EnterWorkspace allow user to enter a joined workspace.
func EnterWorkspace(wid int) *response.WorkspaceResponse {
	return nil
}

// GetWorkspaces can get workspaces that user joined or created.
func GetWorkspaces(id int64, isHandler bool) *response.WorkspaceResponse {
	var workspaces []entity.Workspace
	if isHandler {
		workspaces = mapper.SelectWorkspacesByHandler(id)
		if workspaces != nil {
			return getSuccessWorkspaceResp(workspaces...)
		}
	} else {
		workspaces = mapper.SelectWorkspacesByUid(id)
		if workspaces != nil {
			return getSuccessWorkspaceResp(workspaces...)
		}
	}
	return getFailedWorkspaceResp(4, "Failed to get workspaces information.")
}

// UpdateWorkspace allow handler to update the information of a workspace.
func UpdateWorkspace(workspace entity.Workspace) *response.WorkspaceResponse {
	var cvt bool
	workspace.Due, cvt = util.String2Time(workspace.DueString)
	if !cvt || !mapper.UpdateWorkspace(workspace) {
		return getFailedWorkspaceResp(3, "Failed to update workspace information.")
	}
	return getSuccessWorkspaceResp()
}

func getSuccessWorkspaceResp(workspaces ...entity.Workspace) *response.WorkspaceResponse {
	resp := new(response.WorkspaceResponse)
	resp.New(response.WorkspaceGroupCode, 0, "Success")
	if len(workspaces) == 0 {
		return resp
	}
	resp.SetReturnObjs(workspaces)
	return resp
}

func getFailedWorkspaceResp(failureIdx int, failureMsg string) *response.WorkspaceResponse {
	resp := new(response.WorkspaceResponse)
	resp.New(response.WorkspaceGroupCode, failureIdx, failureMsg)
	return resp
}
