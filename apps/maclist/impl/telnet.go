package impl

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/ziutek/telnet"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/maclist"
	"github.com/canflyx/gosw/apps/switches"
	"github.com/canflyx/gosw/conf"
)

func NewARPTelnet(sw *SwitchesConfig) (error, interface{}) {
	m := newManager()
	m.addProcess(&TelnetSw{})
	m.addProcess(&ByteCheck{Sw: sw.Ip})
	m.addProcess(&SaveARP{
		rep: app.GetInternalApp("maclist-impl").(*MacListServiceImpl),
	})
	return m.run(sw)
}
func NewMacTelnet(sw *SwitchesConfig) (error, interface{}) {
	m := newManager()
	m.addProcess(&TelnetSw{})
	m.addProcess(&ByteCheck{})
	m.addProcess(&Dup{})
	m.addProcess(&SaveMac{
		rep: app.GetInternalApp("maclist-impl").(*MacListServiceImpl),
	})
	return m.run(sw)
}
func NewCuTelnet(sw *SwitchesConfig) (error, interface{}) {
	m := newManager()
	m.addProcess(&TelnetSw{})
	m.addProcess(&SaveLog{
		sw:  *sw,
		rep: app.GetInternalApp("maclist-impl").(*MacListServiceImpl),
	})
	return m.run(sw)
}

type IProcessor interface {
	Process(params interface{}) (error, interface{})
}

type Manager struct {
	ps []IProcessor
}

func newManager() *Manager {
	return &Manager{}
}

// 可以将需要处理的方法都加到一个切片中
func (m *Manager) addProcess(process IProcessor) {
	m.ps = append(m.ps, process)
}

func (m *Manager) run(params interface{}) (error, interface{}) {
	ret := params
	var err error
	for _, v := range m.ps {
		err, ret = v.Process(ret)
		if err != nil {
			return err, nil
		}
	}
	return nil, ret
}

// 读取配置入口：交换机配置，交换机命令，标志项 0: arp 1: mac 模式 2: 批量执行命令返回模式
type SwitchesConfig struct {
	switches.Switches
	BrandCMD conf.TelnetCmd
	Flag     int
	TimeOut  int64
}
type TelnetSw struct {
	tc      *telnet.Conn
	timeout time.Duration
	log     logger.Logger
}

