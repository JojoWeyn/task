package database

import (
	"task/internal/entity"

	"gorm.io/gorm"
)

type ProjectGroupRepository interface {
	AddMember(projectID uint, userID uint, role string) error
	RemoveMember(projectID uint, userID uint) error
	GetMembers(projectID uint) ([]*entity.ProjectGroup, error)
	IsMember(projectID uint, userID uint) (bool, error)
	IsAdmin(projectID uint, userID uint) (bool, error)
}

type ProjectGroupRepositoryDB struct {
	db *gorm.DB
}

func NewProjectGroupRepositoryDB(db *gorm.DB) ProjectGroupRepository {
	return &ProjectGroupRepositoryDB{db: db}
}

func (r *ProjectGroupRepositoryDB) IsAdmin(projectID uint, userID uint) (bool, error) {
	var projectGroup entity.ProjectGroup
	if err := r.db.Where("project_id = ? AND user_id = ? AND role = ?", projectID, userID, "admin").First(&projectGroup).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProjectGroupRepositoryDB) IsMember(projectID uint, userID uint) (bool, error) {
	var projectGroup entity.ProjectGroup
	if err := r.db.Where("project_id = ? AND user_id = ?", projectID, userID).First(&projectGroup).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProjectGroupRepositoryDB) AddMember(projectID uint, userID uint, role string) error {
	return r.db.Create(&entity.ProjectGroup{ProjectID: projectID, UserID: userID, Role: role}).Error
}

func (r *ProjectGroupRepositoryDB) RemoveMember(projectID uint, userID uint) error {
	return r.db.Where("project_id = ? AND user_id = ?", projectID, userID).Delete(&entity.ProjectGroup{}).Error
}

func (r *ProjectGroupRepositoryDB) GetMembers(projectID uint) ([]*entity.ProjectGroup, error) {
	var projectGroups []*entity.ProjectGroup
	if err := r.db.Where("project_id = ?", projectID).Find(&projectGroups).Error; err != nil {
		return nil, err
	}
	return projectGroups, nil
}

func (r *ProjectGroupRepositoryDB) FindByID(id uint) (*entity.ProjectGroup, error) {
	var projectGroup entity.ProjectGroup
	if err := r.db.First(&projectGroup, id).Error; err != nil {
		return nil, err
	}
	return &projectGroup, nil
}
