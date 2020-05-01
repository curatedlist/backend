package list

// DTO the DTO for List
type DTO struct {
	ID          uint
	Name        string
	Description string
	Items       []ItemDTO
	Owner       OwnerDTO
}

// ItemDTO the DTO for item
type ItemDTO struct {
	ID     uint
	Name   string
	URL    string
	PicURL string
}

// OwnerDTO the DTO for Owner
type OwnerDTO struct {
	ID        uint
	Name      string
	Email     string
	AvatarURL string
}

// ToItemsDTO transforms a list of Items into a list of itemsDTO
func ToItemsDTO(items []DatabaseItemDTO) []ItemDTO {
	itemsDTOs := make([]ItemDTO, len(items))

	for i, itm := range items {
		itemsDTOs[i] = ToItemDTO(itm)
	}
	return itemsDTOs
}

// ToItemDTO transforms a item into a itemDTO
func ToItemDTO(item DatabaseItemDTO) ItemDTO {
	return ItemDTO{ID: uint(item.ID.Int64), Name: item.Name.String, URL: item.URL.String, PicURL: item.PicURL.String}
}

// ToOwnerDTO transforms a user into a userDTO
func ToOwnerDTO(owner DatabaseOwnerDTO) OwnerDTO {
	return OwnerDTO{ID: uint(owner.ID), Name: owner.Name.String, Email: owner.Email, AvatarURL: owner.AvatarURL.String}
}

// ToDTO transforms a List into a ListDTO
func ToDTO(list DatabaseDTO) DTO {
	return DTO{ID: list.ID, Name: list.Name, Description: list.Description, Items: ToItemsDTO(list.Items), Owner: ToOwnerDTO(list.Owner)}
}

// ToDTOs transforms a list of Lists into a list of ListDTOs
func ToDTOs(lists []DatabaseDTO) []DTO {
	listDTOs := make([]DTO, len(lists))

	for i, itm := range lists {
		listDTOs[i] = ToDTO(itm)
	}

	return listDTOs
}
