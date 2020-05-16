package list

import (
	"backend/internal/item"
	"backend/internal/list/commands"
)

// Service is a service that provides basic operations over Lists
type Service struct {
	repository Repository
}

// NewService is a Constructor of the ListService
func NewService(repository Repository) Service {
	return Service{repository: repository}
}

// FindAll finds all the lists availables
func (serv *Service) FindAll() []DTO {
	return ToLists(serv.repository.FindAll())
}

// Get a list by id
func (serv *Service) Get(id string) DTO {
	return serv.repository.Get(id).ToList()
}

// GetItem a item by id
func (serv *Service) GetItem(id string) ItemDTO {
	return serv.repository.GetItem(id).ToItem()
}

// CreateList creates a list
func (serv *Service) CreateList(userID string, createListCommand commands.CreateList) int64 {
	return serv.repository.CreateList(userID, createListCommand)
}

// CreateItem creates a item for a list
func (serv *Service) CreateItem(userID string, createItemCommand commands.CreateItem) ItemDTO {
	url := item.GetMetaData(createItemCommand.URL)
	createItemCommand.PicURL = url
	return serv.repository.CreateItem(userID, createItemCommand).ToItem()
}

// DeleteItem creates a item for a list
func (serv *Service) DeleteItem(itemID string) ItemDTO {
	return serv.repository.DeleteItem(itemID).ToItem()
}

// FavList favs a list
func (serv *Service) FavList(listID string, userID string) DTO {
	return serv.repository.FavList(listID, userID).ToList()
}
