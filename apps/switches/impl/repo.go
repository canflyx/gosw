package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/canflyx/gosw/apps/switches"
	"gorm.io/gorm/clause"
)

var _ switches.Repositoryer = (*SwitchesServiceImpl)(nil)

func (sw *SwitchesServiceImpl) Save(ctx context.Context, ins []*switches.Switches) error {
	var sws *switches.Switches
	Db := sw.db
	for i, ins := range ins {
		if i == 0 {
			Db = Db.Where("ip=?", ins.Ip)
			continue
		}
		Db = Db.Or("ip=?", ins.Ip)
	}
	if Db.Find(&sws).RowsAffected > 0 {
		return errors.New("ip address already exists")
	}
	return sw.db.Create(&ins).Error
}

func (sw *SwitchesServiceImpl) QueryByKws(kws map[string]interface{}, offset, pagesize int) (int64, []*switches.Switches) {
	var sws []*switches.Switches
	Db := sw.db

	if v, ok := kws["ip"]; ok {
		if v != nil || v != "" {
			Db = Db.Where(fmt.Sprintf("ip LIKE '%%%s%%'", v))
		}
	}
	var count int64
	Db.Model(&sws).Count(&count)
	Db = Db.Order("ip").Offset(offset - 1).Limit(pagesize)
	Db.Find(&sws)
	return count, sws
}

func (sw *SwitchesServiceImpl) DescById(id uint) *switches.Switches {
	var sws *switches.Switches
	sw.db.First(&sws, id)
	return sws
}

func (sw *SwitchesServiceImpl) Update(ctx context.Context, ins *switches.Switches) error {
	return sw.db.Model(&ins).Updates(&ins).Error
}

func (sw *SwitchesServiceImpl) Delete(ctx context.Context, id int) error {
	var sws switches.Switches
	result := sw.db.Unscoped().Delete(&sws, id)
	if result.RowsAffected < 1 {
		return errors.New("delete id non-existent")
	}
	return result.Error

}

func (sw *SwitchesServiceImpl) SaveMac(ctx context.Context, macs []*switches.MacAddress) error {

	return sw.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "mac"}},
		DoUpdates: clause.AssignmentColumns([]string{"port"}),
	}).Create(&macs).Error
}

func (sw *SwitchesServiceImpl) SaveArp(ctx context.Context, arpList []*switches.ARPList) error {

	return sw.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "arp_ip"}},
		DoUpdates: clause.AssignmentColumns([]string{"mac_address"}),
	}).Create(&arpList).Error
}
