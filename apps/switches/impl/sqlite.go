package impl

import (
	"log/slog"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/conf"
	"gorm.io/gorm"
)

var impl = &SwitchesServiceImpl{}

func NewSwitchImpl() *SwitchesServiceImpl {
	return &SwitchesServiceImpl{
		l:  conf.GetNameLog("Switch"),
		db: conf.C().Sqlite.GetDB(),
	}
}

type SwitchesServiceImpl struct {
	l  *slog.Logger
	db *gorm.DB
}

func (sw *SwitchesServiceImpl) Name() string {
	return "switches-impl"
}

func (sw *SwitchesServiceImpl) Config() error {
	sw.l = conf.GetNameLog("SwitchImpl")
	sw.db = conf.C().Sqlite.GetDB()
	return nil
}

func init() {
	app.RegistryInternalApp(impl)
}
