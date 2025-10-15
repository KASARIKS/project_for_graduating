package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kasariks/project_for_graduating/internal/db"
	envsettings "github.com/kasariks/project_for_graduating/internal/env_settings"
	"github.com/kasariks/project_for_graduating/internal/server"
)

func main() {
	envsettings.Init()
	logger := &log.Logger{}
	router, err := server.CreateServer(logger)
	if err != nil {
		log.Fatalf("Error with creating the server: %s\n", err.Error())
	}
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatalf("Error with creating the db: %s\n", err.Error())
	}
	defer db.Close()
	fmt.Println("Port:", router.Server.Addr[1:])
	if err := http.ListenAndServe(router.Server.Addr, router.Server.Handler); err != nil {
		log.Fatalf("Error with staring the server: %s\n", err.Error())
	}
}
