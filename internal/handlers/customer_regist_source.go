package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type CustomerRegistSourceHandler struct {
	BaseHandler
	service services.CustomerRegistSourceService
}

func NewCustomerRegistSourceHandler(service services.CustomerRegistSourceService) *CustomerRegistSourceHandler {
	return &CustomerRegistSourceHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Customer Registration Sources
// @Description  Get list of customer registration sources with pagination
// @Tags         Customer Registration Sources
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /customer-regist-sources [get]
func (h *CustomerRegistSourceHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Customer Registration Source
// @Tags         Customer Registration Sources
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateCustomerRegistSourceRequest true "Customer Registration Source Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /customer-regist-sources [post]
func (h *CustomerRegistSourceHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateCustomerRegistSourceRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Customer registration source created successfully")
}

// GetByID godoc
// @Summary      Get Customer Registration Source By ID
// @Tags         Customer Registration Sources
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Customer Registration Source Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /customer-regist-sources/{id} [get]
func (h *CustomerRegistSourceHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Customer Registration Source
// @Tags         Customer Registration Sources
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Customer Registration Source Hash ID"
// @Param        request body requests.UpdateCustomerRegistSourceRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /customer-regist-sources/{id} [put]
func (h *CustomerRegistSourceHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateCustomerRegistSourceRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Customer registration source updated successfully")
}

// Delete godoc
// @Summary      Delete Customer Registration Source
// @Tags         Customer Registration Sources
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Customer Registration Source Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /customer-regist-sources/{id} [delete]
func (h *CustomerRegistSourceHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Customer registration source deleted successfully")
}
