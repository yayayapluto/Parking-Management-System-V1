package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type ReportCacheHandler struct {
	*BaseHandler
	service services.ReportCacheService
}

func NewReportCacheHandler(s services.ReportCacheService) *ReportCacheHandler {
	return &ReportCacheHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all report caches with pagination
func (h *ReportCacheHandler) GetAll(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusOK).JSON(result)
}

// Create creates a new report cache
func (h *ReportCacheHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateReportCacheRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// FIX: h.Success(ctx, data, message)
	return h.Success(c, result, "Report cache created successfully")
}

// GetByID retrieves a report cache by ID
func (h *ReportCacheHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, result, "Report cache retrieved successfully")
}

// Update updates a report cache
func (h *ReportCacheHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req requests.UpdateReportCacheRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Update(c.Context(), id, req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return h.Success(c, result, "Report cache updated successfully")
}

// Delete deletes a report cache
func (h *ReportCacheHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, nil, "Report cache deleted successfully")
}
