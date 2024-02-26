package api

import (
	"errors"

	"github.com/canflyx/gosw/apps/maclist"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

// @Summary 扫描选中交换机
// @Tags 交换机ID数组
// @Accept application/json
// @Produce application/json
// @Param list{[]int} body  true
// @Success 0 {object}
func (h *Handler) scanSw(c *gin.Context) {
	var ins maclist.ListData
	if err := c.ShouldBindJSON(&ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	if ins.Flag == 2 && len(ins.ReadCmd) < 1 {
		response.Failed(c.Writer, errors.New("read cmd nil"))
		return
	}
	err := h.svc.ScanSwitch(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, ins)
}

// @Summary 查询已扫描后的结果
// @Tags 查询 map
// @Accept application/json
// @Produce application/json
// @Param object body {page_number: page_size: kws:{'field':'value'}} false "每页显示多少行，默认为20行"

func (h *Handler) queryMacList(c *gin.Context) {
	ins := maclist.NewKwRequest()
	if err := c.ShouldBindJSON(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	set, err := h.svc.QueryMacList(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}

// @Summary 查询扫描后log
// @Tags 查询 map
// @Accept application/json
// @Produce application/json
// @Param object body {page_number: page_size: kws:{'field':'value'}} false "每页显示多少行，默认为20行"

func (h *Handler) logList(c *gin.Context) {
	ins := maclist.NewKwRequest()
	if err := c.ShouldBindJSON(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	set, err := h.svc.QueryLogList(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}
