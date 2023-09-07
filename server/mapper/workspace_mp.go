package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
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
	_, err := engine.Insert(&uw)
	if err != nil {
		util.ErrorLogger(err, "InsertUserWorkspace")
		return false
	}
	return true
}

func SelectWorkspaceByWid(wid int) *entity.Workspace {
	engine := getEngine()
	workspace := new(entity.Workspace)
	get, err := engine.ID(wid).Get(workspace)
	if err != nil || !get {
		return nil
	}
	return workspace
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

func SelectUsersByWid(wid int) []entity.User {
	engine := getEngine()
	users := make([]entity.User, 0)
	err := engine.Table("user").
		Join("INNER", "user_workspace", "user_workspace.uid=user.uid").
		Where("user_workspace.wid=?", wid).Find(&users)
	if err != nil {
		return nil
	}
	return users
}

func UpdateWorkspace(workspace entity.Workspace) bool {
	engine := getEngine()
	_, err := engine.
		ID(workspace.Wid).
		Where("handler=?", workspace.Handler).
		Update(workspace)
	if err != nil {
		return false
	}
	return true
}
