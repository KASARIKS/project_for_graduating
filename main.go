package main

import (
	"log"
	"net/http"

	"github.com/kasariks/project_for_graduating/internal/db"
	"github.com/kasariks/project_for_graduating/internal/server"
)

func main() {
	logger := &log.Logger{}
	router, err := server.CreateServer(logger)
	if err != nil {
		log.Fatalf("Error with creating the server: %s\n", err.Error())
	}
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatalf("Error with creating the db: %s\n", err.Error())
	}
	if err := http.ListenAndServe(router.Server.Addr, router.Server.Handler); err != nil {
		log.Fatalf("Error with staring the server: %s\n", err.Error())
	}
}
