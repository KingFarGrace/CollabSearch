package entity

import "time"

type Workspace struct {
	Wid         int       `xorm:"pk index autoincr"`
	Topic       string    `xorm:"notnull"`
	Description string    `xorm:"text notnull"`
	Handler     int64     `xorm:"notnull"`
	Due         time.Time `xorm:"notnull"`
}
