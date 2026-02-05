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
// @Router       /vehicles [post]
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
