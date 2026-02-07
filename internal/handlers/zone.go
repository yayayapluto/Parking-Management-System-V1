package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type ZoneHandler struct {
	BaseHandler
	service services.ZoneService
}

func NewZoneHandler(service services.ZoneService) *ZoneHandler {
	return &ZoneHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Zones
// @Description  Get list of zones with pagination
// @Tags         Zones
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /zones [get]
func (h *ZoneHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Zone
// @Tags         Zones
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateZoneRequest true "Zone Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zones [post]
func (h *ZoneHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateZoneRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Zone created successfully")
}

// GetByID godoc
// @Summary      Get Zone By ID
// @Tags         Zones
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Zone Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /zones/{id} [get]
func (h *ZoneHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Zone
// @Tags         Zones
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Zone Hash ID"
// @Param        request body requests.UpdateZoneRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zones/{id} [put]
func (h *ZoneHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateZoneRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Zone updated successfully")
}

// Delete godoc
// @Summary      Delete Zone
// @Tags         Zones
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Zone Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zones/{id} [delete]
func (h *ZoneHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Zone deleted successfully")
}
