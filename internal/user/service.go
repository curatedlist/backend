package user

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
	return ToDTO(serv.repository.Get(id))
}
