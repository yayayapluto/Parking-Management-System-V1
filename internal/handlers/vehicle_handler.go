package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type VehicleHandler struct {
	BaseHandler
	service services.VehicleService
}

func NewVehicleHandler(service services.VehicleService) *VehicleHandler {
	return &VehicleHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Vehicles
// @Description  Get list of vehicles with pagination
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by plate_number or description"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /vehicles [get]
func (h *VehicleHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Vehicle
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateVehicleRequest true "Vehicle Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /vehicles [post]
func (h *VehicleHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateVehicleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Vehicle created successfully")
}

// GetByID godoc
// @Summary      Get Vehicle By ID
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vehicle Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /vehicles/{id} [get]
func (h *VehicleHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Vehicle
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Vehicle Hash ID"
// @Param        request body requests.UpdateVehicleRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /vehicles/{id} [put]
func (h *VehicleHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateVehicleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Vehicle updated successfully")
}

// Delete godoc
// @Summary      Delete Vehicle
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vehicle Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /vehicles/{id} [delete]
func (h *VehicleHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Vehicle deleted successfully")
}
