package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AgentAppid  string
	AgentSecret string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查路径")
	}
	LoadServer(file)
	LoadData(file)
	LoadAgent(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("bdfskd983usdhf03")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("523672645")
	DbName = file.Section("database").Key("DbName").MustString("star")
}

func LoadAgent(file *ini.File) {
	AgentAppid = file.Section("agent").Key("Appid  ").MustString("wx5bbeae76ae17f1da")
	AgentSecret = file.Section("agent").Key("Secret ").MustString("2c8e552595f7ec4331661e7fde1425a1")
}
