package impl

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/switches"
	"github.com/canflyx/gosw/apps/tools"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

var _ switches.Service = (*SwitchService)(nil)

type SwitchService struct {
	log logger.Logger
	rep switches.Repositoryer
}

func (sw *SwitchService) CreateSwitch(ctx context.Context, ins *switches.Switches) (*switches.Switches, error) {
	var sws []*switches.Switches
	password, err := tools.EncryptByAes(ins.Password)
	if err != nil {
		return nil, errors.New("encrypt password error: ")
	}
	if strings.Contains(ins.Ip, ";") {
		ips := strings.Split(ins.Ip, ";")
		for _, ip := range ips {
			if net.ParseIP(ip) == nil {
				continue
			}
			sws = append(sws, &switches.Switches{
				Ip:       ip,
				User:     ins.User,
				Password: password,
				IsCore:   ins.IsCore,
				Brand:    ins.Brand,
			})
		}
	} else if strings.Contains(ins.Ip, "-") {
		ips := strings.Split(ins.Ip, "-")
		if net.ParseIP(ips[0]) == nil {
			return nil, errors.New("ip input error: " + ips[0])
		}

		ipStart, err := strconv.Atoi(strings.Split(ips[0], ".")[3])
		if err != nil {
			return nil, errors.New("ip input error: " + ips[0])
		}
		ipEnd, err := strconv.Atoi(ips[1])
		if err != nil {
			return nil, errors.New("ip input error: " + ips[0])
		}
		Ip := ips[0][:strings.LastIndex(ips[0], ".")]
		for i := ipStart; i < ipEnd; i++ {
			newIp := fmt.Sprintf("%s.%d", Ip, i)
			if net.ParseIP(newIp) == nil {
				continue
			}
			sws = append(sws, &switches.Switches{
				Ip:       newIp,
				User:     ins.User,
				Password: password,
				IsCore:   ins.IsCore,
				Brand:    ins.Brand,
			})
		}

	} else {
		if err := ins.Validate(); err != nil {
			return nil, err
		}
		ins.Password = password
		sws = append(sws, ins)
	}
	if err := sw.rep.Save(ctx, sws); err != nil {
		return nil, err
	}
	return ins, nil
}

func (sw *SwitchService) QuerySwitchFromHttp(ctx context.Context, req *switches.QuerySwitchRequest) (*switches.SwitchesSet, error) {
	total, item := sw.rep.QueryByKws(req.Keyword, req.OffSet(), req.GetPageSize())
	for i := 0; i < len(item); i++ {
		item[i].Password = "******"
	}
	return &switches.SwitchesSet{
		Total: total,
		Items: item,
	}, nil
}

func (sw *SwitchService) DescribeHost(ctx context.Context, req *switches.QuerySwitchRequestById) (*switches.Switches, error) {
	if req.Id < 1 {
		return nil, errors.New("id is nil")
	}
	sws := sw.rep.DescById(req.Id)
	return sws, nil
}

func (sw *SwitchService) UpdateSwitch(ctx context.Context, ins *switches.Switches) (*switches.Switches, error) {
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	if ins.ID < 1 {
		return nil, errors.New("id is nil")
	}
	if ins.Password == "******" {
		ins.Password = ""
	}
	if ins.Password != "" {
		ins.Password, _ = tools.DecryptByAes(ins.Password)
	}
	newSwitches :=
		&switches.Switches{
			Model: gorm.Model{
				ID: ins.ID,
			},
			User:     ins.User,
			Password: ins.Password,
			IsCore:   ins.IsCore,
			Brand:    ins.Brand,
			Status:   ins.Status,
			Note:     ins.Note,
		}
	if err := sw.rep.Update(ctx, newSwitches); err != nil {
		return nil, err
	}
	return ins, nil
}

func (sw *SwitchService) DeleteSwitch(ctx context.Context, id int) error {
	if id < 1 {
		return errors.New("id is nil")
	}
	return sw.rep.Delete(ctx, id)
}

var srv = &SwitchService{}

func (sw *SwitchService) Name() string {
	return "switches"
}

func (sw *SwitchService) Config() error {
	sw.log = zap.L().Named("Switches")
	// sw.db = conf.C().Sqlite.GetDB()
	// NewSwitchImpl()
	sw.rep = app.GetInternalApp("switches-impl").(*SwitchesServiceImpl)
	return nil
}

func init() {
	app.RegistryInternalApp(srv)
}
