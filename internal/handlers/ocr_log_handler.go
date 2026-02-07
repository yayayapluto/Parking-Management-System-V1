package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type OCRLogHandler struct {
	*BaseHandler
	service services.OCRLogService
}

func NewOCRLogHandler(s services.OCRLogService) *OCRLogHandler {
	return &OCRLogHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

func (h *OCRLogHandler) GetAll(c *fiber.Ctx) error {
	var req requests.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		// FIX: Tambah nil di akhir untuk interface{}
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

func (h *OCRLogHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateOCRLogRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// FIX: h.Success want (*fiber.Ctx, interface{}, string)
	return h.Success(c, result, "OCR log created successfully")
}

func (h *OCRLogHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		// FIX: Tambah nil
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	// FIX: Swap position based on your error log (data first, then message)
	return h.Success(c, result, "OCR log retrieved successfully")
}

func (h *OCRLogHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	// FIX: Swap position
	return h.Success(c, nil, "OCR log deleted successfully")
}
