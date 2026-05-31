package search

// Result is a single external-catalog search hit, normalized across providers.
type Result struct {
	Category    string `json:"category"` // "book" | "movie" | "music"
	Name        string `json:"name"`
	URL         string `json:"url"`
	PicURL      string `json:"pic_url"`
	Description string `json:"description"`
}
