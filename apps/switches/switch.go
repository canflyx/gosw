package switches

import (
	"context"
)

type Service interface {
	CreateSwitch(context.Context, *Switches) (*Switches, error)
	QuerySwitchFromHttp(context.Context, *QuerySwitchRequest) (*SwitchesSet, error)
	DescribeHost(context.Context, *QuerySwitchRequestById) (*Switches, error)

	UpdateSwitch(context.Context, *Switches) (*Switches, error)
	DeleteSwitch(context.Context, int) error
}

type Repositoryer interface {
	Save(context.Context, []*Switches) error
	QueryByKws(map[string]interface{}, int, int) (int64, []*Switches)
	DescById(uint) *Switches
	Update(context.Context, *Switches) error
	Delete(context.Context, int) error
}

type Arper interface {
	AddArp([]*ARPList) error
	AddMac([]*MacAddress) error
}
