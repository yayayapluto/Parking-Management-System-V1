package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type RoleHandler struct {
	BaseHandler
	service services.RoleService
}

func NewRoleHandler(service services.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Roles
// @Description  Get list of roles with pagination
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /roles [get]
func (h *RoleHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Role
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateRoleRequest true "Role Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /roles [post]
func (h *RoleHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateRoleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Role created successfully")
}

// GetByID godoc
// @Summary      Get Role By ID
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Role Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /roles/{id} [get]
func (h *RoleHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Role fetched successfully")
}

// Update godoc
// @Summary      Update Role
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Role Hash ID"
// @Param        request body requests.UpdateRoleRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /roles/{id} [put]
func (h *RoleHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateRoleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Role updated successfully")
}

// Delete godoc
// @Summary      Delete Role
// @Tags         Roles
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Role Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /roles/{id} [delete]
func (h *RoleHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Role deleted successfully")
}
