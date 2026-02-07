package services

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/pkg/helpers"
)

type UserService interface {
	GetAll(ctx context.Context, p requests.PaginationRequest, url string) (responses.PageResponse, error)
	GetByID(ctx context.Context, hashedID string) (responses.UserResponse, error)
	Create(ctx context.Context, req requests.CreateUserRequest) (responses.UserResponse, error)
	Update(ctx context.Context, hashedID string, req requests.UpdateUserRequest) (responses.UserResponse, error)
	Delete(ctx context.Context, hashedID string) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}

type userService struct {
	repo       repos.UserRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

func NewUserService(r repos.UserRepository, o helpers.IDObfuscator, v *validator.Validate) UserService {
	return &userService{repo: r, obfuscator: o, validator: v}
}

// GetAll retrieves all users with pagination
func (s *userService) GetAll(ctx context.Context, req requests.PaginationRequest, url string) (responses.PageResponse, error) {
	mdls, total, err := s.repo.GetWithPagination(ctx, req, "username", "email", "full_name")
	if err != nil {
		return responses.PageResponse{}, err
	}

	var list []responses.UserResponse
	for _, m := range mdls {
		list = append(list, s.mapToResponse(m))
	}

	return responses.PageResponse{
		Success:    true,
		Message:    "Users retrieved successfully",
		Data:       list,
		Pagination: responses.CreateMeta(req, total, url),
	}, nil
}

// GetByID retrieves a user by hashed ID
func (s *userService) GetByID(ctx context.Context, hashID string) (responses.UserResponse, error) {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.UserResponse{}, errors.New("invalid user id")
	}

	model, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.UserResponse{}, err
	}

	return s.mapToResponse(*model), nil
}

// Create creates a new user with hashed password
func (s *userService) Create(ctx context.Context, req requests.CreateUserRequest) (responses.UserResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.UserResponse{}, err
	}

	// Decode RoleID
	roleID, err := s.obfuscator.Decode(req.RoleID)
	if err != nil {
		return responses.UserResponse{}, errors.New("invalid role id")
	}

	// Hash password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return responses.UserResponse{}, errors.New("failed to hash password")
	}

	user := models.User{
		Username:     req.Username,
		FullName:     req.FullName,
		Email:        req.Email,
		Phone:        req.Phone,
		Password:     string(hashedPassword),
		RoleID:       roleID,
		IsActive:     true,
		IsLocked:     false,
		FailedLoginAttempts: 0,
	}

	if err := s.repo.Create(ctx, &user); err != nil {
		return responses.UserResponse{}, err
	}

	return s.mapToResponse(user), nil
}

// Update updates an existing user
func (s *userService) Update(ctx context.Context, hashID string, req requests.UpdateUserRequest) (responses.UserResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return responses.UserResponse{}, err
	}

	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return responses.UserResponse{}, errors.New("invalid id format")
	}

	// Get existing user
	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return responses.UserResponse{}, err
	}

	// Patch data (only update provided fields)
	if req.FullName != "" {
		existing.FullName = req.FullName
	}
	if req.Email != "" {
		existing.Email = req.Email
	}
	if req.Phone != "" {
		existing.Phone = req.Phone
	}
	if req.RoleID != "" {
		roleID, err := s.obfuscator.Decode(req.RoleID)
		if err != nil {
			return responses.UserResponse{}, errors.New("invalid role id")
		}
		existing.RoleID = roleID
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	// Save
	if err := s.repo.Update(ctx, existing); err != nil {
		return responses.UserResponse{}, err
	}

	return s.mapToResponse(*existing), nil
}

// Delete deletes a user
func (s *userService) Delete(ctx context.Context, hashID string) error {
	id, err := s.obfuscator.Decode(hashID)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.repo.Delete(ctx, id)
}

// GetByEmail retrieves a user by email
func (s *userService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, errors.New("user not found")
}

// GetByUsername retrieves a user by username
func (s *userService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		if u.Username == username {
			return &u, nil
		}
	}

	return nil, errors.New("user not found")
}

// mapToResponse converts a User model to UserResponse DTO
func (s *userService) mapToResponse(m models.User) responses.UserResponse {
	hID, _ := s.obfuscator.Encode(m.ID)
	hRoleID, _ := s.obfuscator.Encode(m.RoleID)

	var lastLoginAt *string
	if m.LastLoginAt != nil {
		lastLoginAtStr := m.LastLoginAt.Format("2006-01-02 15:04:05")
		lastLoginAt = &lastLoginAtStr
	}

	var lockedUntil *string
	if m.LockedUntil != nil {
		lockedUntilStr := m.LockedUntil.Format("2006-01-02 15:04:05")
		lockedUntil = &lockedUntilStr
	}

	return responses.UserResponse{
		ID:                  hID,
		Username:            m.Username,
		FullName:            m.FullName,
		Email:               m.Email,
		Phone:               m.Phone,
		RoleID:              hRoleID,
		IsActive:            m.IsActive,
		IsLocked:            m.IsLocked,
		LastLoginAt:         lastLoginAt,
		LastLoginIP:         m.LastLoginIP,
		FailedLoginAttempts: m.FailedLoginAttempts,
		LockedUntil:         lockedUntil,
		CreatedAt:           m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:           m.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
