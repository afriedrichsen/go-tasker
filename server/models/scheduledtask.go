package models

// ScheduledTask
type ScheduledTask struct{}

// ScheduledTaskRepository ..
type ScheduledTaskRepository interface {
	FindByID(ID int) (*ScheduledTask, error)
	Save(task *ScheduledTask) error
}
