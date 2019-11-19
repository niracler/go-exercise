package base

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"go-exercise/redpacket/infra"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ini.NewIniFileConfigSource("config.ini")
}
