package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main() {
	file := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileConfigSource(file)
	password := conf.GetDefault("mysql.password", "123456")
	fmt.Println(password)
}
