package impl

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/ziutek/telnet"

	"github.com/canflyx/gosw/apps/tools"
	"github.com/canflyx/gosw/conf"
)

type TelnetSw struct {
	tc      *telnet.Conn
	timeout time.Duration
	log     *slog.Logger
}

func (tel *TelnetSw) Login(sw SwitchesConfig, UserFlag, PasswordFlag, LoginFlag string) error {
	var err error
	tel.tc, err = telnet.Dial("tcp", sw.Ip+":23")
	if err != nil {
		return err
	}
	tel.tc.SetUnixWriteMode(true)
	err = tel.expect(UserFlag)
	if err != nil {
		return err
	}
	tel.sendln(sw.User)
	err = tel.expect(PasswordFlag)
	if err != nil {
		return err
	}
	tel.sendln(sw.Password)
	err = tel.expect(LoginFlag)
	if err != nil {
		return err
	}
	tel.log.Info(sw.Ip + " login success")
	return err

}
func (tel *TelnetSw) Process(params interface{}) (interface{}, error) {
	if sw, ok := params.(*SwitchesConfig); !ok {
		return errors.New("TelnetSw input type error"), nil
	} else {
		tel.log = conf.GetNameLog("telnet")
		cmd := sw.BrandCMD
		tel.timeout = time.Duration(sw.TimeOut) * time.Second
		// 返回 telnet.Conn 实例
		err := tel.Login(*sw, cmd.UserFlag, cmd.PasswordFlag, cmd.LoginFlag)

		// err = tel.expect(cmd.LoginFlag)
		if err != nil {
			return nil, err
		}
		// 对自定义进行分别处理,readbytes 或 expect 截断byte
		if sw.Flag == 2 {
			var datas bytes.Buffer
			for _, c := range cmd.UserCmd {
				tel.sendln(c.CMD)
				if !strings.ContainsAny(c.CMDFlag, ">]#") {
					c.CMDFlag = ">"
				}
				data, _ := tel.tc.ReadBytes([]byte(c.CMDFlag)[0])
				datas.Write(data)
			}
			return datas.Bytes(), nil
		}

		cmds := tools.SplicCmd(cmd, sw.Flag)
		for _, c := range cmds {
			_ = tel.sendln(c.CMD)
			if !strings.ContainsAny(c.CMDFlag, ">]#") {
				c.CMDFlag = ">"
			}
			_ = tel.expect(c.CMDFlag)
		}
		if !strings.ContainsAny(cmd.ReadFlag, ">]#") {
			cmd.ReadFlag = ">"
		}
		tlData, _ := tel.tc.ReadBytes([]byte(cmd.ReadFlag)[0])
		if len(cmd.ExitCmd) > 0 {
			for _, c := range cmd.ExitCmd {
				tel.sendln(c.CMD)
			}
		}
		tel.log.Info("telnet return data len = " + fmt.Sprint(len(tlData)))
		return tlData, nil
	}
}

func (tel *TelnetSw) sendln(s string) (err error) {
	err = tel.tc.SetWriteDeadline(time.Now().Add(tel.timeout))
	if err != nil {
		return err
	}
	// 将命令变成[]byte,并在最后面加上 \n
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err = tel.tc.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
func (tel *TelnetSw) expect(d ...string) (err error) {
	tel.tc.SetReadDeadline(time.Now().Add(tel.timeout))
	err = tel.tc.SkipUntil(d...)
	if err != nil {
		return err
	}
	return nil
}
