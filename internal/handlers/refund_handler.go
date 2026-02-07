package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type RefundHandler struct {
	BaseHandler
	service services.RefundService
}

func NewRefundHandler(service services.RefundService) *RefundHandler {
	return &RefundHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Refunds
// @Description  Get list of refunds with pagination
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by reason"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /financials/refunds [get]
func (h *RefundHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Refund
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateRefundRequest true "Refund Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /financials/refunds [post]
func (h *RefundHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateRefundRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Refund created successfully")
}

// GetByID godoc
// @Summary      Get Refund By ID
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Refund Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /financials/refunds/{id} [get]
func (h *RefundHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Refund retrieved successfully")
}

// Update godoc
// @Summary      Update Refund
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        id      path      string                        true  "Refund Hash ID"
// @Param        request body requests.UpdateRefundRequest true "Update Refund Request"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /financials/refunds/{id} [put]
func (h *RefundHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateRefundRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Refund updated successfully")
}

// Approve godoc
// @Summary      Approve Refund
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        id      path      string                        true  "Refund Hash ID"
// @Param        request body requests.ApproveRefundRequest true "Approve Refund Request"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /financials/refunds/{id}/approve [post]
func (h *RefundHandler) Approve(ctx *fiber.Ctx) error {
	var req requests.ApproveRefundRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Approve(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Refund approved successfully")
}

// Reject godoc
// @Summary      Reject Refund
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        id      path      string                       true  "Refund Hash ID"
// @Param        request body requests.RejectRefundRequest true "Reject Refund Request"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /financials/refunds/{id}/reject [post]
func (h *RefundHandler) Reject(ctx *fiber.Ctx) error {
	var req requests.RejectRefundRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Reject(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Refund rejected successfully")
}

// Delete godoc
// @Summary      Delete Refund
// @Tags         Financials
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Refund Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /financials/refunds/{id} [delete]
func (h *RefundHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}

	return h.Success(ctx, nil, "Refund deleted successfully")
}
