package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

type VehicleTypeHandler struct {
	BaseHandler
	service services.VehicleTypeService
}

func NewVehicleTypeHandler(service services.VehicleTypeService) *VehicleTypeHandler {
	return &VehicleTypeHandler{service: service}
}

// GetAll godoc
// @Summary      Get All Vehicle Types
// @Description  Get list of vehicle types with pagination
// @Tags         Vehicle Types
// @Accept       json
// @Produce      json
// @Param        page      query    int     false  "Page number"    default(1)
// @Param        limit     query    int     false  "Items per page" default(10)
// @Param        search    query    string  false  "Search by name"
// @Success      200      {object}  responses.PageResponse
// @Failure      400      {object}  responses.BaseResponse
// @Router       /vehicle-types [get]
func (h *VehicleTypeHandler) GetAll(ctx *fiber.Ctx) error {
	req := helpers.ParsePaginationRequest(ctx)
	url := ctx.Protocol() + "://" + ctx.Hostname() + ctx.Path()

	results, err := h.service.GetAll(ctx.Context(), req, url)
	if err != nil {
		return err
	}

	return h.Pagination(ctx, results)
}

// Create godoc
// @Summary      Create New Vehicle Type
// @Tags         Vehicle Types
// @Accept       json
// @Produce      json
// @Param        request body requests.CreateVehicleTypeRequest true "Vehicle Type Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /vehicle-types [post]
func (h *VehicleTypeHandler) Create(ctx *fiber.Ctx) error {
	var req requests.CreateVehicleTypeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Vehicle type created successfully")
}

// GetByID godoc
// @Summary      Get Vehicle Type By ID
// @Tags         Vehicle Types
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vehicle Type Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Router       /vehicle-types/{id} [get]
func (h *VehicleTypeHandler) GetByID(ctx *fiber.Ctx) error {
	res, err := h.service.GetByID(ctx.Context(), ctx.Params("id"))
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Fetch By ID successfully")
}

// Update godoc
// @Summary      Update Vehicle Type
// @Tags         Vehicle Types
// @Accept       json
// @Produce      json
// @Param        id      path      string  true  "Vehicle Type Hash ID"
// @Param        request body requests.UpdateVehicleTypeRequest true "Update Request Body"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /vehicle-types/{id} [put]
func (h *VehicleTypeHandler) Update(ctx *fiber.Ctx) error {
	var req requests.UpdateVehicleTypeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.service.Update(ctx.Context(), ctx.Params("id"), req)
	if err != nil {
		return err
	}

	return h.Success(ctx, res, "Vehicle type updated successfully")
}

// Delete godoc
// @Summary      Delete Vehicle Type
// @Tags         Vehicle Types
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Vehicle Type Hash ID"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Router       /vehicle-types/{id} [delete]
func (h *VehicleTypeHandler) Delete(ctx *fiber.Ctx) error {
	if err := h.service.Delete(ctx.Context(), ctx.Params("id")); err != nil {
		return err
	}
	return h.Success(ctx, nil, "Vehicle type deleted successfully")
}
