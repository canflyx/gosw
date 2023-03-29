package impl

import (
	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/conf"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"gorm.io/gorm"
)

var impl = &MacListServiceImpl{}

func NewSwitchImpl() *MacListServiceImpl {
	return &MacListServiceImpl{
		l:  zap.L().Named("Switch"),
		db: conf.C().Sqlite.GetDB(),
	}
}

type MacListServiceImpl struct {
	l  logger.Logger
	db *gorm.DB
}

func (sw *MacListServiceImpl) Name() string {
	return "maclist-impl"
}

func (sw *MacListServiceImpl) Config() error {
	sw.l = zap.L().Named("MacListImpl")
	sw.db = conf.C().Sqlite.GetDB()
	return nil
}

func init() {
	app.RegistryInternalApp(impl)
}
