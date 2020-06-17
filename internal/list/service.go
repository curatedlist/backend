package list

import (
	"backend/internal/item"
	"backend/internal/list/commands"
)

// Service provides basic operations over Lists
type Service struct {
	repository Repository
}

// NewService returns a new Service
func NewService(repository Repository) Service {
	return Service{repository: repository}
}

// FindAll the available lists
func (serv *Service) FindAll(filter string) []DTO {
	return ToLists(serv.repository.FindAll(filter))
}

// Get a list id
func (serv *Service) Get(id int64) DTO {
	return serv.repository.Get(id).ToList()
}

// GetItem of a list
func (serv *Service) GetItem(id int64) ItemDTO {
	return serv.repository.GetItem(id).ToItem()
}

// Create a list
func (serv *Service) Create(userID int64, command commands.CreateList) DTO {
	return serv.repository.Create(userID, command).ToList()
}

// Delete a list
func (serv *Service) Delete(id int64) DTO {
	return serv.repository.Delete(id).ToList()
}

// CreateItem for a list
func (serv *Service) CreateItem(id int64, command commands.CreateItem) ItemDTO {
	url := item.GetMetaData(command.URL)
	command.PicURL = url
	return serv.repository.CreateItem(id, command).ToItem()
}

// DeleteItem from a list
func (serv *Service) DeleteItem(itemID int64) ItemDTO {
	return serv.repository.DeleteItem(itemID).ToItem()
}

// Fav a list
func (serv *Service) Fav(id int64, userID int64) DTO {
	return serv.repository.Fav(id, userID).ToList()
}

// Unfav a list
func (serv *Service) Unfav(id int64, userID int64) DTO {
	return serv.repository.Unfav(id, userID).ToList()
}
