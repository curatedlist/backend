package user

import "backend/internal/user/commands"

// Service is a service that provides basic operations over Users
type Service struct {
	repository Repository
}

// NewService is a Constructor of the Service
func NewService(repository Repository) Service {
	return Service{repository: repository}
}

// Get a list by id
func (serv *Service) Get(id string) DTO {
	return serv.repository.GetByID(id).ToUser()
}

// GetByEmail a list by email
func (serv *Service) GetByEmail(email string) DTO {
	return serv.repository.GetByEmail(email).ToUser()
}

// GetByUsername a list by email
func (serv *Service) GetByUsername(email string) DTO {
	return serv.repository.GetByUsername(email).ToUser()
}

// CreateUser creates an user
func (serv *Service) CreateUser(email string) DTO {
	return serv.repository.CreateUser(email).ToUser()
}

// UpdateUser creates an user
func (serv *Service) UpdateUser(id string, updateCommand commands.Update) DTO {
	return serv.repository.UpdateUser(id, updateCommand).ToUser()
}
