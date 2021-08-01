package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"

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

	// Map the name of jobs to handler functions
	pool.Job("send_command", (*Context).SendCommand)

	// Customize options:
	// pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func (w *Worker) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8020", w.Router))
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
	// return nil
}

// func (c *Context) FindTask(job *work.Job, next work.NextMiddlewareFunc) error {
// 	// If there's a task_id param, set it in the context for future middleware and handlers to use.
// 	if _, ok := job.Args["task_id"]; ok {
// 		c.taskID = job.ArgString("task_id")
// 		if err := job.ArgError(); err != nil {
// 			return err
// 		}
// 	}

// 	return next()
// }

// SendCommand sends our command (or, more generally, runs a command in a child process, but pretend it's running remote)
func (c *Context) SendCommand(job *work.Job) error {
	// Extract arguments:
	command := job.ArgString("command")
	// subject := job.ArgString("subject")
	if err := job.ArgError(); err != nil {
		return err
	}

	// Run command
	cmd := exec.Command(command)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Run()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		return err
	}

	// read command's stdout line by line
	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Printf(in.Text())
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}

	return nil
}

// func (c *Context) Export(job *work.Job) error {
// 	return nil
// }
