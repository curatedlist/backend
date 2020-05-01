package list

import "database/sql"

// DatabaseDTO the DTO for List
type DatabaseDTO struct {
	ID          uint
	Name        string
	Description string
	Items       []DatabaseItemDTO
	Owner       DatabaseOwnerDTO
}

// DatabaseItemDTO the DTO for item
type DatabaseItemDTO struct {
	ID     sql.NullInt64
	Name   sql.NullString
	URL    sql.NullString
	PicURL sql.NullString
}

// DatabaseOwnerDTO the DTO for Owner
type DatabaseOwnerDTO struct {
	ID        uint
	Name      sql.NullString
	Email     string
	AvatarURL sql.NullString
}
