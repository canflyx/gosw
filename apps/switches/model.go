package switches

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

type Switches struct {
	gorm.Model
	Ip       string `json:"ip"  binding:"required"`
	User     string `json:"user" validator:"required" binding:"required"`
	Brand    string `json:"brand" `
	Password string `json:"password" binding:"required"`
	IsCore   *int   `gorm:"force" json:"iscore" validator:"required" binding:"required"`
	SwType   *int   `gorm:"force" json:"swtype" validator:"required" binding:"required"`
	Status   *int   `gorm:"force" json:"status"`
	Note     string `json:"note"`
}

type MacAddress struct {
	gorm.Model
	Mac      string `json:"mac_address"`
	Port     string `json:"port"`
	SwitchIp string `json:"switch_ip"`
}

type ARPList struct {
	gorm.Model
	ARPIP      string `json:"arp_ip"`
	MacAddress string `json:"mac_address"`
}

type QuerySwitchRequest struct {
	PageSize   int                    `json:"page_size,omitempty"`
	PageNumber int                    `json:"page_number,omitempty"`
	Keyword    map[string]interface{} `json:"kws"`
}

func NewSwitchRequest() *QuerySwitchRequest {
	return &QuerySwitchRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type QuerySwitchRequestById struct {
	Id uint `json:"id"`
}

func NewSwitchRequestById(id uint) *QuerySwitchRequestById {
	return &QuerySwitchRequestById{
		Id: id,
	}
}

// QueryKwRequest 获取Mac列表query string参数

// func NewQuerySwitchFromHttp(r *http.Request) *QuerySwitchRequest {
// 	req := NewSwitchRequest()
// 	qs := r.URL.Query()
// 	pss := qs.Get("page_size")
// 	if pss != "" {
// 		size, _ := strconv.Atoi(pss)
// 		req.PageSize = int(size)
// 	}
// 	pns := qs.Get("page_number")
// 	if pns != "" {
// 		number, _ := strconv.Atoi(pns)
// 		req.PageNumber = int(number)
// 	}
// 	req.Keyword = qs.Get("kws")
// 	return req
// }

func NewSwitch() *Switches {
	return &Switches{}
}

// 查询的返回数据
type SwitchesSet struct {
	Total int64       `json:"total"`
	Items []*Switches `json:"items"`
}

func NewSwitchSet() *SwitchesSet {
	return &SwitchesSet{}
}

func (s *Switches) Validate() error {
	return validate.Struct(s)
}

func (req *QuerySwitchRequest) GetPageSize() int {
	return int(req.PageSize)
}
func (req *QuerySwitchRequest) OffSet() int {
	return int(req.PageNumber)
}
