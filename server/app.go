package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/afriedrichsen/go-tasker/server/controllers"
	"github.com/afriedrichsen/go-tasker/server/repositories"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {

	// Make a redis pool
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "redis-cart.default.svc.cluster.local:6379")
		},
	}

	// // Make an enqueuer with a particular namespace
	var enqueuer = work.NewEnqueuer("go_tasker", redisPool)

	// Create repos
	taskRepo := repositories.NewScheduledTaskRepo(enqueuer)

	a.Router = mux.NewRouter()

	// tom: this line is added after initializeRoutes is created later on
	a.initializeRoutes(taskRepo)

}

func (a *App) initializeRoutes(repo *repositories.ScheduledTaskRepo) {
	c := controllers.NewBaseController(repo)
	a.Router.HandleFunc("/", c.HomeHandler).Methods("GET")
	a.Router.HandleFunc("/task", c.GetScheduledTask).Methods("GET")
	a.Router.HandleFunc("/task/create", c.CreateScheduledTask).Methods("POST")
	a.Router.HandleFunc("/task/:id", c.GetScheduledTask).Methods("GET")
	a.Router.HandleFunc("/task/update", c.UpdateScheduledTask).Methods("PATCH")
	a.Router.HandleFunc("/task/delete", c.DeleteScheduledTask).Methods("POST", "DELETE")
	fmt.Println("Go Tasker - Routes initialized!")

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}
