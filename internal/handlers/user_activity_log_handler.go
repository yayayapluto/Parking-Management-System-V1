package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type UserActivityLogHandler struct {
	*BaseHandler
	service services.UserActivityLogService
}

func NewUserActivityLogHandler(s services.UserActivityLogService) *UserActivityLogHandler {
	return &UserActivityLogHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all user activity logs with pagination
func (h *UserActivityLogHandler) GetAll(c *fiber.Ctx) error {
	var req requests.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		// FIX: Tambah nil di argumen ke-4 (detail error)
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
		// FIX: Tambah nil
		return h.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// Create creates a new user activity log
func (h *UserActivityLogHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateUserActivityLogRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// FIX: h.Success(ctx, data, message) - Status code diurus internal BaseHandler
	return h.Success(c, result, "User activity log created successfully")
}

// GetByID retrieves a user activity log by ID
func (h *UserActivityLogHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, result, "User activity log retrieved successfully")
}

// Delete deletes a user activity log
func (h *UserActivityLogHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, nil, "User activity log deleted successfully")
}
