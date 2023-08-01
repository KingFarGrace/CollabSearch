package conf

import (
	"github.com/KingFarGrace/CollabSearch/server/util"
	"path/filepath"
	"sync"
)

type Config struct {
	Xorm  []Xorm  `json:"xorm"`
	Redis []Redis `json:"redis"`
	Cors  Cors    `json:"cors"`
}

type Xorm struct {
	Driver    string `json:"driver"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
	Database  string `json:"database"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parseTime"`
	Loc       string `json:"loc"`
}

type Redis struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

type Cors struct {
	AllowAllOrigins        bool     `json:"AllowAllOrigins"`
	AllowOrigins           []string `json:"AllowOrigins"`
	AllowMethods           []string `json:"AllowMethods"`
	AllowHeaders           []string `json:"AllowHeaders"`
	ExposeHeaders          []string `json:"ExposeHeaders"`
	AllowCredentials       bool     `json:"AllowCredentials"`
	MaxAge                 int64    `json:"MaxAge"`
	AllowWildcard          bool     `json:"AllowWildcard"`
	AllowBrowserExtensions bool     `json:"AllowBrowserExtensions"`
	AllowWebSockets        bool     `json:"AllowWebSockets"`
	AllowFiles             bool     `json:"AllowFiles"`
}

var configList *Config
var once sync.Once

const configRelativePath = "D:\\code\\CollabSearch\\server\\config.json"

func GetConfig() *Config {
	once.Do(func() {
		configPath, err := filepath.Abs(configRelativePath)
		if err != nil {
			util.FatalLogger(err, "filepath.Abs()")
			return
		}
		util.ParseJSONFile(configPath, &configList)
	})
	return configList
}

func (this Config) GetXormConf() Xorm {
	return this.Xorm[0]
}

func (this Config) GetXormConfs() []Xorm {
	return this.Xorm
}

func (this Config) GetRedisConf() Redis {
	return this.Redis[0]
}

func (this Config) GetRedisConfs() []Redis {
	return this.Redis
}
