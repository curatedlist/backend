package list

import "database/sql"

// Aggregate the Aggregate for List
type Aggregate struct {
	ID          uint            `db:"id"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Owner       OwnerAggregate  `db:"-"`
	Items       []ItemAggregate `db:"-"`
}

// OwnerAggregate the Aggregate for List
type OwnerAggregate struct {
	ID        uint           `db:"id"`
	Name      sql.NullString `db:"name"`
	Email     string         `db:"email"`
	AvatarURL sql.NullString `db:"avatar_url"`
}

// ItemAggregate the DTO for item
type ItemAggregate struct {
	ID     sql.NullInt64  `db:"id"`
	Name   sql.NullString `db:"name"`
	URL    sql.NullString `db:"url"`
	PicURL sql.NullString `db:"pic_url"`
}
