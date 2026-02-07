package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/services"
	"parking-management-system-v1/pkg/helpers"
)

// ParkingTransactionHandler handles HTTP requests related to parking transactions
type ParkingTransactionHandler struct {
	BaseHandler
	service    services.ParkingTransactionService
	obfuscator helpers.IDObfuscator
}

// NewParkingTransactionHandler creates a new instance of ParkingTransactionHandler
func NewParkingTransactionHandler(
	service services.ParkingTransactionService,
	obfuscator helpers.IDObfuscator,
) *ParkingTransactionHandler {
	return &ParkingTransactionHandler{
		service:    service,
		obfuscator: obfuscator,
	}
}

// Entry godoc
// @Summary      Register Vehicle Entry (Gate In)
// @Description  Record a vehicle entering the parking zone and generate a parking ticket
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.ParkingEntryRequest true "Parking Entry Request"
// @Success      201  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Failure      401  {object}  responses.BaseResponse
// @Failure      500  {object}  responses.BaseResponse
// @Router       /transactions/entry [post]
// @Security     Bearer
func (h *ParkingTransactionHandler) Entry(ctx *fiber.Ctx) error {
	// Parse the request body
	var req requests.ParkingEntryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return h.Error(ctx, fiber.StatusBadRequest, "invalid request payload", err.Error())
	}

	// Extract operator ID from JWT context
	// In a real scenario, this would come from the JWT token claims
	// For now, we'll use a default or extract from the request if provided
	var operatorID uint = 1 // TODO: Extract from JWT context

	// If OperatorID is provided in the request, decode it
	if req.OperatorID != "" {
		decodedID, err := h.obfuscator.Decode(req.OperatorID)
		if err != nil {
			return h.Error(ctx, fiber.StatusBadRequest, "invalid operator id", err.Error())
		}
		operatorID = decodedID
	}

	// Call the service to register the entry
	result, err := h.service.RegisterEntry(ctx.Context(), req, operatorID)
	if err != nil {
		return h.Error(ctx, fiber.StatusInternalServerError, "failed to register entry", err.Error())
	}

	// Return success response with 201 Created status
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Vehicle entry recorded successfully",
		"data":    result,
	})
}

// Exit godoc
// @Summary      Register Vehicle Exit (Gate Out)
// @Description  Record a vehicle exiting the parking zone and calculate parking fee
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body requests.ParkingExitRequest true "Parking Exit Request"
// @Success      200  {object}  responses.BaseResponse
// @Failure      400  {object}  responses.BaseResponse
// @Failure      404  {object}  responses.BaseResponse
// @Failure      500  {object}  responses.BaseResponse
// @Router       /transactions/exit [post]
// @Security     Bearer
func (h *ParkingTransactionHandler) Exit(ctx *fiber.Ctx) error {
	// Parse the request body
	var req requests.ParkingExitRequest
	if err := ctx.BodyParser(&req); err != nil {
		return h.Error(ctx, fiber.StatusBadRequest, "invalid request payload", err.Error())
	}

	// Call the service to register the exit
	result, err := h.service.RegisterExit(ctx.Context(), req)
	if err != nil {
		// Check if it's a "not found" error
		if err.Error() == "no active parking found for the given RFID" {
			return h.Error(ctx, fiber.StatusNotFound, "no active parking found", err.Error())
		}
		return h.Error(ctx, fiber.StatusInternalServerError, "failed to register exit", err.Error())
	}

	// Return success response with 200 OK status
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Vehicle exit recorded successfully",
		"data":    result,
	})
}
