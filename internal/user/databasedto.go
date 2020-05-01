package user

import "database/sql"

// DataBaseDTO the DataBaseDTO for User
type DataBaseDTO struct {
	ID        uint
	Name      sql.NullString
	Email     string
	AvatarURL sql.NullString
}
