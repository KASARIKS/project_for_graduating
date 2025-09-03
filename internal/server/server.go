package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kasariks/project_for_graduating/internal/handlers"
	"github.com/kasariks/project_for_graduating/internal/handlers/taskhandlers"
)

type routerData struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateServer(logger *log.Logger) (*routerData, error) {
	initHandlers()
	routerData := newRouterData(logger)
	return routerData, nil
}

func initHandlers() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/api/nextdate", handlers.GetNextDate)
	http.HandleFunc("/api/task", handlers.Auth(taskhandlers.Task))
	http.HandleFunc("/api/tasks", taskhandlers.GetTasks)
	http.HandleFunc("/api/task/done", handlers.Auth(taskhandlers.TaskDone))
	http.HandleFunc("/api/signin", handlers.SignIn)
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
