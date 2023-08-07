package mapper

import "github.com/KingFarGrace/CollabSearch/server/entity"

func InsertWorkspace(workspace entity.Workspace) bool {
	engine := getEngine()
	affected, err := engine.Insert(workspace)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func SelectWorkspacesByUid(uid int64) []entity.Workspace {
	return nil
}

func SelectWorkspacesByHandler(handlerID int64) []entity.Workspace {
	engine := getEngine()
	workspaces := make([]entity.Workspace, 0)
	for _, workspace := range workspaces {
		workspace.Handler = handlerID
	}
	err := engine.Find(&workspaces)
	if err != nil {
		return nil
	}
	return workspaces
}
