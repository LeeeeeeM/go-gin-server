package main

import (
	"context"
	"fmt"
	"go-gin-server/common"
	router "go-gin-server/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
