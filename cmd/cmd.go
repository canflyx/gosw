package cmd

import (
	"fmt"
	"log/slog"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/tools"
	"github.com/canflyx/gosw/conf"
	"github.com/canflyx/gosw/protocol"
)

func Start() {
	err := conf.LoadConfigFromYaml(confFile)
	if err != nil {
		fmt.Println(err)
	}
	conf.Zlog = tools.SlogInit()
	conf.ScanPool = 0
	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		fmt.Println(err)
	}

	svc := newManager()
	svc.Start()
}

// 引用protocol 中的服务
type manager struct {
	http *protocol.HttpService
	l    *slog.Logger
}

func newManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    conf.GetNameLog("manager"),
	}
}
func (m *manager) Start() error {
	return m.http.Start()
}
