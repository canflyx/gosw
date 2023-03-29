package maclist

import (
	"context"

	"github.com/canflyx/gosw/apps/switches"
)

type Service interface {
	ScanSwitch(context.Context, []int) error
	QueryMacList(ctx context.Context, req *QueryMacRequest) (*MacSet, error)
	SaveAll(ctx context.Context, sw *switches.Switches) error
	TelnetSwitch(sw *switches.Switches) ([]*MacList, error)
}

type Repositoryer interface {
	SaveMac([]*MacAddrs) error
	SaveARP([]*ARPList) error
	QueryByKws(map[string]interface{}, int, int) (uint64, []*MacList)
	DescBySWIP(kws map[string]interface{}) []*MacAddrs
}
