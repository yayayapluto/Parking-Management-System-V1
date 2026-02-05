package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type ZoneTypeHandler struct {
	BaseHandler
	service services.ZoneTypeService
}

func NewZoneTypeHandler(service services.ZoneTypeService) *ZoneTypeHandler {
	return &ZoneTypeHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Zone Types
// @Description  Get list of zone types with pagination
// @Tags         Zone Types
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /zone-types [get]
func (h *ZoneTypeHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Zone Type
// @Tags         Zone Types
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateZoneTypeRequest true "Zone Type Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zone-types [post]
func (h *ZoneTypeHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateZoneTypeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Zone type created successfully")
}

// GetByID godoc
// @Summary      Get Zone Type By ID
// @Tags         Zone Types
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Zone Type Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /zone-types/{id} [get]
func (h *ZoneTypeHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Zone Type
// @Tags         Zone Types
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Zone Type Hash ID"
// @Param        request body requests.UpdateZoneTypeRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zone-types/{id} [put]
func (h *ZoneTypeHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateZoneTypeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Zone type updated successfully")
}

// Delete godoc
// @Summary      Delete Zone Type
// @Tags         Zone Types
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Zone Type Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /zone-types/{id} [delete]
func (h *ZoneTypeHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Zone type deleted successfully")
}
