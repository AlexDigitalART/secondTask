package taskService

import (
	"github.com/google/uuid"
)

type TasksService interface {
	CreateTask(text string) (Task, error)
	GetAllTask() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id string, req UpdateTaskRequest) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TasksService { // Возвращаем интерфейс
	return &taskService{repo: r} // Возвращаем указатель
}

func (s *taskService) CreateTask(text string) (Task, error) {

	task := Task{
		ID:     uuid.NewString(),
		Task:   text,
		IsDone: false,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) GetAllTask() ([]Task, error) {
	var task []Task
	var err error
	if task, err = s.repo.GetAllTask(); err != nil {
		return []Task{}, err
	}
	return task, nil
}

func (s *taskService) GetTaskByID(id string) (Task, error) {
	var task Task
	var err error
	if task, err = s.repo.GetTaskByID(id); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) UpdateTask(id string, req UpdateTaskRequest) (Task, error) {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	if req.Task != nil {
		task.Task = *req.Task
	}
	if req.IsDone != nil {
		task.IsDone = *req.IsDone
	}

	if err := s.repo.UpdateTaskBy(task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
