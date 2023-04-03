package impl

import (
	"fmt"
	"testing"

	"github.com/canflyx/gosw/apps/maclist"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestQuery(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../../../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	impl := &MacListServiceImpl{
		db: db,
	}
	i, b := impl.QueryByKws(map[string]interface{}{"arp_ip": "172.17"}, 1, 20)
	fmt.Println(i, b)
}
func TestSaveMac(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	impl := &MacListServiceImpl{
		db: db,
	}
	var macs []*maclist.MacAddrs
	mac := &maclist.MacAddrs{
		MacAddress: "e050-8be1-bb15",
		Port:       "Eth0/0/3",
		SwitchIp:   "172.17.3.11",
	}
	macs = append(macs, mac)

	err := impl.SaveMac(macs)
	fmt.Println(err)
}
func TestSaveARP(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	impl := &MacListServiceImpl{
		db: db,
	}
	var macs []*maclist.ARPList
	mac := &maclist.ARPList{
		MacAddress: "e050-8be1-bb15",
		ARPIP:      "192.168.1.1",
	}
	macs = append(macs, mac)

	err := impl.SaveARP(macs)
	fmt.Println(err)
}

func TestSaveLog(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("../../../app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if !db.Migrator().HasTable(&maclist.ScanLog{}) {
		db.AutoMigrate(&maclist.ScanLog{})
	}
	impl := &MacListServiceImpl{
		db: db,
	}

	mac := &maclist.ScanLog{
		SwitchID: 11,
		Log:      "192.168.1.1",
	}

	err := impl.SaveLog(mac)
	fmt.Println(err)
}
