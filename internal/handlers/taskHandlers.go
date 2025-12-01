package handlers

import (
	"context"
	"errors"
	"firstTask/internal/taskService"
	"firstTask/internal/web/tasks"
	"fmt"

	"github.com/google/uuid"
)

type TaskHandler struct {
	service taskService.TasksService
}

func NewTaskHandler(service taskService.TasksService) *TaskHandler {
	return &TaskHandler{service: service}
}

func toUUID(v interface{}) (uuid.UUID, error) {
	if u, ok := v.(uuid.UUID); ok {
		return u, nil
	}

	type strer interface{ String() string }
	if s, ok := v.(strer); ok {
		return uuid.Parse(s.String())
	}

	s := fmt.Sprintf("%v", v)
	return uuid.Parse(s)
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id, err := toUUID(request.Id)
	if err != nil {
		return nil, err
	}

	if err := h.service.DeleteTask(id); err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id, err := toUUID(request.Id)
	if err != nil {
		return nil, err
	}

	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}

	updateReq := taskService.UpdateTaskRequest{}
	if request.Body.Task != nil {
		updateReq.Task = request.Body.Task
	}
	if request.Body.IsDone != nil {
		updateReq.IsDone = request.Body.IsDone
	}

	updatedTask, err := h.service.UpdateTask(id, updateReq)
	if err != nil {
		return nil, err
	}

	resp := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
		UserId: &updatedTask.UserID,
	}
	return resp, nil
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	tasksFromDB, err := h.service.GetAllTask()
	if err != nil {
		return nil, err
	}

	resp := tasks.GetTasks200JSONResponse{}

	for _, t := range tasksFromDB {
		item := tasks.Task{
			Id:     &t.ID,
			Task:   &t.Task,
			IsDone: &t.IsDone,
			UserId: &t.UserID,
		}
		resp = append(resp, item)
	}

	return resp, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}

	userID, err := toUUID(request.Body.UserId)
	if err != nil {
		return nil, errors.New("field 'user_id' is required and must be uuid")
	}

	isDone := false
	if request.Body.IsDone != nil {
		isDone = *request.Body.IsDone
	}

	createReq := taskService.CreateTaskRequest{
		Task:   request.Body.Task,
		IsDone: isDone,
		UserID: userID,
	}

	createdTask, err := h.service.CreateTask(createReq)
	if err != nil {
		return nil, err
	}

	resp := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return resp, nil
}
