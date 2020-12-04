package list

import "database/sql"

// Aggregate for List
type Aggregate struct {
	ID          sql.NullInt64   `db:"id"`
	Name        sql.NullString  `db:"name"`
	Description sql.NullString  `db:"description"`
	Deleted     sql.NullBool    `db:"deleted"`
	Owner       OwnerAggregate  `db:"-"`
	Items       []ItemAggregate `db:"-"`
	Favs        uint            `db:"-"`
}

// ToList transforms a ListAggregate into a ListDTO
func (list Aggregate) ToList() DTO {
	if list.ID.Valid {
		return DTO{
			ID:          list.ID.Int64,
			Name:        list.Name.String,
			Description: list.Description.String,
			Deleted:     list.Deleted.Bool,
			Items:       ToItems(list.Items),
			Owner:       list.Owner.ToOwner(),
			Favs:        list.Favs,
		}
	}
	return DTO{}
}

// ToLists transforms a slice of ListAggregate into a slice of ListDTO
func ToLists(lists []Aggregate) []DTO {
	listDTOs := make([]DTO, len(lists))
	for i, itm := range lists {
		listDTOs[i] = itm.ToList()
	}
	return listDTOs
}

// OwnerAggregate for Owner
type OwnerAggregate struct {
	ID        sql.NullInt64  `db:"id"`
	Name      sql.NullString `db:"name"`
	Email     sql.NullString `db:"email"`
	Username  sql.NullString `db:"username"`
	Bio       sql.NullString `db:"bio"`
	AvatarURL sql.NullString `db:"avatar_url"`
}

// ToOwner transforms a OwnerAggregate into a OwnerDTO
func (owner OwnerAggregate) ToOwner() OwnerDTO {
	if owner.ID.Valid {
		return OwnerDTO{
			ID:        owner.ID.Int64,
			Name:      owner.Name.String,
			Username:  owner.Username.String,
			Bio:       owner.Bio.String,
			Email:     owner.Email.String,
			AvatarURL: owner.AvatarURL.String,
		}
	}
	return OwnerDTO{}
}

// ItemAggregate for item
type ItemAggregate struct {
	ID          sql.NullInt64  `db:"id"`
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
	URL         sql.NullString `db:"url"`
	PicURL      sql.NullString `db:"pic_url"`
	Deleted     sql.NullBool   `db:"deleted"`
	ListID      sql.NullInt64  `db:"list_id"`
}

// ToItem transforms a ItemAggregate into a ItemDTO
func (item ItemAggregate) ToItem() ItemDTO {
	if item.ID.Valid {
		return ItemDTO{
			ID:          item.ID.Int64,
			Name:        item.Name.String,
			Description: item.Description.String,
			URL:         item.URL.String,
			PicURL:      item.PicURL.String,
			Deleted:     item.Deleted.Bool,
			ListID:      item.ListID.Int64,
		}
	}
	return ItemDTO{}
}

// ToItems transforms a slice of ItemAggregate into a slice of ItemsDTO
func ToItems(items []ItemAggregate) []ItemDTO {
	its := make([]ItemDTO, len(items))
	for i, itm := range items {
		its[i] = itm.ToItem()
	}
	return its
}
