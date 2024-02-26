package switches

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func TestNewSwitchRequestById(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("./app.db"), &gorm.Config{})
	db.AutoMigrate(&Switches{}, &MacAddress{}, &ARPList{})

}
