package api

import (
	"github.com/canflyx/gosw/app"
	"github.com/canflyx/gosw/apps/maclist"
	"github.com/gin-gonic/gin"
)

func NewHostHttpHandler() *Handler {
	return &Handler{}
}

// 写一个实例类，把内部的接口通过 http 协议 暴露出来
// 所以需要依赖内部接口的实现
// 该实体类，会实现 Gin 的 http Handler

type Handler struct {
	svc maclist.Service
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/scan", h.scanSw)
	r.POST("/list", h.queryMacList)
}

func (h *Handler) Config() error {
	h.svc = app.GetInternalApp(maclist.AppName).(maclist.Service)
	return nil
}
func (h *Handler) Name() string {
	return maclist.AppName
}
func (h *Handler) Version() string {
	return "v1"
}

var handler = &Handler{}

func init() {
	app.RegistryGinApp(handler)
}
