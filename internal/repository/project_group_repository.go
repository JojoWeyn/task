package repository

import "task/internal/entity"

type ProjectGroupRepository interface {
	AddMember(projectID uint, userID uint, role string) error
	RemoveMember(projectID uint, userID uint) error
	GetMembers(projectID uint) ([]*entity.ProjectGroup, error)
	IsMember(projectID uint, userID uint) (bool, error)
	IsAdmin(projectID uint, userID uint) (bool, error)
}
