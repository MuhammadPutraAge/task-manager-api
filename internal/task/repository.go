package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Task, error)
	FindById(id int) (Task, error)
	Create(task *Task) (Task, error)
	Update(task *Task) (Task, error)
	Delete(task *Task) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Task, error) {
	tasks := []Task{}
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repository) Create(task *Task) (Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return Task{}, err
	}

	return *task, nil
}

func (r *repository) FindById(id int) (Task, error) {
	var task Task

	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *repository) Update(task *Task) (Task, error) {
	if err := r.db.Save(task).Error; err != nil {
		return Task{}, err
	}

	return *task, nil
}

func (r *repository) Delete(task *Task) error {
	return r.db.Delete(task).Error
}
