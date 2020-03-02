package main

import (
	"fmt"
	"github.com/tietang/dbx"
	"github.com/tietang/go-utils"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"go-exercise/redpacket/infra"
)

func main() {
	x := dbx.Database{}
	fmt.Println(x)
	utils.GetAllIP()

	//加载和解析配置文件
	file := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileConfigSource(file)
	app := infra.New(conf)
	app.Start()
}
