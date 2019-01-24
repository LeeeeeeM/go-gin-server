package main

import (
	"fmt"
	"go-gin-server/common"
	router "go-gin-server/routers"
	"net/http"
)

func main() {
	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", common.HTTPPort),
		Handler:        router,
		ReadTimeout:    common.ReadTimeout,
		WriteTimeout:   common.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
