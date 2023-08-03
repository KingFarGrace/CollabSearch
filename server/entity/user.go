package entity

// User ORM model -- table 'user'
type User struct {
	Uid      int64  `xorm:"pk index" json:"uid"`
	Email    string `xorm:"notnull" json:"email"`
	Username string `xorm:"notnull" json:"username"`
	Password string `xorm:"notnull" json:"password"`
	Profile  string `json:"profile"`
	Avatar   string `xorm:"notnull" json:"avatar"`
}

// RegisterJSON is a JSON obj to receive user's sign up data.
type RegisterJSON struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,alphanum,min=1,max=32"`
	Password        string `json:"password" validate:"required,alphanumunicode,min=6,max=16"`
	ConfirmPassword string `json:"confirm" validate:"eqcsfield=Password"`
}
