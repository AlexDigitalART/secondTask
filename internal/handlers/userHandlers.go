package handlers

import (
	"context"
	"errors"
	"firstTask/internal/userService"
	"firstTask/internal/web/users"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	usersFromDB, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, user := range usersFromDB {
		userResp := users.User{
			Id:        &user.ID,
			Email:     &user.Email,
			Password:  &user.Password,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		}
		if user.DeletedAt != nil {
			userResp.DeletedAt = user.DeletedAt
		}
		response = append(response, userResp)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}

	if request.Body.Email == "" {
		return nil, errors.New("field 'email' is required")
	}
	if request.Body.Password == "" {
		return nil, errors.New("field 'password' is required")
	}

	userToCreate := userService.CreateUserRequest{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	createdUser, err := h.service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:        &createdUser.ID,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		CreatedAt: &createdUser.CreatedAt,
		UpdatedAt: &createdUser.UpdatedAt,
	}
	if createdUser.DeletedAt != nil {
		response.DeletedAt = createdUser.DeletedAt
	}

	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}

	updateRequest := userService.UpdateUserRequest{}

	if request.Body.Email != nil {
		updateRequest.Email = request.Body.Email
	}
	if request.Body.Password != nil {
		updateRequest.Password = request.Body.Password
	}

	updatedUser, err := h.service.UpdateUser(request.Id, updateRequest)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:        &updatedUser.ID,
		Email:     &updatedUser.Email,
		Password:  &updatedUser.Password,
		CreatedAt: &updatedUser.CreatedAt,
		UpdatedAt: &updatedUser.UpdatedAt,
	}
	if updatedUser.DeletedAt != nil {
		response.DeletedAt = updatedUser.DeletedAt
	}

	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUser(request.Id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
