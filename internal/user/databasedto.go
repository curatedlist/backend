package user

import "database/sql"

// DataBaseDTO the DataBaseDTO for User
type DataBaseDTO struct {
	ID        uint
	Name      sql.NullString
	Email     string
	AvatarURL sql.NullString
	Lists     []DatabaseListDTO
}

// DatabaseListDTO the DTO for List
type DatabaseListDTO struct {
	ID          sql.NullInt64
	Name        sql.NullString
	Description sql.NullString
}
