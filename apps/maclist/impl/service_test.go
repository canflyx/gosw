package impl

import (
	"fmt"

	"github.com/canflyx/gosw/apps/maclist"
	"github.com/canflyx/gosw/conf"
	"github.com/infraboard/mcube/logger/zap"
)

// func TestQueryMacList(t *testing.T) {
// 	re := maclist.QueryKwRequest{PageSize: 10, PageNumber: 1, Keyword: nil}
// 	ctl := gomock.NewController(t)
// 	c := []*maclist.MacList{
// 		{MacAddrs: maclist.MacAddrs{MacAddress: "0000-0000-0000", Port: "ETH0", SwitchIp: "192.168.1.1"}, ARPIP: "192,168.3.2"},
// 		{MacAddrs: maclist.MacAddrs{MacAddress: "0000-0000-0001", Port: "ETH1", SwitchIp: "192.168.1.1"}, ARPIP: "192,168.3.1"},
// 	}
// 	mockPerson := mock_dal.NewMockRepositoryer(ctl)
// 	mockPerson.EXPECT().QueryByKws(gomock.Any(), gomock.Any(), gomock.Any()).Return(uint64(2), c).AnyTimes()
// 	rep := MacListService{rep: mockPerson}

// 	set, _ := rep.QueryMacList(context.Background(), &re)
// 	fmt.Printf("%d,%v  \n", set.Total, set.Items)
// 	for _, mac := range set.Items {
// 		ok, _ := json.Marshal(mac)
// 		fmt.Println(string(ok))
// 	}
// }

type TestRep struct {
}

func (t *TestRep) SaveMac(sws []*maclist.MacAddrs) error {
	fmt.Println(sws)
	return nil
}
func (t *TestRep) SaveLog(log *maclist.ScanLog) error {
	fmt.Println(log)
	return nil
}
func (t *TestRep) SaveARP(sws []*maclist.ARPList) error {
	fmt.Println(sws)
	return nil
}
func (t *TestRep) QueryByKws(kws map[string]interface{}, i int, c int) (uint64, []*maclist.MacList) {
	fmt.Println(kws, i, c)
	return 1, nil
}
func (t *TestRep) QueryLogByKws(kws map[string]interface{}, i int, c int) (uint64, []*maclist.LogList) {
	fmt.Println(kws, i, c)
	return 1, nil
}
func (t *TestRep) DescBySWIP(kws map[string]interface{}) []*maclist.MacAddrs {
	fmt.Println(kws)
	return nil
}

// func TestSaveAll(t *testing.T) {
// 	core := 1
// 	err := conf.LoadConfigFromYaml("config.yaml")
// 	if err := loadGlobalLogger(); err != nil {
// 		fmt.Println(err)
// 	}
// 	sws := &switches.Switches{
// 		Ip:       "172.17.80.1",
// 		User:     "dai",
// 		Password: "dai2018",
// 		IsCore:   &core,
// 	}
// 	a := &MacListService{
// 		rep: &TestRep{},
// 		log: conf.GetNameLog("maclist"),
// 	}
// 	c := []conf.CMD{{CMD: "dis users", CMDFlag: "]"}, {CMD: "dis ver", CMDFlag: "]"}}
// 	err = a.SaveAll(context.Background(), sws, c)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	lc := conf.C().Log
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s,use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level :%s", lv)
	}
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level
	zapConfig.Files.RotateOnStartup = false
	switch lc.To {
	case conf.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	conf.GetNameLog("INIT").Info(logInitMsg)
	return nil
}
