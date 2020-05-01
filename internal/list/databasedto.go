package list

import "database/sql"

// DatabaseDTO the DTO for List
type DatabaseDTO struct {
	ID          uint
	Name        string
	Description string
	Items       []DatabaseItemDTO
}

// DatabaseItemDTO the DTO for item
type DatabaseItemDTO struct {
	ID     sql.NullInt64
	Name   sql.NullString
	URL    sql.NullString
	PicURL sql.NullString
}
