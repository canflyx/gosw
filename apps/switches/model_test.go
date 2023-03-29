package switches

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewSwitchRequestById(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("./app.db"), &gorm.Config{})
	db.AutoMigrate(&Switches{}, &MacAddress{}, &ARPList{})

}
