package repositories

import (
	"log"

	"github.com/afriedrichsen/go-tasker/server/models"
	"github.com/gocraft/work"
)

// ScheduledTaskRepo implements models.ScheduledTask
type ScheduledTaskRepo struct {
	enqueuer *work.Enqueuer
}

func NewScheduledTaskRepo(enqueuer *work.Enqueuer) *ScheduledTaskRepo {
	return &ScheduledTaskRepo{
		enqueuer: enqueuer,
	}
}

// FindByID ..
func (r *ScheduledTaskRepo) FindByID(ID string) (*models.ScheduledTask, error) {
	return &models.ScheduledTask{}, nil
}

// Save ..
func (r *ScheduledTaskRepo) Save(user *models.ScheduledTask) error {
	// Enqueue a job named "send_email" with the specified parameters.
	_, err := r.enqueuer.Enqueue("send_email", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
