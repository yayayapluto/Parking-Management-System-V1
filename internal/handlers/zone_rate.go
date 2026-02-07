package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type ZoneRateHandler struct {
	BaseHandler
	service services.ZoneRateService
}

func NewZoneRateHandler(service services.ZoneRateService) *ZoneRateHandler {
	return &ZoneRateHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Zone Rates
// @Description  Get list of zone rates with pagination
// @Tags         Zone Rates
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /zone-rates [get]
func (h *ZoneRateHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Zone Rate
// @Tags         Zone Rates
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateZoneRateRequest true "Zone Rate Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zone-rates [post]
func (h *ZoneRateHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateZoneRateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Zone rate created successfully")
}

// GetByID godoc
// @Summary      Get Zone Rate By ID
// @Tags         Zone Rates
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Zone Rate Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /zone-rates/{id} [get]
func (h *ZoneRateHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Zone Rate
// @Tags         Zone Rates
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Zone Rate Hash ID"
// @Param        request body requests.UpdateZoneRateRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zone-rates/{id} [put]
func (h *ZoneRateHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateZoneRateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Zone rate updated successfully")
}

// Delete godoc
// @Summary      Delete Zone Rate
// @Tags         Zone Rates
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Zone Rate Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zone-rates/{id} [delete]
func (h *ZoneRateHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Zone rate deleted successfully")
}
