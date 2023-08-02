package mapper

import (
	"github.com/KingFarGrace/CollabSearch/server/conf"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"xorm.io/xorm"
)

func getEngine() *xorm.Engine {
	return conf.GetMySQLEngine()
}

func ExistEmail(email string) bool {
	engine := getEngine()
	exist, err := engine.Table("user").Where("email=?", email).Exist()
	if err != nil {
		util.ErrorLogger(err, "ExistEmail()")
		return false
	}
	return exist
}

func InsertUser(user entity.User) {
	engine := getEngine()
	_, err := engine.Insert(user)
	if err != nil {
		util.ErrorLogger(err, "engine.Insert(user)")
		return
	}
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
