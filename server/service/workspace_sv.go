package service

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/mapper"
	router "github.com/KingFarGrace/CollabSearch/server/router/response"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

var emptyWorkspaces = make([]entity.Workspace, 0)

// CreateWorkspace can create a new workspace.
func CreateWorkspace(workspace entity.Workspace) *router.WorkspaceResponse {
	resp := new(router.WorkspaceResponse)
	var err error
	workspace.Due, err = util.String2Time(workspace.DueString)
	if err != nil || !mapper.InsertWorkspace(workspace) {
		resp.New(router.WorkspaceGroupCode, 1, "Failed to create the workspace.")
		resp.SetReturnObjs(emptyWorkspaces)
		return resp
	}
	resp.New(router.WorkspaceGroupCode, 0, "Success.")
	resp.SetReturnObjs(emptyWorkspaces)
	return resp
}

// JoinWorkspace allow user to join an existing workspace.
func JoinWorkspace(uw entity.UserWorkspace) *router.WorkspaceResponse {
	resp := new(router.WorkspaceResponse)
	if mapper.InsertUserWorkspace(uw) {
		resp.New(router.WorkspaceGroupCode, 0, "Success.")
		resp.SetReturnObjs(emptyWorkspaces)
		return resp
	}
	resp.New(router.WorkspaceGroupCode, 2, "Failed to join the workspace.")
	resp.SetReturnObjs(emptyWorkspaces)
	return resp
}

// EnterWorkspace allow user to enter a joined workspace.
func EnterWorkspace(wid int) *router.WorkspaceResponse {
	return nil
}

// GetWorkspaces can get workspaces that user joined or created.
func GetWorkspaces(id int64, isHandler bool) *router.WorkspaceResponse {
	resp := new(router.WorkspaceResponse)
	var workspaces []entity.Workspace
	if isHandler {
		workspaces = mapper.SelectWorkspacesByHandler(id)
		if workspaces != nil {
			resp.New(router.WorkspaceGroupCode, 0, "Success.")
			resp.SetReturnObjs(workspaces)
			return resp
		}
	} else {
		workspaces = mapper.SelectWorkspacesByUid(id)
		if workspaces != nil {
			resp.New(router.WorkspaceGroupCode, 0, "Success.")
			resp.SetReturnObjs(workspaces)
			return resp
		}
	}
	resp.New(router.WorkspaceGroupCode, 4, "Failed to get workspaces information.")
	resp.SetReturnObjs(emptyWorkspaces)
	return resp
}

// UpdateWorkspace allow handler to update the information of a workspace.
func UpdateWorkspace(workspace entity.Workspace) *router.WorkspaceResponse {
	resp := new(router.WorkspaceResponse)
	var err error
	workspace.Due, err = util.String2Time(workspace.DueString)
	if err != nil || !mapper.UpdateWorkspace(workspace) {
		resp.New(router.WorkspaceGroupCode, 3, "Failed to update workspace information.")
		resp.SetReturnObjs(emptyWorkspaces)
		return resp
	}
	resp.New(router.WorkspaceGroupCode, 0, "Success.")
	resp.SetReturnObjs(emptyWorkspaces)
	return resp
}
