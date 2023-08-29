package entity

// User ORM model -- table 'user'
type User struct {
	Uid      int64  `xorm:"pk index" json:"uid" validate:"required,number,max=9223372036854775807"`
	Email    string `xorm:"notnull" json:"email" validate:"omitempty,email"`
	Username string `xorm:"notnull" json:"username" validate:"omitempty,alphanum|ascii,min=1,max=32"`
	Password string `xorm:"notnull" json:"password" validate:"omitempty,alphanum|ascii,min=6,max=16"`
	Profile  string `json:"profile" validate:"omitempty,alphanumunicode|ascii"`
	Avatar   string `xorm:"mediumtext" json:"avatar" validate:"omitempty,base64url"`
}

// RegisterJSON is a JSON entity to receive and pass user's sign up data.
type RegisterJSON struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,alphanum|ascii,min=1,max=32"`
	Password        string `json:"password" validate:"required,alphanum|ascii,min=6,max=16"`
	ConfirmPassword string `json:"confirm" validate:"eqcsfield=Password"`
}

// LoginJSON is a JSON entity to receive and pass user's sign in data.
type LoginJSON struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum|ascii,min=6,max=16"`
}

// UidJSON is a JSON entity to receive and pass user ID.
type UidJSON struct {
	Uid int64 `json:"uid" validate:"required,number,max=9223372036854775807"`
}

// TokenClaims is a JSON entity used to initialize a token
type TokenClaims struct {
	Email string `json:"email"`
}
