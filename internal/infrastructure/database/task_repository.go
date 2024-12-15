package database

import (
	"task/internal/entity"
	"task/internal/repository"

	"gorm.io/gorm"
)

type TaskRepositoryDB struct {
	db *gorm.DB
}

func NewTaskRepositoryDB(db *gorm.DB) repository.TaskRepository {
	return &TaskRepositoryDB{db: db}
}

func (r *TaskRepositoryDB) Create(task *entity.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepositoryDB) Update(task *entity.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepositoryDB) FindByID(id uint) (*entity.Task, error) {
	var task entity.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryDB) ListByProject(projectID uint) ([]*entity.Task, error) {
	var tasks []*entity.Task
	if err := r.db.Where("project_id = ?", projectID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
