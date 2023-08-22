package entity

import "time"

type Workspace struct {
	Wid         int       `xorm:"pk index autoincr" json:"wid" validate:"omitempty,number"`
	Topic       string    `xorm:"notnull" json:"topic" validate:"required,alphanumunicode|ascii,min=1,max=255"`
	Description string    `xorm:"text notnull" json:"description" validate:"required,alphanumunicode|ascii,min=1"`
	Handler     int64     `xorm:"notnull" json:"handler" validate:"required,number"`
	DueString   string    `xorm:"-" json:"due" validate:"required,number|ascii"`
	Due         time.Time `xorm:"notnull" json:"-"`
}
