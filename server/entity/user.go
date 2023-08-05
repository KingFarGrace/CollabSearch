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

// RegisterJSON is a JSON entity to receive and pass user's sign up data.
type RegisterJSON struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,alphanum,min=1,max=32"`
	Password        string `json:"password" validate:"required,alphanumunicode,min=6,max=16"`
	ConfirmPassword string `json:"confirm" validate:"eqcsfield=Password"`
}

// LoginJSON is a JSON entity to receive and pass user's sign in data.
type LoginJSON struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"Password" validate:"required,alphanumunicode,min=6,max=16"`
}

// TokenClaims is a JSON entity used to initialize a token
type TokenClaims struct {
	Email string `json:"email"`
}
