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
	DbTimeZone string

	AgentAppid  string
	AgentSecret string
)

func init() {
	file, err := ini.Load("config/releaseConfig.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查路径")
	}
	LoadServer(file)
	LoadData(file)
	LoadAgent(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
	JwtKey = file.Section("server").Key("JwtKey").String()
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").String()
	DbHost = file.Section("database").Key("DbHost").String()
	DbPort = file.Section("database").Key("DbPort").String()
	DbUser = file.Section("database").Key("DbUser").String()
	DbPassWord = file.Section("database").Key("DbPassWord").String()
	DbName = file.Section("database").Key("DbName").String()
	DbTimeZone = file.Section("database").Key("DbTimeZone").String()
}

func LoadAgent(file *ini.File) {
	AgentAppid = file.Section("agent").Key("Appid").String()
	AgentSecret = file.Section("agent").Key("Secret").String()
}
