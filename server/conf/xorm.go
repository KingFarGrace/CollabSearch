package conf

import (
	"github.com/KingFarGrace/CollabSearch/server/util"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var xormConf = GetConfig().GetXormConf()
var engine *xorm.Engine

// InitPersistenceLayer to connect MySQL database.
func InitPersistenceLayer() {
	url := util.Concat(
		xormConf.Username,
		":",
		xormConf.Password,
		"@tcp(",
		xormConf.Hostname,
		":",
		xormConf.Port,
		")/",
		xormConf.Database,
		"?charset=",
		xormConf.Charset,
		"&parseTime=",
		xormConf.ParseTime,
		"&loc=",
		xormConf.Loc,
	)
	var err error
	engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		util.PlainErrorLogger(err, "xorm.NewEngine()")
		util.LaunchLogger("xorm", false)
		return
	}
	util.LaunchLogger("xorm", true)
}

// GetEngine
// TODO: test
func GetEngine() *xorm.Engine {
	return engine
}
