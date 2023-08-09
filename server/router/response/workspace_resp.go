package router

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

type WorkspaceResponse struct {
	Response
	Workspaces []entity.Workspace
}

func (receiver *WorkspaceResponse) SetReturnObjs(objs interface{}) {
	if workspaces, ok := objs.([]entity.Workspace); ok {
		receiver.Workspaces = workspaces
	} else {
		util.WarnLogger("SetReturnObjs()", "Failed to set return objs.")
	}
}
