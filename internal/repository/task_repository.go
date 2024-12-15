package repository

import "task/internal/entity"

type TaskRepository interface {
	Create(task *entity.Task) error
	Update(task *entity.Task) error
	FindByID(id uint) (*entity.Task, error)
	ListByProject(projectID uint) ([]*entity.Task, error)
}
