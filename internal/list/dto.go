package list

// DTO the DTO for List
type DTO struct {
	ID          uint
	Name        string
	Description string
	Items       []ItemDTO
	Owner       OwnerDTO
}

// ItemDTO the DTO for item
type ItemDTO struct {
	ID      uint
	Name    string
	URL     string
	PicURL  string
	Deleted bool
	ListID  uint
}

// OwnerDTO the DTO for Owner
type OwnerDTO struct {
	ID        uint
	Name      string
	Email     string
	AvatarURL string
}
