package database

import (
	"task/internal/entity"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *entity.Project) error
	Update(project *entity.Project) error
	FindByID(id uint) (*entity.Project, error)
	ListByUser(userID uint) ([]*entity.Project, error)
}

type ProjectRepositoryDB struct {
	db *gorm.DB
}

func NewProjectRepositoryDB(db *gorm.DB) ProjectRepository {
	return &ProjectRepositoryDB{db: db}
}

func (r *ProjectRepositoryDB) Create(project *entity.Project) error {
	return r.db.Create(project).Error
}

func (r *ProjectRepositoryDB) Update(project *entity.Project) error {
	return r.db.Save(project).Error
}

func (r *ProjectRepositoryDB) FindByID(id uint) (*entity.Project, error) {
	var project entity.Project
	if err := r.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepositoryDB) ListByUser(userID uint) ([]*entity.Project, error) {
	var projects []*entity.Project
	if err := r.db.Where("created_by = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
