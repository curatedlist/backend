package list

import "database/sql"

// Aggregate the Aggregate for List
type Aggregate struct {
	ID          sql.NullInt64   `db:"id"`
	Name        sql.NullString  `db:"name"`
	Description sql.NullString  `db:"description"`
	Owner       OwnerAggregate  `db:"-"`
	Items       []ItemAggregate `db:"-"`
	Favs        uint            `db:"-"`
}

// OwnerAggregate the Aggregate for List
type OwnerAggregate struct {
	ID        sql.NullInt64  `db:"id"`
	Name      sql.NullString `db:"name"`
	Email     sql.NullString `db:"email"`
	Username  sql.NullString `db:"username"`
	Bio       sql.NullString `db:"bio"`
	AvatarURL sql.NullString `db:"avatar_url"`
}

// ItemAggregate the DTO for item
type ItemAggregate struct {
	ID          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
	URL         sql.NullString `db:"url"`
	PicURL      sql.NullString `db:"pic_url"`
	Deleted     sql.NullBool   `db:"deleted"`
	ListID      sql.NullInt64  `db:"list_id"`
}

// ToItem transforms a item into a itemDTO
func (item ItemAggregate) ToItem() ItemDTO {
	if item.ID.Valid {
		return ItemDTO{ID: uint(item.ID.Int64), Name: item.Name.String, Description: item.Description.String, URL: item.URL.String, PicURL: item.PicURL.String, Deleted: item.Deleted.Bool, ListID: uint(item.ListID.Int64)}
	}
	return ItemDTO{}
}

// ToOwner transforms a user into a userDTO
func (owner OwnerAggregate) ToOwner() OwnerDTO {
	if owner.ID.Valid {
		return OwnerDTO{ID: uint(owner.ID.Int64), Name: owner.Name.String, Username: owner.Username.String, Bio: owner.Bio.String, Email: owner.Email.String, AvatarURL: owner.AvatarURL.String}
	}
	return OwnerDTO{}
}

// ToList transforms a List into a ListDTO
func (list Aggregate) ToList() DTO {
	if list.ID.Valid {
		return DTO{ID: uint(list.ID.Int64), Name: list.Name.String, Description: list.Description.String, Items: ToItems(list.Items), Owner: list.Owner.ToOwner(), Favs: list.Favs}
	}
	return DTO{}
}

// ToItems transforms a list of Items into a list of itemsDTO
func ToItems(items []ItemAggregate) []ItemDTO {
	its := make([]ItemDTO, len(items))
	for i, itm := range items {
		its[i] = itm.ToItem()
	}
	return its
}

// ToLists transforms a list of Lists into a list of ListDTOs
func ToLists(lists []Aggregate) []DTO {
	listDTOs := make([]DTO, len(lists))

	for i, itm := range lists {
		listDTOs[i] = itm.ToList()
	}

	return listDTOs
}
