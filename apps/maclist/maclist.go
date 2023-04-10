package maclist

import (
	"context"

	"github.com/canflyx/gosw/apps/switches"
)

type Service interface {
	ScanSwitch(context.Context, ListData) error
	QueryMacList(ctx context.Context, req *QueryKwRequest) (*MacSet, error)
	QueryLogList(ctx context.Context, req *QueryKwRequest) (*LogSet, error)
	SaveAll(ctx context.Context, sw *switches.Switches, cmd []CMD) error
	TelnetSwitch(sw *switches.Switches) ([]*MacList, error)
}

type Repositoryer interface {
	SaveMac([]*MacAddrs) error
	SaveARP([]*ARPList) error
	SaveLog(*ScanLog) error
	QueryByKws(map[string]interface{}, int, int) (uint64, []*MacList)
	DescBySWIP(kws map[string]interface{}) []*MacAddrs
	QueryLogByKws(map[string]interface{}, int, int) (uint64, []*LogList)
}
