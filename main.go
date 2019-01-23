package main

import (
	"fmt"
	"go-gin-server/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", common.HTTPPort),
		Handler:        router,
		ReadTimeout:    common.ReadTimeout,
		WriteTimeout:   common.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
