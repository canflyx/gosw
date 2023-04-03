package impl

import (
	"context"
	"errors"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/maclist"
	"github.com/canflyx/gosw/apps/switches"
	swimpl "github.com/canflyx/gosw/apps/switches/impl"
	"github.com/canflyx/gosw/apps/tools"
	"github.com/canflyx/gosw/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// var _ maclist.Service = (*MacListService)(nil)

type MacListService struct {
	log logger.Logger
	rep maclist.Repositoryer
}

// 依据交换机 ID 数组进行扫描
func (ms *MacListService) ScanSwitch(ctx context.Context, ins maclist.ListData) error {
	if len(ins.List) < 1 {
		return errors.New("not find switch")
	}
	sw := swimpl.NewSwitchImpl()
	for i := 0; i < len(ins.List); i++ {
		sws := sw.DescById(uint(ins.List[i]))
		//
		if sws == nil {
			return errors.New("switch is null")
		}
		sws.Password, _ = tools.DecryptByAes(sws.Password)
		go ms.SaveAll(ctx, sws, ins.ReadCmd)

	}
	return nil
}

// 查询数据给 gin 使用
func (ms *MacListService) QueryMacList(ctx context.Context, req *maclist.QueryMacRequest) (*maclist.MacSet, error) {
	total, item := ms.rep.QueryByKws(req.Keyword, req.OffSet(), req.GetPageSize())
	return &maclist.MacSet{
		Total: total,
		Items: item,
	}, nil
}

// 传入交换机进行 telnet 返回数据依据类型进行保存
func (ms *MacListService) SaveAll(ctx context.Context, sw *switches.Switches, value []maclist.CMD) error {
	cmds := conf.C().TelnetCmd()
	var cmd conf.TelnetCmd
	v, ok := cmds[sw.Brand]
	if !ok {
		v1, ok1 := cmds["default"]
		if !ok1 {
			return errors.New("config not brand or default ")
		}
		cmd = v1
	} else {
		cmd = v
	}
	if len(value) > 0 {
		for _, v := range value {
			cmd.ReadCmd = append(cmd.ReadCmd, conf.CMD{CMD: v.Cmd, CMDFlag: v.Flag})
		}
		err, _ := NewCuTelnet(&SwitchesConfig{Switches: *sw, BrandCMD: cmd, Flag: 2, TimeOut: 5})
		return err
	}
	if *sw.IsCore == 1 {
		err, _ := NewARPTelnet(&SwitchesConfig{Switches: *sw, BrandCMD: cmd, Flag: 1, TimeOut: 5})
		return err
	}
	err, _ := NewMacTelnet(&SwitchesConfig{Switches: *sw, BrandCMD: cmd, Flag: 0, TimeOut: 5})
	return err

	// datas, err := ms.TelnetSwitch(sw)
	// if err != nil {
	// 	return err
	// }
	// if len(datas) < 1 {
	// 	return nil
	// }
	// // 根据 arpip 来判断返回的值的交换机类型
	// if datas[0].ARPIP != "" {
	// 	var result []*maclist.ARPList
	// 	for _, d := range datas {
	// 		result = append(result, &maclist.ARPList{ARPIP: d.ARPIP, MacAddress: d.MacAddress})
	// 	}
	// 	return ms.rep.SaveARP(result)
	// }

	// var result []*maclist.MacAddrs
	// for _, d := range datas {
	// 	result = append(result, &maclist.MacAddrs{MacAddress: d.MacAddress, Port: d.Port, SwitchIp: d.SwitchIp})
	// }
	// return ms.rep.SaveMac(result)

}

var srv = &MacListService{}

func (ms *MacListService) Name() string {
	return "maclist"
}

func (ms *MacListService) Config() error {
	ms.log = zap.L().Named("maclist")
	// sw.db = conf.C().Sqlite.GetDB()
	// NewSwitchImpl()
	ms.rep = app.GetInternalApp("maclist-impl").(*MacListServiceImpl)
	return nil
}

func init() {
	app.RegistryInternalApp(srv)
}
