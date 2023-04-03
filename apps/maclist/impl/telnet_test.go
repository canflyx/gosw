package impl

import (
	"fmt"
	"testing"

	"github.com/canflyx/gosw/apps/switches"
)

func TestParse(t *testing.T) {
	core := 1
	a := switches.Switches{
		Ip:       "172.17.80.1",
		User:     "daika",
		Password: "daika2018",
		IsCore:   &core,
	}
	config := &SwitchesConfig{
		Switches: a,
		// BrandCMD: map[string]conf.TelnetCmd{"default": {
		// 	Brand:        "default",
		// 	UserFlag:     "name:",
		// 	PasswordFlag: "ssword:",
		// 	LoginFlag:    ">",
		// 	EnableCmd:    "sys",
		// 	EnableFlag:   "]",
		// 	Cmds:         []conf.CMD{{"user-interface vty 0 4", "]"}, {"screen-length 0", "]"}},
		// 	ReadCmd:      []conf.CMD{{"dis ver", "]"}, {"dis users", "]"}},
		// 	AccessCmd:    "dis mac-add",
		// 	CoreCmd:      "dis arp",
		// 	ReadFlag:     "]",
		// 	ExitCmds:     []conf.CMD{{"screen-length 50", "]"}},
		// }},
		Flag:    2,
		TimeOut: 5,
	}
	err, data := NewCuTelnet(config)

	fmt.Printf("%v,%s", err, data)

}
