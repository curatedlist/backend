package user

import "database/sql"

// Aggregate the Aggregate for User
type Aggregate struct {
	ID        sql.NullInt64   `db:"id"`
	Name      sql.NullString  `db:"name"`
	Email     sql.NullString  `db:"email"`
	AvatarURL sql.NullString  `db:"avatar_url"`
	Lists     []ListAggregate `db:"-"`
}

// ListAggregate the DTO for List
type ListAggregate struct {
	ID          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
}

// ToUser transforms a User into a DTO
func (agg Aggregate) ToUser() DTO {
	if agg.ID.Valid {
		return DTO{ID: uint(agg.ID.Int64), Name: agg.Name.String, Email: agg.Email.String, AvatarURL: agg.AvatarURL.String, Lists: ToLists(agg.Lists)}
	}
	return DTO{}
}

// ToList transforms a item into a itemDTO
func (la ListAggregate) ToList() ListDTO {
	if la.ID.Valid {
		return ListDTO{ID: uint(la.ID.Int64), Name: la.Name.String, Description: la.Description.String}
	}
	return ListDTO{}
}

// ToLists transforms an array of Lists from the database into a ListDTO
func ToLists(lists []ListAggregate) []ListDTO {
	listDTOs := make([]ListDTO, len(lists))

	for i, list := range lists {
		listDTOs[i] = list.ToList()
	}
	return listDTOs
}
