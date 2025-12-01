package userService

import (
	"firstTask/internal/taskService"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(request CreateUserRequest) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id uuid.UUID, request UpdateUserRequest) (*User, error)
	DeleteUser(id uuid.UUID) error
	GetTasksForUser(userID uuid.UUID) ([]taskService.Task, error)
}

type userService struct {
	repo        UserRepository
	taskService taskService.TasksService
}

func NewUserService(repo UserRepository, taskService taskService.TasksService) UserService {
	return &userService{
		repo:        repo,
		taskService: taskService,
	}
}

func (s *userService) GetTasksForUser(userID uuid.UUID) ([]taskService.Task, error) {
	return s.taskService.GetTasksByUserID(userID)
}

func (s *userService) CreateUser(request CreateUserRequest) (*User, error) {
	user := &User{
		ID:       uuid.New(),
		Email:    request.Email,
		Password: request.Password,
	}
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) UpdateUser(id uuid.UUID, request UpdateUserRequest) (*User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if request.Email != nil {
		user.Email = *request.Email
	}
	if request.Password != nil {
		user.Password = *request.Password
	}

	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
