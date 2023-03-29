package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_createSw(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	type sw struct {
		name string
		h    *Handler
		args args
	}
	a := &sw{}
	a.h.createSw(a.args.c)

}
