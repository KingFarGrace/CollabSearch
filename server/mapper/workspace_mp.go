package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
)

func InsertWorkspace(workspace entity.Workspace) bool {
	engine := getEngine()
	user := new(entity.User)
	get, err := engine.Table("user").ID(workspace.Handler).Get(user)
	if err != nil || !get {
		return false
	}
	affected, err := engine.Insert(&workspace)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func InsertUserWorkspace(uw entity.UserWorkspace) bool {
	engine := getEngine()
	affected, err := engine.Insert(&uw)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func SelectWorkspacesByUid(uid int64) []entity.Workspace {
	engine := getEngine()
	joins := make([]entity.UserWorkspaceJoin, 0)
	err := engine.Table("workspace").
		Join("INNER", "user_workspace", "workspace.wid=user_workspace.wid").
		Join("INNER", "user", "user_workspace.uid=user.uid").
		Where("user.uid=?", uid).Find(&joins)
	if err != nil {
		return nil
	}
	workspaces := make([]entity.Workspace, 0)
	for _, join := range joins {
		workspaces = append(workspaces, join.Workspace)
	}
	return workspaces
}

func SelectWorkspacesByHandler(handlerID int64) []entity.Workspace {
	engine := getEngine()
	workspaces := make([]entity.Workspace, 0)
	err := engine.Where("handler=?", handlerID).Find(&workspaces)
	if err != nil {
		return nil
	}
	return workspaces
}

func UpdateWorkspace(workspace entity.Workspace) bool {
	engine := getEngine()
	affected, err := engine.ID(workspace.Wid).Update(workspace)
	if err != nil || affected == 0 {
		return false
	}
	return true
}
