package user

import "database/sql"

// Aggregate the Aggregate for User
type Aggregate struct {
	ID        uint            `db:"id"`
	Name      sql.NullString  `db:"name"`
	Email     string          `db:"email"`
	AvatarURL sql.NullString  `db:"avatar_url"`
	Lists     []ListAggregate `db:"-"`
}

// ListAggregate the DTO for List
type ListAggregate struct {
	ID          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
}
