package repository

import "task/internal/entity"

type ProjectRepository interface {
	Create(project *entity.Project) error
	FindByID(id uint) (*entity.Project, error)
	ListByUser(userID uint) ([]*entity.Project, error)
}
