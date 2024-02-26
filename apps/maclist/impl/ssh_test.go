package impl

import (
	"fmt"
	"testing"

	"github.com/canflyx/gosw/apps/tools"
	"github.com/canflyx/gosw/conf"
)

func TestShhProcess(t *testing.T) {
	core := 1
	conf.LoadConfigFromEnv()
	confs := conf.C().TelnetCmd()
	v, ok := confs["default"]
	if !ok {
		fmt.Print("config error")
	}
	v.UserCmd = tools.FormatCmd("dis users;dis vlan")
	v.ReadFlag = ">"
	v.UserFlag = ">"
	conf.Zlog = tools.SlogInit()

	// c := SwitchesConfig{Flag: 2, BrandCMD: v}
	// c := SwitchesConfig{Flag: 0, TimeOut: 10, BrandCMD: v}
	c := SwitchesConfig{Flag: 1, TimeOut: 10, BrandCMD: v}
	c.Ip = "172.30.49.17"
	c.User = "admin"
	c.Password = "123456"
	c.IsCore = &core

	d := SshConfig{}

	e, _ := d.Process(&c)

	if v, ok := e.([]byte); !ok {
		fmt.Println("ByteCheck input type error")
	} else {
		d.log.Info(string(v))
	}
}
