package commands

// CreateList is the command with the properties on create list
type CreateList struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
