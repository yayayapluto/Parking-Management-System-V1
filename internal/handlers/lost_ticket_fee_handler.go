package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type LostTicketFeeHandler struct {
	BaseHandler
	service services.LostTicketFeeService
}

func NewLostTicketFeeHandler(service services.LostTicketFeeService) *LostTicketFeeHandler {
	return &LostTicketFeeHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Lost Ticket Fees
// @Description  Get list of lost ticket fees with pagination
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /financials/lost-ticket-fees [get]
func (h *LostTicketFeeHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Lost Ticket Fee
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateLostTicketFeeRequest true "Lost Ticket Fee Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /financials/lost-ticket-fees [post]
func (h *LostTicketFeeHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateLostTicketFeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Lost ticket fee created successfully")
}

// GetByID godoc
// @Summary      Get Lost Ticket Fee By ID
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Lost Ticket Fee Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /financials/lost-ticket-fees/{id} [get]
func (h *LostTicketFeeHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Lost ticket fee retrieved successfully")
}

// Delete godoc
// @Summary      Delete Lost Ticket Fee
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Lost Ticket Fee Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /financials/lost-ticket-fees/{id} [delete]
func (h *LostTicketFeeHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}

	return h.Success(ctx, nil, "Lost ticket fee deleted successfully")
}
