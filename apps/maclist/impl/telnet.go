package impl

import (
	"errors"
	"net"
	"strings"
	"time"

	"github.com/ziutek/telnet"

	"github.com/canflyx/gosw/apps/maclist"
	"github.com/canflyx/gosw/apps/switches"
	"github.com/canflyx/gosw/conf"
)

// telnet 相关方法
func (ms *MacListService) TelnetSwitch(sw *switches.Switches) ([]*maclist.MacList, error) {
	cmds := conf.C().TelnetCmd()
	if cmds == nil {
		ms.log.Error("config telnet command is null or error")
		return nil, errors.New("config telnet command is null or error")
	}
	var cmd conf.TelnetCmd
	v, ok := cmds[sw.Brand]
	if !ok {
		v1, ok1 := cmds["default"]
		if !ok1 {
			ms.log.Error("config not brand or default ")
			return nil, errors.New("config not brand or default ")
		}
		cmd = v1
	} else {
		cmd = v
	}
	// 返回 telnet.Conn 实例
	t, err := telnet.Dial("tcp", sw.Ip+":23")
	if err != nil {
		ms.log.Info(err)
		return nil, err
	}
	t.SetUnixWriteMode(false)
	_ = expect(t, cmd.UserFlag)
	sendln(t, sw.User)
	_ = expect(t, cmd.PasswordFlag)
	sendln(t, sw.Password)
	_ = expect(t, "[Y/N]:")
	sendln(t, "n")
	err = expect(t, cmd.LoginFlag)
	if err != nil {
		ms.checkErr("登陆:", err)
		return nil, err
	}
	if cmd.EnableCmd != "" {
		sendln(t, cmd.EnableCmd)
		err = expect(t, cmd.EnableFlag)
		if err != nil {
			ms.checkErr("特权模式:", err)
			return nil, err
		}
	}
	if len(cmd.Cmds) > 0 {
		for _, c := range cmd.Cmds {
			sendln(t, c.CMD)
			err = expect(t, c.CMDFlag)
			ms.checkErr(c.CMD, err)
		}
	}
	flag := *sw.IsCore == 1
	if flag {
		sendln(t, cmd.CoreCmd)
		ms.checkErr(cmd.CoreCmd, err)
	} else {
		sendln(t, cmd.ReadCmd)
		ms.checkErr(cmd.ReadCmd, err)
	}

	tlData, err := t.ReadBytes([]byte(cmd.ReadFlag)[0])
	ms.log.Info(string(tlData[:20]))
	if len(cmd.ExitCmds) > 0 {
		for _, c := range cmd.ExitCmds {
			sendln(t, c.CMD)
			err = expect(t, c.CMDFlag)
			ms.checkErr(c.CMD, err)
		}
	}

	// return ms.ByteToSlice(tlData, flag)
	return ms.ByteToSlice(tlData, flag, sw.Ip)

}

func (ms *MacListService) ByteToSlice(data []byte, flag bool, swIP string) ([]*maclist.MacList, error) {
	lines := strings.Split(string(data), "\n")
	okSlice := []*maclist.MacList{}
	for _, line := range lines {
		lineArray := strings.Split(line, " ")
		new := maclist.NewMacList()
		for i := 0; i < len(lineArray); i++ {
			if lineArray[i] == "" {
				continue
			}

			if strings.ContainsAny(lineArray[i], "GEXI") {
				new.Port = lineArray[i]
				continue
			}
			if IsMac(lineArray[i]) {
				new.MacAddress = lineArray[i]
				continue
			}

			if net.ParseIP(lineArray[i]) != nil {
				new.ARPIP = lineArray[i]
				continue
			}

		}
		if new.MacAddress != "" {
			new.SwitchIp = swIP
			okSlice = append(okSlice, new)
		}

	}

	if flag {
		return okSlice, nil
	}
	return Duplicate(okSlice), nil
}

func (ms *MacListService) checkErr(info string, err error) {
	if err != nil {
		ms.log.Info("[ERROR] ", info, "失败.", err)
	} else {
		ms.log.Info("[INFO] ", info, "成功.")
	}
}

func IsMac(mac string) bool {
	return len(strings.Split(mac, "-")) == 3
}

// 保留具有唯一端口的MAC
func Duplicate(b []*maclist.MacList) []*maclist.MacList {
	var ret []*maclist.MacList
	var c = make(map[string]int)
	for _, s := range b {

		v, ok := c[s.Port]
		if !ok {
			c[s.Port] = 1
		} else {
			c[s.Port] = v + 1
		}

	}
	for _, s := range b {
		v, _ := c[s.Port]
		if v == 1 {
			ret = append(ret, s)
		}
	}
	return ret
}

const timeout = 5 * time.Second

func expect(t *telnet.Conn, d ...string) (err error) {
	t.SetReadDeadline(time.Now().Add(timeout))

	err = t.SkipUntil(d...)
	if err != nil {
		return err
	}

	return nil
}

func sendln(t *telnet.Conn, s string) (err error) {
	err = t.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}
	// 将命令变成[]byte,并在最后面加上 \n
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'

	_, err = t.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
