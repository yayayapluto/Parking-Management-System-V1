package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type PermissionHandler struct {
	*BaseHandler
	service services.PermissionService
}

func NewPermissionHandler(s services.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all permissions with pagination
func (h *PermissionHandler) GetAll(c *fiber.Ctx) error {
	var req requests.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		// FIX: Tambah nil di argumen ke-4
		return h.Error(c, fiber.StatusBadRequest, "Invalid pagination parameters", nil)
	}

	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}

	result, err := h.service.GetAll(c.Context(), req, c.OriginalURL())
	if err != nil {
		return h.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.JSON(result)
}

// Create creates a new permission
func (h *PermissionHandler) Create(c *fiber.Ctx) error {
	var req requests.CreatePermissionRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// FIX: h.Success(ctx, data, message)
	return h.Success(c, result, "Permission created successfully")
}

// GetByID retrieves a permission by ID
func (h *PermissionHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, result, "Permission retrieved successfully")
}

// Update updates a permission
func (h *PermissionHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req requests.UpdatePermissionRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Update(c.Context(), id, req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return h.Success(c, result, "Permission updated successfully")
}

// Delete deletes a permission
func (h *PermissionHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, nil, "Permission deleted successfully")
}
