package user

// DTO the DTO for User
type DTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	Lists     []ListDTO `json:"lists"`
	Favs      []uint    `json:"favs"`
}

// ListDTO the DTO for List
type ListDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
