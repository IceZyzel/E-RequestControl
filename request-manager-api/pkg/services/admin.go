package services

import (
	Request_Manager "request_manager_api"
	"request_manager_api/pkg/repository"
)

type AdministratorService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdministratorService {
	return &AdministratorService{repo: repo}
}
func (s *AdministratorService) GetUserByID(userID int) (Request_Manager.User, error) {
	return s.repo.GetUserByID(userID)
}
