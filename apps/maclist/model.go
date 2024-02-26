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
type ScanLog struct {
	gorm.Model
	SwitchIP string `json:"switch_ip"`
	Log      string `json:"log"`
}

// MacList Mac列表接口响应数据

type MacList struct {
	MacAddrs
	ARPIP string `json:"arp_ip"`
}

func NewMacList() *MacList {
	return &MacList{}
}

// QueryKwRequest 获取Mac列表query string参数
type QueryKwRequest struct {
	PageSize   int                    `json:"page_size,omitempty"`
	PageNumber int                    `json:"page_number,omitempty"`
	Keyword    map[string]interface{} `json:"kws"`
}

func NewKwRequest() *QueryKwRequest {
	return &QueryKwRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type ListData struct {
	List  []int  `json:"list"`
	Value int    `json:"value"`
	CuCms string `json:"cu_cmds"`
	Flag  int    `json:"flag"`
}

type CMD struct {
	Cmd  string `json:"cmd"`
	Flag string `json:"flag"`
}

func NewQueryMacFromHttp(r *http.Request) *QueryKwRequest {
	req := NewKwRequest()
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

// 查询的返回数据
type LogSet struct {
	Total uint64     `json:"total"`
	Items []*LogList `json:"items"`
}

type LogList struct {
	Switch_IP  string `json:"switch_ip"`
	Log        string `json:"log"`
	Updated_At string `json:"UpdatedAt"`
}

func NewMacSet() *MacSet {
	return &MacSet{}
}

func (req *QueryKwRequest) GetPageSize() int {
	return int(req.PageSize)
}
func (req *QueryKwRequest) OffSet() int {
	return int(req.PageNumber)
}
