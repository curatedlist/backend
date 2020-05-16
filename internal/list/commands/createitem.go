package commands

// CreateItem command
type CreateItem struct {
	UserID      string `json:"user_id"`
	ListID      string `json:"list_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"pic_url"`
}
