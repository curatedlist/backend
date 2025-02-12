package user

import "database/sql"

// Aggregate for User
type Aggregate struct {
	ID        sql.NullInt64  `db:"id"`
	Name      sql.NullString `db:"name"`
	Email     sql.NullString `db:"email"`
	Username  sql.NullString `db:"username"`
	Bio       sql.NullString `db:"bio"`
	AvatarURL sql.NullString `db:"avatar_url"`
	Favs      []FavAggregate `db:"-"`
	Lists     uint           `db:"-"`
}

// ToUser transforms a UserAggregate into a UserDTO
func (agg Aggregate) ToUser() DTO {
	if agg.ID.Valid {
		return DTO{
			ID:        agg.ID.Int64,
			Name:      agg.Name.String,
			Email:     agg.Email.String,
			Username:  agg.Username.String,
			Bio:       agg.Bio.String,
			AvatarURL: agg.AvatarURL.String,
			Lists:     agg.Lists,
			Favs:      ToFavs(agg.Favs),
		}
	}
	return DTO{}
}

// ListAggregate for List
type ListAggregate struct {
	ID          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
	Deleted     sql.NullBool   `db:"deleted"`
	Owner       Aggregate      `db:"-"`
}

// FavAggregate for Favs
type FavAggregate struct {
	ListID sql.NullInt64 `db:"list_id"`
}

// ToList transforms a ListAggregate into a ListDTO
func (la ListAggregate) ToList() ListDTO {
	if la.ID.Valid {
		return ListDTO{
			ID:          la.ID.Int64,
			Name:        la.Name.String,
			Description: la.Description.String,
			Deleted:     la.Deleted.Bool,
			Owner:       la.Owner.ToUser(),
		}
	}
	return ListDTO{}
}

//ToFavs transforms a slice of FavAggregate into a slice of integers
func ToFavs(favs []FavAggregate) []uint {
	favlist := make([]uint, len(favs))
	for i, fav := range favs {
		favlist[i] = uint(fav.ListID.Int64)
	}
	return favlist
}

// ToLists transforms a slice of ListAggregate into a slice of ListDTO
func ToLists(lists []ListAggregate) []ListDTO {
	listDTOs := make([]ListDTO, len(lists))

	for i, list := range lists {
		listDTOs[i] = list.ToList()
	}
	return listDTOs
}
