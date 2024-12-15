package usecase

import (
	"errors"
	"task/internal/entity"
	"task/internal/infrastructure/database"
	"time"
)

type ProjectUsecase struct {
	projectRepo database.ProjectRepository
}

func NewProjectUsecase(projectRepo database.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{projectRepo: projectRepo}
}

func (u *ProjectUsecase) CreateProject(name, description string, createdBy uint) (*entity.Project, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	project := &entity.Project{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
	}
	if err := u.projectRepo.Create(project); err != nil {
		return nil, err
	}

	return project, nil
}

func (u *ProjectUsecase) UpdateProject(project *entity.Project) error {
	return u.projectRepo.Update(project)
}

func (u *ProjectUsecase) FindByID(id uint) (*entity.Project, error) {
	return u.projectRepo.FindByID(id)
}

func (u *ProjectUsecase) ListByUser(userID uint) ([]*entity.Project, error) {
	return u.projectRepo.ListByUser(userID)
}
