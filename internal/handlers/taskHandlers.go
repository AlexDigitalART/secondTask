package handlers

import (
	"context"
	"errors"
	"firstTask/internal/taskService"
	"firstTask/internal/web/tasks"

	"github.com/google/uuid"
)

type TaskHandler struct {
	service taskService.TasksService
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	idStr := request.Id.String()

	err := h.service.DeleteTask(idStr)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	idStr := request.Id.String()

	if request.Body == nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	updateRequest := taskService.UpdateTaskRequest{
		Task:   request.Body.Task,
		IsDone: request.Body.IsDone,
	}

	updatedTask, err := h.service.UpdateTask(idStr, updateRequest)
	if err != nil {
		return nil, err
	}

	updatedTaskID, err := uuid.Parse(updatedTask.ID)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTaskID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	task, err := h.service.GetAllTask()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range task {
		taskID, err := uuid.Parse(tsk.ID)
		if err != nil {
			return nil, err
		}

		task := tasks.Task{
			Id:     &taskID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, errors.New("request body is empty")
	}
	if request.Body.Task == nil {
		return nil, errors.New("field 'task' is required")
	}

	isDone := false
	if request.Body.IsDone != nil {
		isDone = *request.Body.IsDone
	}

	taskToCreate := taskService.Task{
		Task:   *request.Body.Task,
		IsDone: isDone,
	}

	createdTask, err := h.service.CreateTask(taskToCreate.Task)
	if err != nil {
		return nil, err
	}

	createdTaskID, err := uuid.Parse(createdTask.ID)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTaskID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func NewTaskHandler(service taskService.TasksService) *TaskHandler {
	return &TaskHandler{service: service}
}
