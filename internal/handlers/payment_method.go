package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type PaymentMethodHandler struct {
	BaseHandler
	service services.PaymentMethodService
}

func NewPaymentMethodHandler(service services.PaymentMethodService) *PaymentMethodHandler {
	return &PaymentMethodHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Payment Methods
// @Description  Get list of payment methods with pagination
// @Tags         Payment Methods
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /payment-methods [get]
func (h *PaymentMethodHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Payment Method
// @Tags         Payment Methods
// @Accept       json
// @Produce      json
// @Param        request body requests.CreatePaymentMethodRequest true "Payment Method Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /payment-methods [post]
func (h *PaymentMethodHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreatePaymentMethodRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Payment method created successfully")
}

// GetByID godoc
// @Summary      Get Payment Method By ID
// @Tags         Payment Methods
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Payment Method Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /payment-methods/{id} [get]
func (h *PaymentMethodHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Payment Method
// @Tags         Payment Methods
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Payment Method Hash ID"
// @Param        request body requests.UpdatePaymentMethodRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /payment-methods/{id} [put]
func (h *PaymentMethodHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdatePaymentMethodRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Payment method updated successfully")
}

// Delete godoc
// @Summary      Delete Payment Method
// @Tags         Payment Methods
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Payment Method Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /payment-methods/{id} [delete]
func (h *PaymentMethodHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Payment method deleted successfully")
}
