package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/pkg/config"
	"time"
)

type AuthService interface {
	Login(ctx context.Context, req requests.LoginRequest) (responses.LoginResponse, error)
	LoginWithUsername(ctx context.Context, req requests.LoginWithUsernameRequest) (responses.LoginResponse, error)
	VerifyPassword(hashedPassword, plainPassword string) bool
	GenerateToken(userID uint, username, email string, roleID uint, roleName string, permissions []string, cfg *config.Config) (string, time.Time, error)
}

type authService struct {
	userService UserService
}

func NewAuthService(userService UserService) AuthService {
	return &authService{userService: userService}
}

// Login authenticates a user with email and password
func (s *authService) Login(ctx context.Context, req requests.LoginRequest) (responses.LoginResponse, error) {
	// Find user by email
	user, err := s.userService.GetByEmail(ctx, req.Email)
	if err != nil {
		return responses.LoginResponse{}, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return responses.LoginResponse{}, errors.New("user account is inactive")
	}

	// Check if user is locked
	if user.IsLocked {
		if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
			return responses.LoginResponse{}, errors.New("user account is locked")
		}
	}

	// Verify password
	if !s.VerifyPassword(user.Password, req.Password) {
		return responses.LoginResponse{}, errors.New("invalid email or password")
	}

	// Build response
	userResp := s.mapUserToResponse(user)
	return responses.LoginResponse{
		User: &userResp,
	}, nil
}

// LoginWithUsername authenticates a user with username and password
func (s *authService) LoginWithUsername(ctx context.Context, req requests.LoginWithUsernameRequest) (responses.LoginResponse, error) {
	// Find user by username
	user, err := s.userService.GetByUsername(ctx, req.Username)
	if err != nil {
		return responses.LoginResponse{}, errors.New("invalid username or password")
	}

	// Check if user is active
	if !user.IsActive {
		return responses.LoginResponse{}, errors.New("user account is inactive")
	}

	// Check if user is locked
	if user.IsLocked {
		if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
			return responses.LoginResponse{}, errors.New("user account is locked")
		}
	}

	// Verify password
	if !s.VerifyPassword(user.Password, req.Password) {
		return responses.LoginResponse{}, errors.New("invalid username or password")
	}

	// Build response
	userResp := s.mapUserToResponse(user)
	return responses.LoginResponse{
		User: &userResp,
	}, nil
}

// VerifyPassword compares a hashed password with a plain password
func (s *authService) VerifyPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

// GenerateToken generates a JWT token with user claims
func (s *authService) GenerateToken(userID uint, username, email string, roleID uint, roleName string, permissions []string, cfg *config.Config) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Duration(cfg.JWTConfig.ExpiryHours) * time.Hour)

	claims := jwt.MapClaims{
		"user_id":     userID,
		"username":    username,
		"email":       email,
		"role_id":     roleID,
		"role_name":   roleName,
		"permissions": permissions,
		"exp":         expirationTime.Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWTConfig.Secret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

// mapUserToResponse converts a User model to UserResponse DTO
func (s *authService) mapUserToResponse(user *models.User) responses.UserResponse {
	// Note: This is a simplified mapping. In production, you'd use the obfuscator
	// For now, we'll return basic info
	return responses.UserResponse{
		ID:       fmt.Sprintf("%d", user.ID),
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		IsActive: user.IsActive,
	}
}
