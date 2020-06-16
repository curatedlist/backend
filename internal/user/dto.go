package user

// DTO for User
type DTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
	Lists     uint   `json:"lists"`
	Favs      []uint `json:"favs"`
}

// ListDTO for List
type ListDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Deleted     bool   `json:"deleted"`
	Owner       DTO    `json:"owner"`
}
