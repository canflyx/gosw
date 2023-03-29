package protocol

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestHttp(t *testing.T) {
	r := gin.New()
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "OK2"})
	})

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              "0.0.0.0:8000",
		Handler:           engine,
	}
	serv := &HttpService{
		r:      r,
		server: server,
	}
	if err := serv.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("service is stopped")
		}
		fmt.Printf("start service error, %s", err.Error())
	}
}
