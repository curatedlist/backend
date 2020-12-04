package list

// DTO for List
type DTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deleted     bool      `json:"deleted"`
	Items       []ItemDTO `json:"items"`
	Owner       OwnerDTO  `json:"owner"`
	Favs        uint      `json:"favs"`
}

// ItemDTO for item
type ItemDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"pic_url"`
	Deleted     bool   `json:"deleted"`
	ListID      int64  `json:"list_id"`
}

// OwnerDTO for Owner
type OwnerDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}
