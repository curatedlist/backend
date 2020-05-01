package list

// DTO the DTO for List
type DTO struct {
	ID          uint
	Name        string
	Description string
}

// ToDTO transforms a List into a ListDTO
func ToDTO(list DatabaseDTO) DTO {
	return DTO{ID: list.ID, Name: list.Name, Description: list.Description}
}

// ToDTOs transforms a list of Lists into a list of ListDTOs
func ToDTOs(lists []DatabaseDTO) []DTO {
	listDTOs := make([]DTO, len(lists))

	for i, itm := range lists {
		listDTOs[i] = ToDTO(itm)
	}

	return listDTOs
}
