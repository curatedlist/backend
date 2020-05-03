package user

// DTO the DTO for User
type DTO struct {
	ID        uint
	Name      string
	Email     string
	AvatarURL string
	Lists     []ListDTO
}

// ListDTO the DTO for List
type ListDTO struct {
	ID          uint
	Name        string
	Description string
}
