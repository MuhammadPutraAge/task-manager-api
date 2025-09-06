package task

import "time"

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required"`
	DueDate     string `json:"dueDate" validate:"required"`
}

type UpdateTaskRequest struct {
	CreateTaskRequest
	Completed bool `json:"completed" validate:"boolean"`
}

type TaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Completed   bool      `json:"completed"`
}
