package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type DailySettlementHandler struct {
	*BaseHandler
	service services.DailySettlementService
}

func NewDailySettlementHandler(s services.DailySettlementService) *DailySettlementHandler {
	return &DailySettlementHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all daily settlements with pagination
// @Summary Get all daily settlements
// @Description Retrieve all daily settlements with pagination support
// @Tags Reporting
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search term"
// @Success 200 {object} responses.PageResponse
// @Failure 400 {object} responses.BaseResponse
// @Router /api/v1/admin/reports/daily-settlements [get]
func (h *DailySettlementHandler) GetAll(c *fiber.Ctx) error {
	var req requests.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
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

// Create creates a new daily settlement
// @Summary Create a new daily settlement
// @Description Create a new daily settlement record
// @Tags Reporting
// @Accept json
// @Produce json
// @Param request body requests.CreateDailySettlementRequest true "Daily settlement data"
// @Success 201 {object} responses.BaseResponse
// @Failure 400 {object} responses.BaseResponse
// @Router /api/v1/admin/reports/daily-settlements [post]
func (h *DailySettlementHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateDailySettlementRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return h.Success(c, result, "Daily settlement created successfully")
}

// GetByID retrieves a daily settlement by ID
// @Summary Get daily settlement by ID
// @Description Retrieve a specific daily settlement by its ID
// @Tags Reporting
// @Accept json
// @Produce json
// @Param id path string true "Daily settlement ID"
// @Success 200 {object} responses.BaseResponse
// @Failure 404 {object} responses.BaseResponse
// @Router /api/v1/admin/reports/daily-settlements/{id} [get]
func (h *DailySettlementHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, result, "Daily settlement retrieved successfully")
}

// Reconcile reconciles a daily settlement
// @Summary Reconcile a daily settlement
// @Description Reconcile a specific daily settlement by its ID
// @Tags Reporting
// @Accept json
// @Produce json
// @Param id path string true "Daily settlement ID"
// @Param request body requests.ReconcileDailySettlementRequest true "Reconciliation data"
// @Success 200 {object} responses.BaseResponse
// @Failure 400 {object} responses.BaseResponse
// @Router /api/v1/admin/reports/daily-settlements/{id}/reconcile [post]
func (h *DailySettlementHandler) Reconcile(c *fiber.Ctx) error {
	id := c.Params("id")
	var req requests.ReconcileDailySettlementRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Reconcile(c.Context(), id, req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return h.Success(c, result, "Daily settlement reconciled successfully")
}

// Delete deletes a daily settlement
// @Summary Delete a daily settlement
// @Description Delete a specific daily settlement by its ID
// @Tags Reporting
// @Accept json
// @Produce json
// @Param id path string true "Daily settlement ID"
// @Success 200 {object} responses.BaseResponse
// @Failure 404 {object} responses.BaseResponse
// @Router /api/v1/admin/reports/daily-settlements/{id} [delete]
func (h *DailySettlementHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, nil, "Daily settlement deleted successfully")
}
