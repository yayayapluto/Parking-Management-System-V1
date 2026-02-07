package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type SystemLogHandler struct {
	*BaseHandler
	service services.SystemLogService
}

func NewSystemLogHandler(s services.SystemLogService) *SystemLogHandler {
	return &SystemLogHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all system logs with pagination
func (h *SystemLogHandler) GetAll(c *fiber.Ctx) error {
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
		// FIX: Tambah nil
		return h.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

// Create creates a new system log
func (h *SystemLogHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateSystemLogRequest
	if err := c.BodyParser(&req); err != nil {
		// FIX: Tambah nil
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		// FIX: Tambah nil
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// FIX: h.Success(ctx, data, message)
	return h.Success(c, result, "System log created successfully")
}

// GetByID retrieves a system log by ID
func (h *SystemLogHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		// FIX: Tambah nil
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, result, "System log retrieved successfully")
}

// Delete deletes a system log
func (h *SystemLogHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		// FIX: Tambah nil
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, nil, "System log deleted successfully")
}
