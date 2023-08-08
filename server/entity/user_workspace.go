package entity

type UserWorkspace struct {
	Uid int64 `json:"uid" xorm:"pk"`
	Wid int   `json:"wid" xorm:"pk"`
}

type UserWorkspaceJoin struct {
	User      `xorm:"extends"`
	Workspace `xorm:"extends"`
}
