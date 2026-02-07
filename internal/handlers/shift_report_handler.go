package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type ShiftReportHandler struct {
	*BaseHandler
	service services.ShiftReportService
}

func NewShiftReportHandler(s services.ShiftReportService) *ShiftReportHandler {
	return &ShiftReportHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all shift reports with pagination
func (h *ShiftReportHandler) GetAll(c *fiber.Ctx) error {
	var req requests.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		// FIX: Tambah nil di argumen ke-4 sesuai kontrak BaseHandler
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

	// Menggunakan c.Status().JSON() agar konsisten dengan handler lain yang sudah di-fix
	return c.Status(fiber.StatusOK).JSON(result)
}

// Create creates a new shift report
func (h *ShiftReportHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateShiftReportRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return h.Success(c, result, "Shift report created successfully")
}

// GetByID retrieves a shift report by ID
func (h *ShiftReportHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, result, "Shift report retrieved successfully")
}

// Update updates a shift report
func (h *ShiftReportHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req requests.UpdateShiftReportRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Update(c.Context(), id, req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return h.Success(c, result, "Shift report updated successfully")
}

// Delete deletes a shift report
func (h *ShiftReportHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, nil, "Shift report deleted successfully")
}
