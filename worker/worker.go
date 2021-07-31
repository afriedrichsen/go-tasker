package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
)

type Worker struct {
	Router *mux.Router
}

type Context struct {
	taskID string
}

func (w *Worker) Initialize() {
	// Make a redis pool
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "redis-cart.default.svc.cluster.local:6379")
		},
	}
	pool := work.NewWorkerPool(Context{}, 10, "go_tasker", redisPool)
	pool.Middleware((*Context).Log)
	fmt.Println("Go Tasker - Worker READY!")
}

func (w *Worker) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8020", w.Router))
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (c *Context) FindTask(job *work.Job, next work.NextMiddlewareFunc) error {
	// If there's a task_id param, set it in the context for future middleware and handlers to use.
	if _, ok := job.Args["task_id"]; ok {
		c.taskID = job.ArgString("task_id")
		if err := job.ArgError(); err != nil {
			return err
		}
	}

	return next()
}

func (c *Context) SendCommand(job *work.Job) error {
	// Extract arguments:
	// addr := job.ArgString("address")
	// subject := job.ArgString("subject")
	if err := job.ArgError(); err != nil {
		return err
	}

	// Go ahead and send the email...
	// sendEmailTo(addr, subject)

	return nil
}

// func (c *Context) Export(job *work.Job) error {
// 	return nil
// }
