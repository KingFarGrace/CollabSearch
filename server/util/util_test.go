package util

import (
	"fmt"
	"strconv"
	"testing"
)

type testConfig struct {
	Xorm  []testXorm  `json:"xorm"`
	Redis []testRedis `json:"redis"`
	Cors  testCors    `json:"cors"`
}

type testXorm struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

type testRedis struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database int    `json:"database"`
}
type testCors struct {
	AllowOrigins     []string `json:"AllowOrigins"`
	AllowMethods     []string `json:"AllowMethods"`
	ExposeHeaders    []string `json:"ExposeHeaders"`
	AllowCredentials bool     `json:"AllowCredentials"`
}

func TestConcat(t *testing.T) {
	testStr := Concat("This ", "is ", "a ", false, " string. ", "No.", 1)
	println("Concat string: ", testStr)
}

func TestGetSnowflakeID(t *testing.T) {
	testID := GetSnowflakeID(1)
	println("ID: ", testID.Int64())
	println("ID: ", strconv.FormatInt(testID.Int64(), 2))
}

func TestParseJSONFile(t *testing.T) {
	var test testConfig
	ParseJSONFile("../conf/config.json", &test)
	fmt.Printf("%v\n", test)
}
