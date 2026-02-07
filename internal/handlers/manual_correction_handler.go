package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type ManualCorrectionHandler struct {
	*BaseHandler
	service services.ManualCorrectionService
}

func NewManualCorrectionHandler(s services.ManualCorrectionService) *ManualCorrectionHandler {
	return &ManualCorrectionHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all manual corrections with pagination
// @Summary Get all manual corrections
// @Description Retrieve all manual corrections with pagination support
// @Tags Operations
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search term"
// @Success 200 {object} responses.PageResponse
// @Failure 400 {object} responses.BaseResponse
// @Router /api/v1/operations/manual-corrections [get]
func (h *ManualCorrectionHandler) GetAll(c *fiber.Ctx) error {
	var req requests.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid pagination parameters", nil)
	}

	// Default values logic
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

// Create creates a new manual correction
// @Summary Create a new manual correction
// @Description Create a new manual correction record
// @Tags Operations
// @Accept json
// @Produce json
// @Param request body requests.CreateManualCorrectionRequest true "Manual correction data"
// @Success 201 {object} responses.BaseResponse
// @Failure 400 {object} responses.BaseResponse
// @Router /api/v1/operations/manual-corrections [post]
func (h *ManualCorrectionHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateManualCorrectionRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	// Validate request if your BaseHandler has a validator
	// if err := h.Validate(req); err != nil { return h.Error(c, fiber.StatusBadRequest, err.Error()) }

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// Gunakan helper Success dari BaseHandler agar format JSON konsisten
	return h.Success(c, result, "Manual correction created successfully")
}

// GetByID retrieves a manual correction by ID
// @Summary Get manual correction by ID
// @Description Retrieve a specific manual correction by its ID
// @Tags Operations
// @Accept json
// @Produce json
// @Param id path string true "Manual correction ID"
// @Success 200 {object} responses.BaseResponse
// @Failure 404 {object} responses.BaseResponse
// @Router /api/v1/operations/manual-corrections/{id} [get]
func (h *ManualCorrectionHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.Error(c, fiber.StatusNotFound, "Manual correction not found", nil)
	}

	return h.Success(c, result, "Manual correction retrieved successfully")
}

// Delete deletes a manual correction
// @Summary Delete a manual correction
// @Description Delete a specific manual correction by its ID
// @Tags Operations
// @Accept json
// @Produce json
// @Param id path string true "Manual correction ID"
// @Success 200 {object} responses.BaseResponse
// @Failure 404 {object} responses.BaseResponse
// @Router /api/v1/operations/manual-corrections/{id} [delete]
func (h *ManualCorrectionHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, "Failed to delete: Manual correction not found", nil)
	}

	return h.Success(c, nil, "Manual correction deleted successfully")
}
