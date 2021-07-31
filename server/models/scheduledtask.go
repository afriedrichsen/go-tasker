package models

// ScheduledTask
type ScheduledTask struct {
	command string
	url     string
}

// ScheduledTaskRepository ..
type ScheduledTaskRepository interface {
	FindByID(ID string) (*ScheduledTask, error)
	Save(task *ScheduledTask) error
}
