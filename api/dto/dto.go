package dto

type CreateUserDto struct {
	FullName string `json:"fullName" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,min=10,numeric"`
	Password string `json:"password" validate:"required,min=4"`
	Role     string `json:"role" validate:"required,oneof=admin user guest"`
	Status   string `json:"status" validate:"required,oneof=active inactive deleted"`
}

type UpdateUserDto struct {
	FullName *string `json:"fullName,omitempty" validate:"omitempty,min=3"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone    *string `json:"phone,omitempty" validate:"omitempty,min=10,numeric"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=4"`
	Role     *string `json:"role,omitempty" validate:"omitempty,oneof=admin user guest"`
	Status   *string `json:"status,omitempty" validate:"omitempty,oneof=active inactive deleted"`
}

// for specific end pont
type UpdateStatusDTO struct {
	Status *string `json:"status,omitempty" validate:"required,oneof=active inactive deleted"`
}
