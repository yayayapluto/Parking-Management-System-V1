package requests

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,max=50,alphanum"`
	FullName string `json:"full_name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Phone    string `json:"phone" validate:"omitempty,max=20"`
	Password string `json:"password" validate:"required,min=8"`
	RoleID   string `json:"role_id" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" validate:"omitempty,max=100"`
	Email    string `json:"email" validate:"omitempty,email,max=255"`
	Phone    string `json:"phone" validate:"omitempty,max=20"`
	RoleID   string `json:"role_id" validate:"omitempty"`
	IsActive *bool  `json:"is_active" validate:"omitempty"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}
