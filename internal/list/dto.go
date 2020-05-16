package list

// DTO the DTO for List
type DTO struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Items       []ItemDTO `json:"items"`
	Owner       OwnerDTO  `json:"owner"`
}

// ItemDTO the DTO for item
type ItemDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"pic_url"`
	Deleted     bool   `json:"deleted"`
	ListID      uint   `json:"list_id"`
}

// OwnerDTO the DTO for Owner
type OwnerDTO struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}
