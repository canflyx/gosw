package app

import (
	"fmt"

	"github.com/canflyx/gosw/apps/switches"
)

var (
	internalApps = map[string]InternalApp{}
	HostService  switches.Service
)

// InternalApp 内部服务实例, 不需要暴露
type InternalApp interface {
	Config() error
	Name() string
}

// RegistryInternalApp 服务实例注册
func RegistryInternalApp(app InternalApp) {
	// 已经注册的服务禁止再次注册
	_, ok := internalApps[app.Name()]
	if ok {
		panic(fmt.Sprintf("internal app %s has register", app.Name()))
	}

	internalApps[app.Name()] = app
	if v, ok := app.(switches.Service); ok {
		HostService = v
	}
}

// LoadedInternalApp 查询加载成功的服务
func LoadedInternalApp() (apps []string) {
	for k := range internalApps {
		apps = append(apps, k)
	}
	return
}

func GetInternalApp(name string) InternalApp {
	app, ok := internalApps[name]
	if !ok {
		panic(fmt.Sprintf("internal app %s not register", name))
	}

	return app
}
