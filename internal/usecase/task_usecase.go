package usecase

import (
	"errors"
	"task/internal/entity"
	"task/internal/infrastructure/database"
	"time"
)

type TaskUsecase struct {
	taskRepo database.TaskRepository
}

func NewTaskUsecase(taskRepo database.TaskRepository) *TaskUsecase {
	return &TaskUsecase{taskRepo: taskRepo}
}

func (u *TaskUsecase) CreateTask(projectID uint, name, description, status string, deadline time.Time, assignedTo uint) (*entity.Task, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	task := &entity.Task{
		ProjectID:   projectID,
		Name:        name,
		Description: description,
		Status:      status,
		Deadline:    deadline,
		AssignedTo:  assignedTo,
		CreatedAt:   time.Now(),
	}
	if err := u.taskRepo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (u *TaskUsecase) UpdateTask(id uint, name, description, status string, deadline time.Time, assignedTo uint) (*entity.Task, error) {
	task, err := u.taskRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("task not found")
	}

	task.Name = name
	task.Description = description
	task.Status = status
	task.Deadline = deadline
	task.AssignedTo = assignedTo

	if err := u.taskRepo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (u *TaskUsecase) DeleteTask(id uint) error {
	return u.taskRepo.Delete(id)
}

func (u *TaskUsecase) ListTasksByProject(projectID uint) ([]*entity.Task, error) {
	return u.taskRepo.ListByProject(projectID)
}

func (u *TaskUsecase) FindTaskByID(id uint) (*entity.Task, error) {
	return u.taskRepo.FindByID(id)
}
