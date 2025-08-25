package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

type routerData struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateServer(logger *log.Logger) (*routerData, error) {
	registerHandlers()
	routerData := newRouterData(logger)
	return routerData, nil
}

func registerHandlers() {
	http.Handle("/", http.FileServer(http.Dir("web")))
}

func newRouterData(logger *log.Logger) *routerData {
	todoPort := ":" + os.Getenv("TODO_PORT")
	if todoPort == ":" {
		todoPort = ":7540"
	}

	server := &http.Server{
		Addr:         todoPort,
		Handler:      http.DefaultServeMux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
	}

	router := &routerData{
		Logger: logger,
		Server: server,
	}

	return router
}
