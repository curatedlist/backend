package commands

// CreateItem command
type CreateItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"pic_url"`
}
