package cmd

import (
	"fmt"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/conf"
	"github.com/canflyx/gosw/protocol"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func Start() {
	err := conf.LoadConfigFromYaml(confFile)
	if err != nil {
		fmt.Println(err)
	}
	if err := loadGlobalLogger(); err != nil {
		fmt.Println(err)
	}
	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		fmt.Println(err)
	}

	svc := newManager()
	svc.Start()
}

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
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

// 引用protocol 中的服务
type manager struct {
	http *protocol.HttpService
	l    logger.Logger
}

func newManager() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    zap.L().Named("manager"),
	}
}
func (m *manager) Start() error {
	return m.http.Start()
}
