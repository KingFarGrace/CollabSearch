package router

import "github.com/KingFarGrace/CollabSearch/server/entity"

type WorkspaceResponse struct {
	Response
	Workspaces []entity.Workspace
}

func (receiver *WorkspaceResponse) SetReturnObjs(workspaces interface{}) {
	receiver.Workspaces = workspaces.([]entity.Workspace)
}

func (receiver *WorkspaceResponse) SetReturnObj(workspace interface{}) {
	receiver.Workspaces = make([]entity.Workspace, 0)
	receiver.Workspaces = append(receiver.Workspaces, workspace.(entity.Workspace))
}
