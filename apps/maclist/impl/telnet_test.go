package impl

import (
	"fmt"
	"testing"
	"time"

	"github.com/canflyx/gosw/apps/tools"
	"github.com/canflyx/gosw/conf"
)

func TestProcess(t *testing.T) {
	core := 1
<<<<<<< HEAD
	conf.LoadConfigFromEnv()
	confs := conf.C().TelnetCmd()
	v, ok := confs["default"]
	if !ok {
		fmt.Print("config error")
=======
	a := switches.Switches{
		Ip:       "172.17.80.1",
		User:     "dai",
		Password: "dai2018",
		IsCore:   &core,
>>>>>>> 2f3aec7f5f955a6e829def2cfa6a70188d4a36b7
	}
	// v.UserCmd = tools.FormatCmd("dis users;dis esn")
	v.ReadFlag = ">"
	v.UserFlag = ">"
	conf.Zlog = tools.SlogInit()

	// c := SwitchesConfig{Flag: 2, TimeOut: 10, BrandCMD: v}
	// c := SwitchesConfig{Flag: 0, TimeOut: 10, BrandCMD: v}
	c := SwitchesConfig{Flag: 1, TimeOut: 10, BrandCMD: v}
	c.Ip = "172.30.61.1"
	c.User = "admin"
	c.Password = "123456"
	c.IsCore = &core

	d := TelnetSw{}

	d.timeout = 10 * time.Second
	e, _ := d.Process(&c)

	if _, ok := e.([]byte); !ok {
		fmt.Println("ByteCheck input type error")
	} else {
		d.log.Info("ok")
	}
}
