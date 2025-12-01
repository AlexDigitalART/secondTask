package taskService

import (
	"errors"

	"github.com/google/uuid"
)

type CreateTaskRequest struct {
	Task   string
	IsDone bool
	UserID uuid.UUID
}

type TasksService interface {
	CreateTask(req CreateTaskRequest) (*Task, error)
	GetAllTask() ([]Task, error)
	GetTaskByID(id uuid.UUID) (Task, error)
	UpdateTask(id uuid.UUID, req UpdateTaskRequest) (Task, error)
	DeleteTask(id uuid.UUID) error
	GetTasksByUserID(userID uuid.UUID) ([]Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TasksService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(req CreateTaskRequest) (*Task, error) {
	if req.UserID == uuid.Nil {
		return nil, errors.New("user_id is required")
	}

	task := &Task{
		ID:     uuid.New(),
		Task:   req.Task,
		IsDone: req.IsDone,
		UserID: req.UserID,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskService) GetTasksByUserID(userID uuid.UUID) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *taskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}

func (s *taskService) GetTaskByID(id uuid.UUID) (Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskService) UpdateTask(id uuid.UUID, req UpdateTaskRequest) (Task, error) {
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

	if err := s.repo.UpdateTaskBy(&task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id uuid.UUID) error {
	return s.repo.DeleteTask(id)
}
