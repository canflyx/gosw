package impl

import (
	"context"
	"fmt"
	"testing"

	"github.com/canflyx/gosw/apps/switches"
)

type TestRep struct{}

func (t *TestRep) Save(ctx context.Context, sws []*switches.Switches) error {
	for _, v := range sws {
		fmt.Println("insert db:", v)
	}
	return nil
}
func (t *TestRep) QueryByKws(map[string]interface{}, int, int) (int64, []*switches.Switches) {
	return 0, nil
}
func (t *TestRep) DescById(uint) *switches.Switches {
	return nil
}
func (t *TestRep) Update(context.Context, *switches.Switches) error {
	return nil
}
func (t *TestRep) Delete(context.Context, int) error {
	return nil
}

func TestCreate(t *testing.T) {
	core := 1
	sws := switches.NewSwitch()
	sws.Ip = "1.1.1.10"
	sws.User = "admin"
	sws.Password = "admin"
	sws.IsCore = &core
	a := &SwitchService{
		rep: &TestRep{},
	}
	c, err := a.CreateSwitch(context.Background(), sws)
	fmt.Printf("%v,%v", c, err)
}
