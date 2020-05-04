package list

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
