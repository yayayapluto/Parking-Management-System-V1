package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type HolidayHandler struct {
	BaseHandler
	service services.HolidayService
}

func NewHolidayHandler(service services.HolidayService) *HolidayHandler {
	return &HolidayHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Holidays
// @Description  Get list of holidays with pagination
// @Tags         Holidays
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /holidays [get]
func (h *HolidayHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Holiday
// @Tags         Holidays
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateHolidayRequest true "Holiday Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /holidays [post]
func (h *HolidayHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateHolidayRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Holiday created successfully")
}

// GetByID godoc
// @Summary      Get Holiday By ID
// @Tags         Holidays
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Holiday Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /holidays/{id} [get]
func (h *HolidayHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Holiday
// @Tags         Holidays
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Holiday Hash ID"
// @Param        request body requests.UpdateHolidayRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /holidays/{id} [put]
func (h *HolidayHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateHolidayRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Holiday updated successfully")
}

// Delete godoc
// @Summary      Delete Holiday
// @Tags         Holidays
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Holiday Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /holidays/{id} [delete]
func (h *HolidayHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Holiday deleted successfully")
}
