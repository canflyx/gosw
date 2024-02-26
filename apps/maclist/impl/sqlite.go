package impl

import (
	"log/slog"

	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/conf"
	"gorm.io/gorm"
)

var impl = &MacListServiceImpl{}

func NewSwitchImpl() *MacListServiceImpl {
	return &MacListServiceImpl{
		l:  conf.GetNameLog("Switch"),
		db: conf.C().Sqlite.GetDB(),
	}
}

type MacListServiceImpl struct {
	l  *slog.Logger
	db *gorm.DB
}

func (sw *MacListServiceImpl) Name() string {
	return "maclist-impl"
}

func (sw *MacListServiceImpl) Config() error {
	sw.l = conf.GetNameLog("MacListImpl")
	sw.db = conf.C().Sqlite.GetDB()
	return nil
}

func init() {
	app.RegistryInternalApp(impl)
}
