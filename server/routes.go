package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/afriedrichsen/go-tasker/server/models"
	"github.com/gorilla/mux"
)

// BaseHandler will hold everything we need.
type BaseHandler struct {
	scheduledtask models.ScheduledTaskRepository
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index Page! %v\n", vars["category"])
}

func CreateScheduledTask(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	dispatcher := NewWorkerDispatcher(
		Workers(4),
		JobExpiry(time.Millisecond),
	)
	// Queue a job func
	tracker := dispatcher.QueueFunc(func() error {
		time.Sleep(time.Microsecond)
		return nil
	})

	// Queue a 'JobRunner'
	// dispatcher.Queue(JobRunnerFunc(func() error {
	// 	time.Sleep(time.Microsecond)
	// 	return nil
	// }))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", tracker.ID())
}

func GetScheduledTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func GetScheduledTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func UpdateScheduledTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func DeleteScheduledTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
