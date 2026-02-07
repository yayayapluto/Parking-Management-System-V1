package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type CustomerHandler struct {
	BaseHandler
	service services.CustomerService
}

func NewCustomerHandler(service services.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Customers
// @Description  Get list of customers with pagination
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name, phone, or rfid_uid"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /customers [get]
func (h *CustomerHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Customer
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateCustomerRequest true "Customer Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /customers [post]
func (h *CustomerHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Customer created successfully")
}

// GetByID godoc
// @Summary      Get Customer By ID
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Customer Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /customers/{id} [get]
func (h *CustomerHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Customer
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Customer Hash ID"
// @Param        request body requests.UpdateCustomerRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /customers/{id} [put]
func (h *CustomerHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Customer updated successfully")
}

// Delete godoc
// @Summary      Delete Customer
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Customer Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /customers/{id} [delete]
func (h *CustomerHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Customer deleted successfully")
}
