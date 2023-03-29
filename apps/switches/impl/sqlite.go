package impl

import (
	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

var impl = &SwitchesServiceImpl{}

func NewSwitchImpl() *SwitchesServiceImpl {
	return &SwitchesServiceImpl{
		l:  zap.L().Named("Switch"),
		db: conf.C().Sqlite.GetDB(),
	}
}

type SwitchesServiceImpl struct {
	l  logger.Logger
	db *gorm.DB
}

func (sw *SwitchesServiceImpl) Name() string {
	return "switches-impl"
}

func (sw *SwitchesServiceImpl) Config() error {
	sw.l = zap.L().Named("SwitchImpl")
	sw.db = conf.C().Sqlite.GetDB()
	return nil
}

func init() {
	app.RegistryInternalApp(impl)
}
