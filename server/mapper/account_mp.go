package mapper

import (
	"fmt"
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
	affected, err := engine.Insert(user)
	if err != nil || affected == 0 {
		util.ErrorLogger(err, "engine.Insert(user)")
		return false
	}
	return true
}

func SelectUserByUid(uid int64) *entity.User {
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

func SelectUserByEmail(email string) *entity.User {
	engine := getEngine()
	user := new(entity.User)
	user.Email = email
	notEmpty, err := engine.Get(user)
	if err != nil {
		util.ErrorLogger(err, "SelectUserByEmail()")
		return nil
	}
	if notEmpty {
		return user
	} else {
		return nil
	}
}

// SelectUserByEmailLike
// cmd: SELECT * FROM user WHERE email LIKE %key%
func SelectUserByEmailLike(key string) []entity.User {
	engine := getEngine()
	users := make([]entity.User, 0)
	cond := util.Concat("%", key, "%")
	err := engine.Where("email LIKE ?", cond).Find(&users)
	if err != nil {
		return nil
	}
	return users
}

func SelectUsersByUsername(username string) []entity.User {
	return nil
}

func UpdateUser(newData entity.User) bool {
	engine := getEngine()
	_, err := engine.ID(newData.Uid).Update(&newData)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//func UpdateAvatar(uid int64, newAvatar []uint8) bool {
//	engine := getEngine()
//	user := new(entity.User)
//	_, err := engine.ID(uid).Update(&user)
//	if err != nil {
//		return false
//	}
//	return true
//}

func DeleteUser(user entity.User) bool {
	engine := getEngine()
	affected, err := engine.ID(user.Uid).Delete(&user)
	if err != nil || affected == 0 {
		return false
	}
	return true
}
