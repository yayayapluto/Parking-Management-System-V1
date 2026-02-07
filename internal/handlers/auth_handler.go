package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/config"
)

type AuthHandler struct {
	BaseHandler
	authService services.AuthService
	config      *config.Config
}

func NewAuthHandler(authService services.AuthService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		config:      cfg,
	}
}

// Login godoc
// @Summary      Login with Email
// @Description  Authenticate user with email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body requests.LoginRequest true "Login Request"
// @Success      200  {object}  responses.BaseResponse
// @Failure      401  {object}  responses.BaseResponse
// @Router       /auth/login [post]
func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var req requests.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.authService.Login(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Login successful")
}

// LoginWithUsername godoc
// @Summary      Login with Username
// @Description  Authenticate user with username and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body requests.LoginWithUsernameRequest true "Login Request"
// @Success      200  {object}  responses.BaseResponse
// @Failure      401  {object}  responses.BaseResponse
// @Router       /auth/login-username [post]
func (h *AuthHandler) LoginWithUsername(ctx *fiber.Ctx) error {
	var req requests.LoginWithUsernameRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.authService.LoginWithUsername(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Login successful")
}
