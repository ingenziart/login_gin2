package dto

type CreateUserDto struct {
	FullName     string `json:"fullName" validate:"required,min=3"`
	Email        string `json:"email" validate:"required,email"`
	Phone        string `json:"phone" validate:"required,min=10"`
	PasswordHash string `json:"password" validate:"required,min=4"`
	Role         string `json:"role" validate:"required,oneof=admin user guest"`
	Status       string `json:"status" validate:"required,oneof=active inactive deleted"`
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
