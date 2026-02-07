package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type UserHandler struct {
	BaseHandler
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Users
// @Description  Get list of users with pagination
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by username, email, or full name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /users [get]
func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateUserRequest true "User Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /users [post]
func (h *UserHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "User created successfully")
}

// GetByID godoc
// @Summary      Get User By ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "User fetched successfully")
}

// Update godoc
// @Summary      Update User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "User Hash ID"
// @Param        request body requests.UpdateUserRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /users/{id} [put]
func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "User updated successfully")
}

// Delete godoc
// @Summary      Delete User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "User deleted successfully")
}
