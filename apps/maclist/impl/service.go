package impl

import (
	"context"
	"errors"
	"log/slog"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/maclist"
	"github.com/canflyx/gosw/apps/switches"
	swimpl "github.com/canflyx/gosw/apps/switches/impl"
	"github.com/canflyx/gosw/apps/tools"
	"github.com/canflyx/gosw/conf"
)

var _ maclist.Service = (*MacListService)(nil)

type MacListService struct {
	log *slog.Logger
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
		conf.ScanPool++
		sws.Password, _ = tools.DecryptByAes(sws.Password)
		go func() {
			_ = ms.SaveAll(ctx, sws, ins.CuCms)
			conf.ScanPool--
		}()
	}
	return nil
}

// 查询数据给 gin 使用
func (ms *MacListService) QueryMacList(ctx context.Context, req *maclist.QueryKwRequest) (*maclist.MacSet, error) {
	total, item := ms.rep.QueryByKws(req.Keyword, req.OffSet(), req.GetPageSize())
	return &maclist.MacSet{
		Total: total,
		Items: item,
	}, nil
}

func (ms *MacListService) QueryLogList(ctx context.Context, req *maclist.QueryKwRequest) (*maclist.LogSet, error) {
	total, item := ms.rep.QueryLogByKws(req.Keyword, req.OffSet(), req.GetPageSize())
	return &maclist.LogSet{
		Total: total,
		Items: item,
	}, nil
}

// 传入交换机进行 telnet 返回数据依据类型进行保存
func (ms *MacListService) SaveAll(ctx context.Context, sw *switches.Switches, value string) error {
	cmd, err := GetBrandCmd(sw.Brand)
	if err != nil {
		return err
	}
	if len(value) > 0 {
		// 自定义为 PreCmd
		cmd.UserCmd = tools.FormatCmd(value)

		err, _ := NewCuTelnet(&SwitchesConfig{Switches: *sw, BrandCMD: cmd, Flag: 2, TimeOut: 5})
		if err != nil {
			ms.log.Error("switch:"+sw.Ip, "error:", err)
			return errors.New("switch telnet error")
		}
		return nil
	}
	if *sw.IsCore == 1 {
		err, _ := NewARPTelnet(&SwitchesConfig{Switches: *sw, BrandCMD: cmd, Flag: 1, TimeOut: 5})
		if err != nil {
			ms.log.Error("switch:"+sw.Ip, "error:", err)
			return errors.New("switch telnet error")
		}
		return nil
	}
	_, err = NewMacTelnet(&SwitchesConfig{Switches: *sw, BrandCMD: cmd, Flag: 0, TimeOut: 5})
	if err != nil {
		ms.log.Error("switch:"+sw.Ip, "error:", err)
		return errors.New("switch telnet error")
	}
	return nil

}

func GetBrandCmd(brand string) (conf.TelnetCmd, error) {
	cmds := conf.C().TelnetCmd()
	var cmd conf.TelnetCmd
	v, ok := cmds[brand]
	if !ok {
		v1, ok1 := cmds["default"]
		if !ok1 {
			return cmd, errors.New("config not brand or default ")
		}
		cmd = v1
	} else {
		cmd = v
	}
	return cmd, nil
}

var srv = &MacListService{}

func (ms *MacListService) Name() string {
	return "maclist"
}

func (ms *MacListService) Config() error {
	ms.log = conf.GetNameLog("maclist")
	// sw.db = conf.C().Sqlite.GetDB()
	// NewSwitchImpl()
	ms.rep = app.GetInternalApp("maclist-impl").(*MacListServiceImpl)
	return nil
}

func init() {
	app.RegistryInternalApp(srv)
}