func (tel *TelnetSw) Process(params interface{}) (error, interface{}) {
	if sw, ok := params.(*SwitchesConfig); !ok {
		return errors.New("TelnetSw input type error"), nil
	} else {
		tel.log = zap.L().Named("telnet")
		cmd := sw.BrandCMD
		tel.timeout = time.Duration(sw.TimeOut) * time.Second
		// 返回 telnet.Conn 实例
		var err error
		tel.tc, err = telnet.Dial("tcp", sw.Ip+":23")
		if err != nil {
			tel.log.Errorf("telnet sw:%s dial error", sw.Ip)
			return nil, err
		}
		tel.tc.SetUnixWriteMode(false)
		_ = tel.expect(cmd.UserFlag)
		tel.sendln(sw.User)
		_ = tel.expect(cmd.PasswordFlag)
		tel.sendln(sw.Password)
		_ = tel.expect("[Y/N]:")
		tel.sendln("n")
		err = tel.expect(cmd.LoginFlag)
		if err != nil {
			return nil, err
		}
		if cmd.EnableCmd != "" {
			tel.sendln(cmd.EnableCmd)
			err = tel.expect(cmd.EnableFlag)
			if err != nil {
				return nil, err
			}
		}
		if len(cmd.Cmds) > 0 {
			for _, c := range cmd.Cmds {
				tel.sendln(c.CMD)
				err = tel.expect(c.CMDFlag)
				tel.checkErr(c.CMD, err)
			}
		}
		if sw.Flag == 2 {
			var datas bytes.Buffer
			for _, c := range cmd.ReadCmd {
				err = tel.sendln(c.CMD)
				tel.checkErr(c.CMD, err)
				data, _ := tel.tc.ReadBytes([]byte(c.CMDFlag)[0])
				datas.Write(data)
			}
			return nil, datas.Bytes()
		}
		if sw.Flag == 1 {
			err = tel.sendln(cmd.CoreCmd)
			tel.checkErr(cmd.CoreCmd, err)
		}
		if sw.Flag == 0 {
			err = tel.sendln(cmd.AccessCmd)
			tel.checkErr(cmd.AccessCmd, err)
		}

		tlData, err := tel.tc.ReadBytes([]byte(cmd.ReadFlag)[0])
		if len(cmd.ExitCmds) > 0 {
			for _, c := range cmd.ExitCmds {
				tel.sendln(c.CMD)
				err = tel.expect(c.CMDFlag)
				tel.checkErr(c.CMD, err)
			}
		}
		tel.log.Info(string(tlData[:20]))
		return nil, tlData
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

func (tel *TelnetSw) checkErr(info string, err error) {
	if err != nil {
		tel.log.Error("[ERROR] ", info, "失败.", err)
	} else {
		tel.log.Info("[INFO] ", info, "成功.")
	}
}

type ByteCheck struct {
	Sw string
}

// 未处理 mac 未加入交换机IP BUG
func (b *ByteCheck) Process(params interface{}) (error, interface{}) {
	if v, ok := params.([]byte); !ok {
		return errors.New("ByteCheck input type error"), nil
	} else {
		lines := strings.Split(string(v), "\n")
		okSlice := []*maclist.MacList{}
		for _, line := range lines {
			lineArray := strings.Split(line, " ")
			new := maclist.NewMacList()
			for i := 0; i < len(lineArray); i++ {
				if lineArray[i] == "" {
					continue
				}
				if strings.ContainsAny(lineArray[i], "GEXIF") {
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
				if new.ARPIP == "" {
					new.SwitchIp = b.Sw
				}
				// new.SwitchIp = swIP
				okSlice = append(okSlice, new)
			}
		}
		return nil, okSlice
	}
}

type Dup struct {
}

func (d *Dup) Process(params interface{}) (error, interface{}) {
	if v, ok := params.([]*maclist.MacList); !ok {
		return errors.New("Duplicate input type error"), nil
	} else {
		var ret []*maclist.MacList
		var c = make(map[string]int)
		for _, s := range v {

			v, ok := c[s.Port]
			if !ok {
				c[s.Port] = 1
			} else {
				c[s.Port] = v + 1
			}

		}
		for _, s := range v {
			v, _ := c[s.Port]
			if v == 1 {
				ret = append(ret, s)
			}
		}
		return nil, ret
	}
}

type SaveMac struct {
	rep maclist.Repositoryer
}

func (s *SaveMac) Process(params interface{}) (error, interface{}) {
	if v, ok := params.([]*maclist.MacList); !ok {
		return errors.New("saveMac input type error"), nil
	} else {
		var result []*maclist.MacAddrs
		for _, d := range v {
			result = append(result, &maclist.MacAddrs{MacAddress: d.MacAddress, Port: d.Port, SwitchIp: d.SwitchIp})
		}
		return s.rep.SaveMac(result), nil
	}

}

type SaveARP struct {
	rep maclist.Repositoryer
}

func (s *SaveARP) Process(params interface{}) (error, interface{}) {
	if v, ok := params.([]*maclist.MacList); !ok {
		return errors.New("saveArp input type error"), nil
	} else {
		var result []*maclist.ARPList
		for _, d := range v {
			result = append(result, &maclist.ARPList{ARPIP: d.ARPIP, MacAddress: d.MacAddress})
		}
		return s.rep.SaveARP(result), nil
	}

}

type SaveLog struct {
	rep maclist.Repositoryer
	sw  SwitchesConfig
}

func (s *SaveLog) Process(param interface{}) (error, interface{}) {
	log := ""
	if v, ok := param.(([]byte)); ok {
		log = string(v)
	} else {
		log = fmt.Sprintf("%v", param)

	}
	return s.rep.SaveLog(&maclist.ScanLog{
		SwitchID: s.sw.ID,
		Log:      log,
	}), nil
}

func IsMac(mac string) bool {
	return len(strings.Split(mac, "-")) == 3 || len(strings.Split(mac, ".")) == 3
}
