package user

import (
	"backend/internal/user/commands"
)

// Service is a service that provides basic operations over Users
type Service struct {
	repository Repository
}

// NewService is a Constructor of the Service
func NewService(repository Repository) Service {
	return Service{repository: repository}
}

// Get a user
func (serv *Service) Get(id int64) DTO {
	return serv.repository.GetByID(id).ToUser()
}

// GetByEmail a user
func (serv *Service) GetByEmail(email string) DTO {
	return serv.repository.GetByEmail(email).ToUser()
}

// GetByIss a user
func (serv *Service) GetByIss(iss string) DTO {
	return serv.repository.GetByIss(iss).ToUser()
}

// GetByUsername a user
func (serv *Service) GetByUsername(email string) DTO {
	return serv.repository.GetByUsername(email).ToUser()
}

// GetLists get lists for an user
func (serv *Service) GetLists(user DTO) []ListDTO {
	lists := ToLists(serv.repository.GetLists(user.ID))
	for i := range lists {
		lists[i].Owner = user
	}
	return lists
}

// GetFavs get favorites lists for an user
func (serv *Service) GetFavs(user DTO) []ListDTO {
	return ToLists(serv.repository.GetFavs(user.ID))
}

// Create an user
func (serv *Service) Create(email string) DTO {
	return serv.repository.Create(email).ToUser()
}

// Update an user
func (serv *Service) Update(id int64, updateCommand commands.Update) DTO {
	return serv.repository.Update(id, updateCommand).ToUser()
}
