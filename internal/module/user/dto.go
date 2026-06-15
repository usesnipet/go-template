package user

// CreateUserDTO represents the payload to create a user.
type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     Role   `json:"role" validate:"required,oneof=user admin"`
}

func (dto *CreateUserDTO) ToModel() *User {
	return &User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     dto.Role,
	}
}
