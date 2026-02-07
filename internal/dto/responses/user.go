package responses

type UserResponse struct {
	ID                  string `json:"id"`
	Username            string `json:"username"`
	FullName            string `json:"full_name"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	RoleID              string `json:"role_id"`
	Role                *RoleResponse `json:"role,omitempty"`
	IsActive            bool   `json:"is_active"`
	IsLocked            bool   `json:"is_locked"`
	LastLoginAt         *string `json:"last_login_at"`
	LastLoginIP         string `json:"last_login_ip"`
	FailedLoginAttempts int    `json:"failed_login_attempts"`
	LockedUntil         *string `json:"locked_until"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}
