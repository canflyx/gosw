package impl

import (
	"context"
	"fmt"
	"testing"

	"github.com/canflyx/gosw/apps/switches"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestImpl(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	impl := &SwitchesServiceImpl{
		db: db,
	}
	var bb []*switches.Switches
	b := &switches.Switches{
		Ip:       "192.168.1.33",
		User:     "admin",
		Password: "admin",
	}
	bb = append(bb, b)
	err := impl.Save(context.Background(), bb)
	if err != nil {
		fmt.Println(err)
	}
}

func TestQuery(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	impl := &SwitchesServiceImpl{
		db: db,
	}
	i, b := impl.QueryByKws(nil, 1, 20)
	fmt.Println(i, b)
}

func TestDelete(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	impl := &SwitchesServiceImpl{
		db: db,
	}

	err := impl.Delete(context.Background(), 2)
	fmt.Println(err)
}

func TestUpdate(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	impl := &SwitchesServiceImpl{
		db: db,
	}
	core := 0
	b := &switches.Switches{

		Password: "admin123",
		IsCore:   &core,
	}
	b.ID = 3
	err := impl.Update(context.Background(), b)
	fmt.Println(err)
}
func TestSaveMac(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	impl := &SwitchesServiceImpl{
		db: db,
	}
	var macs []*switches.MacAddress
	mac := &switches.MacAddress{
		Mac:      "e050-8be1-bb15",
		Port:     "Eth0/0/3",
		SwitchIp: "172.17.3.11",
	}
	macs = append(macs, mac)

	err := impl.SaveMac(context.Background(), macs)
	fmt.Println(err)
}
