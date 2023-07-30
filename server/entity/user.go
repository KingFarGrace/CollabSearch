package entity

type User struct {
	Uid      int64  `xorm:"pk index"`
	Email    string `xorm:"notnull"`
	Username string `xorm:"notnull"`
	Password string `xorm:"notnull"`
	Profile  string
	Avatar   string `xorm:"notnull"`
}
