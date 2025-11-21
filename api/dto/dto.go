package dto

type CreateUserDto struct {
	FullName     string `json:"fullName" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	PasswordHash string `json:"password" validate:"required"`
	Role         string `json:"role" validate:"required"`
	Status       string `json:"status" validate:"required"`
}

type UpdateUserDto struct {
	FullName     *string `json:"fullName,omitempty"`
	Email        *string `json:"email,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	PasswordHash *string `json:"password,omitempty"`
	Role         *string `json:"role,omitempty"`
	Status       *string `json:"status,omitempty"`
}

// for specific end pont
type UpdateStatusDTO struct {
	Status *string `json:"status,omitempty"`
}
