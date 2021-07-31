package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

var redisPool *redis.Pool

func (a *App) Initialize() {

	// Make a redis pool
	redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "redis-cart.default.svc.cluster.local:6379")
		},
	}

	// Make an enqueuer with a particular namespace
	var enqueuer = work.NewEnqueuer("go_tasker", redisPool)

	// Enqueue a job named "send_email" with the specified parameters.
	_, err := enqueuer.Enqueue("send_email", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	// tom: this line is added after initializeRoutes is created later on
	a.initializeRoutes()

	// s.Scheduler.StartAsync()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", HomeHandler)
	a.Router.HandleFunc("/task", GetScheduledTask)
	a.Router.HandleFunc("/task/create", CreateScheduledTask)
	a.Router.HandleFunc("/task/:id", GetScheduledTask)
	a.Router.HandleFunc("/task/update", UpdateScheduledTask)
	a.Router.HandleFunc("/task/delete", DeleteScheduledTask)
	fmt.Println("Go Tasker - Routes initialized!")

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

// func caller() {
// 	svc := sts.New(session.New())
// 	input := &sts.GetCallerIdentityInput{}

// 	result, err := svc.GetCallerIdentity(input)
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			switch aerr.Code() {
// 			default:
// 				fmt.Println(aerr.Error())
// 			}
// 		} else {
// 			// Print the error, cast err to awserr.Error to get the Code and
// 			// Message from an error.
// 			fmt.Println(err.Error())
// 		}
// 		return
// 	}

// 	fmt.Println(result)
// }
