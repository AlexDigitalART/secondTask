package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTask() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTaskBy(task Task) error
	DeleteTask(id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) CreateTask(task Task) error {
	return t.db.Create(&task).Error
}

func (t *taskRepository) GetAllTask() ([]Task, error) {
	var tasks []Task
	err := t.db.Find(&tasks).Error
	return tasks, err
}

func (t *taskRepository) GetTaskByID(id string) (Task, error) {
	var task Task
	err := t.db.First(&task, "id = ?", id).Error
	return task, err
}

func (t *taskRepository) UpdateTaskBy(task Task) error {
	return t.db.Save(&task).Error
}

func (t *taskRepository) DeleteTask(id string) error {
	return t.db.Delete(&Task{}, "id = ?", id).Error
}
