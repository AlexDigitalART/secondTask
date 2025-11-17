package userService

type UserService interface {
	CreateUser(request CreateUserRequest) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id int, request UpdateUserRequest) (*User, error)
	DeleteUser(id int) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(request CreateUserRequest) (*User, error) {
	user := &User{
		Email:    request.Email,
		Password: request.Password,
	}
	return s.repo.CreateUser(user)
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) UpdateUser(id int, request UpdateUserRequest) (*User, error) {
	// Сначала получаем пользователя
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Обновляем поля, если они переданы
	if request.Email != nil {
		user.Email = *request.Email
	}
	if request.Password != nil {
		user.Password = *request.Password
	}

	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
