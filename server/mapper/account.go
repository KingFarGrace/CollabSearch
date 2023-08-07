package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
)

func ExistEmail(email string) bool {
	engine := getEngine()
	exist, err := engine.Table("user").Where("email=?", email).Exist()
	if err != nil {
		util.ErrorLogger(err, "ExistEmail()")
		return false
	}
	return exist
}

func InsertUser(user entity.User) bool {
	engine := getEngine()
	_, err := engine.Insert(user)
	if err != nil {
		util.ErrorLogger(err, "engine.Insert(user)")
		return false
	}
	return true
}

func GetUserByUid(uid int64) *entity.User {
	engine := getEngine()
	user := new(entity.User)
	user.Uid = uid
	notEmpty, err := engine.Get(user)
	if err != nil {
		util.ErrorLogger(err, "GetUserById()")
		return nil
	}
	if notEmpty {
		return user
	} else {
		return nil
	}
}

func GetUserByEmail(email string) *entity.User {
	engine := getEngine()
	user := new(entity.User)
	user.Email = email
	notEmpty, err := engine.Get(user)
	if err != nil {
		util.ErrorLogger(err, "GetUserByEmail()")
		return nil
	}
	if notEmpty {
		return user
	} else {
		return nil
	}
}

// UpdateUserData can update table 'user'
func UpdateUserData(newData entity.User) bool {
	engine := getEngine()
	infected, err := engine.ID(newData.Uid).Update(newData)
	if err != nil || infected == 0 {
		return false
	}
	return true
}
