package taskService

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *Task) error
	GetAllTask() ([]Task, error)
	GetTaskByID(id uuid.UUID) (Task, error)
	UpdateTaskBy(task *Task) error
	DeleteTask(id uuid.UUID) error
	GetTasksByUserID(userID uuid.UUID) ([]Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task *Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetTasksByUserID(userID uuid.UUID) ([]Task, error) {
	var tasks []Task
	result := r.db.Where("user_id = ?", userID).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *taskRepository) GetAllTask() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskByID(id uuid.UUID) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskRepository) UpdateTaskBy(task *Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) DeleteTask(id uuid.UUID) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
