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

// ToItem transforms a item into a itemDTO
func (item ItemAggregate) ToItem() ItemDTO {
	return ItemDTO{ID: uint(item.ID.Int64), Name: item.Name.String, URL: item.URL.String, PicURL: item.PicURL.String}
}

// ToOwner transforms a user into a userDTO
func (owner OwnerAggregate) ToOwner() OwnerDTO {
	return OwnerDTO{ID: uint(owner.ID), Name: owner.Name.String, Email: owner.Email, AvatarURL: owner.AvatarURL.String}
}

// ToList transforms a List into a ListDTO
func (list Aggregate) ToList() DTO {
	return DTO{ID: list.ID, Name: list.Name, Description: list.Description, Items: ToItems(list.Items), Owner: list.Owner.ToOwner()}
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
