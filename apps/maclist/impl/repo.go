package impl

import (
	"fmt"

	"github.com/canflyx/gosw/apps/maclist"
)

var _ maclist.Repositoryer = (*MacListServiceImpl)(nil)

func (ms *MacListServiceImpl) QueryByKws(kws map[string]interface{}, offset, pagesize int) (uint64, []*maclist.MacList) {
	Db := ms.db
	if kws != nil {
		for k, v := range kws {
			if v == nil || v == "" {
				continue
			}
			if k == "arp_ip" {
				Db = Db.Where(fmt.Sprintf("arp_lists.%s  LIKE '%%%s%%' ", k, v.(string)))
				continue
			}
			Db = Db.Where(fmt.Sprintf("mac_addrs.%s  LIKE '%%%s%%' ", k, v.(string)))
		}
	}
	Db = Db.Table("mac_addrs").Select("mac_addrs.mac_address,mac_addrs.port,mac_addrs.switch_ip,arp_lists.arp_ip,mac_addrs.updated_at").Joins("join arp_lists on arp_lists.mac_address = mac_addrs.mac_address")
	var data []*maclist.MacList
	var count int64
	Db.Model(&data).Count(&count)
	// Count := uint64(result.RowsAffected)
	Db = Db.Order("mac_addrs.switch_ip,mac_addrs.port").Offset(offset - 1).Limit(pagesize)
	Db.Find(&data)
	return uint64(count), data
}

func (ms *MacListServiceImpl) QueryLogByKws(kws map[string]interface{}, offset, pagesize int) (uint64, []*maclist.LogList) {
	Db := ms.db
	if kws != nil {
		for k, v := range kws {
			if v == nil || v == "" {
				continue
			}
			if k == "ids" {

				if sw, ok := v.([]int); ok {
					str := ""
					for i := 0; i < len(sw)-1; i++ {
						str = str + fmt.Sprintf("%v,", sw[i])
					}
					str = str + fmt.Sprintf("%v", sw[len(sw)-1])
					Db = Db.Where(fmt.Sprintf("switches.id in (%v)", str))
					continue
				}

			}
			Db = Db.Where(fmt.Sprintf("switches.%s  LIKE '%%%s%%' ", k, v.(string)))
		}
	}
	Db = Db.Table("scan_logs").Select("switches.ip as switch_ip,scan_logs.log,scan_logs.updated_at").Joins("left join switches switches on scan_logs.switch_id = switches.id")
	var data []*maclist.LogList
	var count int64
	Db.Model(&data).Count(&count)
	// Count := uint64(result.RowsAffected)
	Db = Db.Order("scan_logs.updated_at desc").Offset(offset - 1).Limit(pagesize)
	Db.Find(&data)
	return uint64(count), data
}

func (ms *MacListServiceImpl) DescBySWIP(kws map[string]interface{}) []*maclist.MacAddrs {
	var mss []*maclist.MacAddrs
	ms.db.Where(&kws).Find(&mss)
	return mss
}

func (ms *MacListServiceImpl) SaveMac(macs []*maclist.MacAddrs) error {

	return ms.db.Save(&macs).Error
}
func (ms *MacListServiceImpl) SaveLog(log *maclist.ScanLog) error {

	return ms.db.Save(&log).Error
}

func (ms *MacListServiceImpl) SaveARP(arpList []*maclist.ARPList) error {

	return ms.db.Save(&arpList).Error

}
