package api

import (
	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/switches"
	"github.com/gin-gonic/gin"
)

func NewHostHttpHandler() *Handler {
	return &Handler{}
}

// 写一个实例类，把内部的接口通过 http 协议 暴露出来
// 所以需要依赖内部接口的实现
// 该实体类，会实现 Gin 的 http Handler

type Handler struct {
	svc switches.Service
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/", h.createSw)
	r.POST("/list", h.querySw)
	r.GET(":id", h.queryDesc)
	r.PATCH(":id", h.updateSw)
	r.DELETE("/:id", h.deleteSw)
}

func (h *Handler) Config() error {
	h.svc = app.GetInternalApp(switches.AppName).(switches.Service)
	return nil
}
func (h *Handler) Name() string {
	return switches.AppName
}
func (h *Handler) Version() string {
	return "v1"
}

var handler = &Handler{}

func init() {
	app.RegistryGinApp(handler)
}
