package handlers

import (
	"context"
	"errors"
	"firstTask/internal/userService"
	"firstTask/internal/web/users"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(
	ctx context.Context,
	request users.GetUsersRequestObject,
) (users.GetUsersResponseObject, error) {

	usersFromDB, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, user := range usersFromDB {

		userResp := users.User{
			Id:       &user.ID,
			Email:    &user.Email,
			Password: &user.Password,
		}

		response = append(response, userResp)
	}

	return response, nil
}

func (h *UserHandler) GetUsersIdTasks(
	ctx context.Context,
	request users.GetUsersIdTasksRequestObject,
) (users.GetUsersIdTasksResponseObject, error) {

	userUUID := uuid.UUID(request.Id)

	tasksList, err := h.service.GetTasksForUser(userUUID)
	if err != nil {
		return nil, err
	}

	response := users.GetUsersIdTasks200JSONResponse{}

	for _, t := range tasksList {
		item := users.Task{
			Id:     &t.ID,
			Task:   &t.Task,
			IsDone: &t.IsDone,
			UserId: &t.UserID,
		}

		response = append(response, item)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(
	ctx context.Context,
	request users.PostUsersRequestObject,
) (users.PostUsersResponseObject, error) {

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

	resp := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return resp, nil
}

func (h *UserHandler) PatchUsersId(
	ctx context.Context,
	request users.PatchUsersIdRequestObject,
) (users.PatchUsersIdResponseObject, error) {

	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}

	updateReq := userService.UpdateUserRequest{}

	if request.Body.Email != nil {
		updateReq.Email = request.Body.Email
	}
	if request.Body.Password != nil {
		updateReq.Password = request.Body.Password
	}

	userUUID := uuid.UUID(request.Id)

	updatedUser, err := h.service.UpdateUser(userUUID, updateReq)
	if err != nil {
		return nil, err
	}

	userID := types.UUID(updatedUser.ID)

	resp := users.PatchUsersId200JSONResponse{
		Id:       &userID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return resp, nil
}

func (h *UserHandler) DeleteUsersId(
	ctx context.Context,
	request users.DeleteUsersIdRequestObject,
) (users.DeleteUsersIdResponseObject, error) {

	err := h.service.DeleteUser(request.Id)
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
