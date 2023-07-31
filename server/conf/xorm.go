package conf

import (
	"github.com/KingFarGrace/CollabSearch/server/util"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// xorm configuration
const (
	username = "root"
	password = "132546"
	addr     = "localhost:3306"
	database = "collabsearch"
	charset  = "utf-8"
)

var engine *xorm.Engine

// InitPersistenceLayer to connect MySQL database.
func InitPersistenceLayer() {
	url := util.Concat(username, ":", password, "@", addr, "/", database, "?", charset)
	var err error
	engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		util.PlainErrorLogger(err, "xorm.NewEngine()")
		return
	}
}

// GetEngine
// TODO: test
func GetEngine() *xorm.Engine {
	return engine
}
