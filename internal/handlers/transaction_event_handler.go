package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
)

type TransactionEventHandler struct {
	*BaseHandler
	service services.TransactionEventService
}

func NewTransactionEventHandler(s services.TransactionEventService) *TransactionEventHandler {
	return &TransactionEventHandler{
		BaseHandler: &BaseHandler{},
		service:     s,
	}
}

// GetAll retrieves all transaction events with pagination
// @Summary Get all transaction events
// @Description Retrieve all transaction events with pagination support
// @Tags Operations
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search term"
// @Success 200 {object} responses.PageResponse
// @Failure 400 {object} responses.BaseResponse
// @Router /api/v1/operations/transaction-events [get]
func (h *TransactionEventHandler) GetAll(c *fiber.Ctx) error {
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

// Create creates a new transaction event
// @Summary Create a new transaction event
// @Description Create a new transaction event record
// @Tags Operations
// @Accept json
// @Produce json
// @Param request body requests.CreateTransactionEventRequest true "Transaction event data"
// @Success 201 {object} responses.BaseResponse
// @Failure 400 {object} responses.BaseResponse
// @Router /api/v1/operations/transaction-events [post]
func (h *TransactionEventHandler) Create(c *fiber.Ctx) error {
	var req requests.CreateTransactionEventRequest
	if err := c.BodyParser(&req); err != nil {
		return h.Error(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	result, err := h.service.Create(c.Context(), req)
	if err != nil {
		return h.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return h.Success(c, result, "Transaction event created successfully")
}

// GetByID retrieves a transaction event by ID
// @Summary Get transaction event by ID
// @Description Retrieve a specific transaction event by its ID
// @Tags Operations
// @Accept json
// @Produce json
// @Param id path string true "Transaction event ID"
// @Success 200 {object} responses.BaseResponse
// @Failure 404 {object} responses.BaseResponse
// @Router /api/v1/operations/transaction-events/{id} [get]
func (h *TransactionEventHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	result, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, result, "Transaction event retrieved successfully")
}

// Delete deletes a transaction event
// @Summary Delete a transaction event
// @Description Delete a specific transaction event by its ID
// @Tags Operations
// @Accept json
// @Produce json
// @Param id path string true "Transaction event ID"
// @Success 200 {object} responses.BaseResponse
// @Failure 404 {object} responses.BaseResponse
// @Router /api/v1/operations/transaction-events/{id} [delete]
func (h *TransactionEventHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.Error(c, fiber.StatusNotFound, err.Error(), nil)
	}

	return h.Success(c, nil, "Transaction event deleted successfully")
}
