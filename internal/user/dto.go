package user

// DTO the DTO for User
type DTO struct {
	ID        uint
	Name      string
	Email     string
	AvatarURL string
}

// ToDTO transforms a User into a DTO
func ToDTO(user DataBaseDTO) DTO {
	return DTO{ID: user.ID, Name: user.Name.String, Email: user.Email, AvatarURL: user.AvatarURL.String}
}
