package entity

type UserWorkspace struct {
	Uid int64 `json:"uid" xorm:"pk" validate:"required,number"`
	Wid int   `json:"wid" xorm:"pk" validate:"required,number"`
}

type UserWorkspaceJoin struct {
	User      `xorm:"extends"`
	Workspace `xorm:"extends"`
}
