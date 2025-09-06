package task

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/muhammadputraage/task-manager-api/pkg/utils"
)

type Handler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repository Repository
}

func NewHandler(repository Repository) Handler {
	return &handler{repository: repository}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repository.FindAll()
	if err != nil {
		utils.APIResponse(w, http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to fetch tasks",
			Error:   err.Error(),
		})
		return
	}

	var tasksResponse []TaskResponse
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     task.DueDate,
			Completed:   task.Completed,
		})
	}

	utils.APIResponse(w, http.StatusOK, utils.Response{
		Success: true,
		Message: "Tasks fetched successfully",
		Data:    tasksResponse,
	})
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var request CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
		return
	}

	if validationErr := utils.ValidateRequest(request); len(validationErr) > 0 {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Validation failed",
			Error:   validationErr,
		})
		return
	}

	dueDate, err := time.Parse("02/01/2006", request.DueDate)
	if err != nil {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Validation failed",
			Error: map[string]string{
				"dueDate": "Invalid date format",
			},
		})
		return
	}

	newTask := Task{
		Title:       request.Title,
		Description: request.Description,
		DueDate:     dueDate,
		Completed:   false,
	}

	createdTask, err := h.repository.Create(&newTask)
	if err != nil {
		utils.APIResponse(w, http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to create task",
			Error:   err.Error(),
		})
		return
	}

	taskResponse := TaskResponse{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		DueDate:     createdTask.DueDate,
		Completed:   createdTask.Completed,
	}

	utils.APIResponse(w, http.StatusCreated, utils.Response{
		Success: true,
		Message: "Task created successfully",
		Data:    taskResponse,
	})
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid task ID",
			Error:   err.Error(),
		})
		return
	}

	task, err := h.repository.FindById(id)
	if err != nil {
		utils.APIResponse(w, http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Task not found",
			Error:   err.Error(),
		})
		return
	}

	taskResponse := TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Completed:   task.Completed,
	}

	utils.APIResponse(w, http.StatusOK, utils.Response{
		Success: true,
		Message: "Task fetched successfully",
		Data:    taskResponse,
	})
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid task ID",
			Error:   err.Error(),
		})
		return
	}

	var request UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
		return
	}

	if validationErr := utils.ValidateRequest(request); len(validationErr) > 0 {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Validation failed",
			Error:   validationErr,
		})
		return
	}

	dueDate, err := time.Parse("02/01/2006", request.DueDate)
	if err != nil {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Validation failed",
			Error: map[string]string{
				"dueDate": "Invalid date format",
			},
		})
		return
	}

	task, err := h.repository.FindById(id)
	if err != nil {
		utils.APIResponse(w, http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Task not found",
			Error:   err.Error(),
		})
		return
	}

	task.Title = request.Title
	task.Description = request.Description
	task.DueDate = dueDate
	task.Completed = request.Completed

	updatedTask, err := h.repository.Update(&task)
	if err != nil {
		utils.APIResponse(w, http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update task",
			Error:   err.Error(),
		})
		return
	}

	taskResponse := TaskResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		DueDate:     updatedTask.DueDate,
		Completed:   updatedTask.Completed,
	}

	utils.APIResponse(w, http.StatusOK, utils.Response{
		Success: true,
		Message: "Task updated successfully",
		Data:    taskResponse,
	})
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.APIResponse(w, http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid task ID",
			Error:   err.Error(),
		})
		return
	}

	task, err := h.repository.FindById(id)
	if err != nil {
		utils.APIResponse(w, http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Task not found",
			Error:   err.Error(),
		})
		return
	}

	if err := h.repository.Delete(&task); err != nil {
		utils.APIResponse(w, http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to delete task",
			Error:   err.Error(),
		})
		return
	}

	utils.APIResponse(w, http.StatusOK, utils.Response{
		Success: true,
		Message: "Task deleted successfully",
	})
}
