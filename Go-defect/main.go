package main

import (
	"Demo/routers"
	"fmt"
	"net/http"

	"Demo/pkg/setting"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 2 << 20,
	}

	s.ListenAndServe()
}
