package main

import (
	"fmt"
	"github.com/tietang/dbx"
	"github.com/tietang/go-utils"
	"go-exercise/redpacket/infra"
	"go-exercise/redpacket/infra/base"
)

func main() {
	x := dbx.Database{}
	fmt.Println(x)
	utils.GetAllIP()

	// 读取配置文件
	propsStarter := base.PropsStarter{}
	propsStarter.Init(infra.StarterContext{}) //初始化配置文件
	conf := base.Props()
	password := conf.GetDefault("mysql.password", "123456")
	fmt.Println(password)
}
