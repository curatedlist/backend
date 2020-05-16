package user

import "database/sql"

// Aggregate the Aggregate for User
type Aggregate struct {
	ID        sql.NullInt64   `db:"id"`
	Name      sql.NullString  `db:"name"`
	Email     sql.NullString  `db:"email"`
	AvatarURL sql.NullString  `db:"avatar_url"`
	Lists     []ListAggregate `db:"-"`
	Favs      []FavAggregate  `db:"-"`
}

// ListAggregate the DTO for List
type ListAggregate struct {
	ID          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
}

// ListAggregate the DTO for List
type FavAggregate struct {
	ListID sql.NullInt64 `db:"list_id"`
}

// ToUser transforms a User into a DTO
func (agg Aggregate) ToUser() DTO {
	if agg.ID.Valid {
		return DTO{ID: uint(agg.ID.Int64), Name: agg.Name.String, Email: agg.Email.String, AvatarURL: agg.AvatarURL.String, Lists: ToLists(agg.Lists), Favs: ToFavs(agg.Favs)}
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

//ToFavs to favs
func ToFavs(favs []FavAggregate) []uint {
	favlist := make([]uint, len(favs))
	for i, fav := range favs {
		favlist[i] = uint(fav.ListID.Int64)
	}
	return favlist
}

// ToLists transforms an array of Lists from the database into a ListDTO
func ToLists(lists []ListAggregate) []ListDTO {
	listDTOs := make([]ListDTO, len(lists))

	for i, list := range lists {
		listDTOs[i] = list.ToList()
	}
	return listDTOs
}
