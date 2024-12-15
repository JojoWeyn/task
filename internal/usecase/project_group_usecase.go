package usecase

import (
	"task/internal/entity"
	"task/internal/infrastructure/database"
)

type ProjectGroupUsecase struct {
	projectGroupRepo database.ProjectGroupRepository
}

func NewProjectGroupUsecase(projectGroupRepo database.ProjectGroupRepository) *ProjectGroupUsecase {
	return &ProjectGroupUsecase{projectGroupRepo: projectGroupRepo}
}

func (u *ProjectGroupUsecase) AddMember(projectID uint, userID uint, role string) error {
	return u.projectGroupRepo.AddMember(projectID, userID, role)
}

func (u *ProjectGroupUsecase) RemoveMember(projectID uint, userID uint) error {
	return u.projectGroupRepo.RemoveMember(projectID, userID)
}

func (u *ProjectGroupUsecase) GetMembers(projectID uint) ([]*entity.ProjectGroup, error) {
	return u.projectGroupRepo.GetMembers(projectID)
}

func (u *ProjectGroupUsecase) IsMember(projectID uint, userID uint) (bool, error) {
	return u.projectGroupRepo.IsMember(projectID, userID)
}

func (u *ProjectGroupUsecase) IsAdmin(projectID uint, userID uint) (bool, error) {
	return u.projectGroupRepo.IsAdmin(projectID, userID)
}
