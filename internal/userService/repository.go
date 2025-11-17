package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *User) (*User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id int) (*User, error) {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *User) (*User, error) {
	result := r.db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id int) error {
	result := r.db.Delete(&User{}, id)
	return result.Error
}
