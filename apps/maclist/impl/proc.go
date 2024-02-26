package impl

import (
	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/switches"
	"github.com/canflyx/gosw/conf"
)

// 读取配置入口：交换机配置，交换机命令，标志项 0: arp 1: mac 模式 2: 批量执行命令返回模式
type SwitchesConfig struct {
	switches.Switches
	BrandCMD conf.TelnetCmd
	Flag     int
	TimeOut  int64
}

func NewARPTelnet(sw *SwitchesConfig) (interface{}, error) {
	m := newManager()
	if *sw.SwType == 0 {
		m.addProcess(&TelnetSw{})
	} else {
		m.addProcess(&SshConfig{})
	}
	m.addProcess(&ByteCheck{Sw: sw.Ip})
	m.addProcess(&SaveARP{
		rep: app.GetInternalApp("maclist-impl").(*MacListServiceImpl),
	})
	return m.run(sw)
}
func NewMacTelnet(sw *SwitchesConfig) (interface{}, error) {
	m := newManager()
	if *sw.SwType == 0 {
		m.addProcess(&TelnetSw{})
	} else {
		m.addProcess(&SshConfig{})
	}
	m.addProcess(&ByteCheck{Sw: sw.Ip})
	m.addProcess(&Dup{})
	m.addProcess(&SaveMac{
		rep: app.GetInternalApp("maclist-impl").(*MacListServiceImpl),
	})
	return m.run(sw)
}
func NewCuTelnet(sw *SwitchesConfig) (interface{}, error) {
	m := newManager()
	if *sw.SwType == 0 {
		m.addProcess(&TelnetSw{})
	} else {
		m.addProcess(&SshConfig{})
	}
	m.addProcess(&SaveLog{
		sw:  *sw,
		rep: app.GetInternalApp("maclist-impl").(*MacListServiceImpl),
	})
	return m.run(sw)
}

type IProcessor interface {
	Process(params interface{}) (interface{}, error)
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

func (m *Manager) run(params interface{}) (interface{}, error) {
	ret := params
	var err error
	for _, v := range m.ps {
		ret, err = v.Process(ret)
		if err != nil {
			log := conf.GetNameLog("process.run:")
			log.Error(err.Error())
			return nil, err
		}
	}
	return ret, nil
}
