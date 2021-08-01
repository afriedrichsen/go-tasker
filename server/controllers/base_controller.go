package controllers

import (
	"fmt"
	"net/http"

	"github.com/afriedrichsen/go-tasker/server/models"
	"github.com/gorilla/mux"
)

// BaseHandler will hold everything we need.
type BaseController struct {
	scheduledtask models.ScheduledTaskRepository
}

// NewBaseController returns a new BaseController
func NewBaseController(taskRepo models.ScheduledTaskRepository) *BaseController {
	return &BaseController{
		scheduledtask: taskRepo,
	}
}

func (c *BaseController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index Page! %v\n", vars["category"])
}

func (c *BaseController) CreateScheduledTask(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	new_task := &models.ScheduledTask{}
	c.scheduledtask.Save(new_task)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task Created: %v\n", "Task Created!")
}

func (c *BaseController) GetScheduledTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func (c *BaseController) GetScheduledTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func (c *BaseController) UpdateScheduledTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func (c *BaseController) DeleteScheduledTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
