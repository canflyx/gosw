package maclist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

type MacAddrs struct {
	gorm.Model
	MacAddress string `json:"mac_address" gorm:"unique"`
	Port       string `json:"port"`
	SwitchIp   string `json:"switch_ip"`
	Note       string `json:"note"`
}

type ARPList struct {
	gorm.Model
	ARPIP      string `json:"arp_ip"`
	MacAddress string `json:"mac_address" gorm:"unique"`
	Note       string `json:"note"`
}

// MacList Mac列表接口响应数据

type MacList struct {
	MacAddrs
	ARPIP string `json:"arp_ip"`
}

func NewMacList() *MacList {
	return &MacList{}
}

// QueryMacRequest 获取Mac列表query string参数
type QueryMacRequest struct {
	PageSize   int                    `json:"page_size,omitempty"`
	PageNumber int                    `json:"page_number,omitempty"`
	Keyword    map[string]interface{} `json:"kws"`
}

func NewMacRequest() *QueryMacRequest {
	return &QueryMacRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type ListData struct {
	List []int `json:"list"`
}

func NewQueryMacFromHttp(r *http.Request) *QueryMacRequest {
	req := NewMacRequest()
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		size, _ := strconv.Atoi(pss)
		req.PageSize = int(size)
	}
	pns := qs.Get("page_number")
	if pns != "" {
		number, _ := strconv.Atoi(pns)
		req.PageNumber = int(number)
	}
	fmt.Println(qs.Get("kws"))
	err := json.Unmarshal([]byte(qs.Get("kws")), &req.Keyword)
	fmt.Println(req.Keyword)
	if err != nil {
		req.Keyword = nil
	}
	return req
}

// 查询的返回数据
type MacSet struct {
	Total uint64     `json:"total"`
	Items []*MacList `json:"items"`
}

func NewMacSet() *MacSet {
	return &MacSet{}
}

func (req *QueryMacRequest) GetPageSize() int {
	return int(req.PageSize)
}
func (req *QueryMacRequest) OffSet() int {
	return int(req.PageNumber)
}
