package user

// ToDTO transforms a User into a DTO
func ToDTO(user Aggregate) DTO {
	return DTO{ID: user.ID, Name: user.Name.String, Email: user.Email, AvatarURL: user.AvatarURL.String, Lists: ToListDTOs(user.Lists)}
}

// ToListDTOs transforms an array of Lists from the database into a ListDTO
func ToListDTOs(lists []ListAggregate) []ListDTO {
	listDTOs := make([]ListDTO, len(lists))

	for i, list := range lists {
		listDTOs[i] = ToListDTO(list)
	}
	return listDTOs
}

// ToListDTO transforms a item into a itemDTO
func ToListDTO(list ListAggregate) ListDTO {
	return ListDTO{ID: uint(list.ID.Int64), Name: list.Name.String, Description: list.Description.String}
}
