package api

import (
	"strconv"

	"github.com/canflyx/gosw/apps/switches"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

// @Summary 创建交换机
// @Tags switches
// @Accept application/json
// @Produce application/json
// @Param object body {ip: user: brand: password: iscore: status: note:} false "每页显示多少行，默认为20行"

func (h *Handler) createSw(c *gin.Context) {
	ins := switches.NewSwitch()
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	ins, err := h.svc.CreateSwitch(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, ins)
}

func (h *Handler) querySw(c *gin.Context) {
	ins := switches.NewSwitchRequest()
	_ = c.ShouldBindJSON(ins)
	set, err := h.svc.QuerySwitchFromHttp(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}

func (h *Handler) queryDesc(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	set, err := h.svc.DescribeHost(c.Request.Context(), switches.NewSwitchRequestById(uint(id)))
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}

func (h *Handler) updateSw(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	ins := switches.NewSwitch()
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	ins.ID = uint(id)
	set, err := h.svc.UpdateSwitch(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, set)
}

func (h *Handler) deleteSw(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	if err := h.svc.DeleteSwitch(c.Request.Context(), id); err != nil {
		response.Failed(c.Writer, err)
		return
	}
}
