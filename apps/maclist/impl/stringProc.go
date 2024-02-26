package impl

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/canflyx/gosw/apps/maclist"
)

type ByteCheck struct {
	Sw string
}

// 未处理 mac 未加入交换机IP
func (b *ByteCheck) Process(params interface{}) (interface{}, error) {
	if v, ok := params.([]byte); !ok {
		return errors.New("ByteCheck input type error"), nil
	} else {
		lines := strings.Split(string(v), "\n")
		okSlice := []*maclist.MacList{}
		for _, line := range lines {
			lineArray := strings.Split(line, " ")
			new := maclist.NewMacList()
			for i := 0; i < len(lineArray); i++ {
				if lineArray[i] == "" {
					continue
				}
				if strings.ContainsAny(lineArray[i], "GEXIF") {
					new.Port = lineArray[i]
					continue
				}
				if IsMac(lineArray[i]) {
					new.MacAddress = lineArray[i]
					continue
				}

				if net.ParseIP(lineArray[i]) != nil {
					new.ARPIP = lineArray[i]
					continue
				}
			}

			if new.MacAddress != "" {
				if new.ARPIP == "" {
					new.SwitchIp = b.Sw
				}
				// new.SwitchIp = swIP
				okSlice = append(okSlice, new)
			}
		}
		return okSlice, nil
	}
}

type Dup struct {
}

// 对maclist去重
func (d *Dup) Process(params interface{}) (interface{}, error) {
	if v, ok := params.([]*maclist.MacList); !ok {
		return nil, errors.New("duplicate input type error")
	} else {
		var ret []*maclist.MacList
		var c = make(map[string]int)
		for _, s := range v {

			v, ok := c[s.Port]
			if !ok {
				c[s.Port] = 1
			} else {
				c[s.Port] = v + 1
			}

		}
		for _, s := range v {
			v := c[s.Port]
			if v == 1 {
				ret = append(ret, s)
			}
		}
		return ret, nil
	}
}

type SaveMac struct {
	rep maclist.Repositoryer
}

// 保存交换机 mac-addr
func (s *SaveMac) Process(params interface{}) (interface{}, error) {
	if v, ok := params.([]*maclist.MacList); !ok {
		return errors.New("saveMac input type error"), nil
	} else {
		var result []*maclist.MacAddrs
		for _, d := range v {
			result = append(result, &maclist.MacAddrs{MacAddress: d.MacAddress, Port: d.Port, SwitchIp: d.SwitchIp})
		}
		return s.rep.SaveMac(result), nil
	}

}

type SaveARP struct {
	rep maclist.Repositoryer
}

func (s *SaveARP) Process(params interface{}) (interface{}, error) {
	if v, ok := params.([]*maclist.MacList); !ok {
		return errors.New("saveArp input type error"), nil
	} else {
		var result []*maclist.ARPList
		for _, d := range v {
			result = append(result, &maclist.ARPList{ARPIP: d.ARPIP, MacAddress: d.MacAddress})
		}
		return s.rep.SaveARP(result), nil
	}

}

type SaveLog struct {
	rep maclist.Repositoryer
	sw  SwitchesConfig
}

func (s *SaveLog) Process(param interface{}) (interface{}, error) {
	log := ""
	if v, ok := param.(([]byte)); ok {
		log = string(v)
	} else {
		log = fmt.Sprintf("%v", param)

	}
	return s.rep.SaveLog(&maclist.ScanLog{
		SwitchIP: s.sw.Ip,
		Log:      log,
	}), nil
}

func IsMac(mac string) bool {
	return len(strings.Split(mac, "-")) == 3 || len(strings.Split(mac, ".")) == 3
}
