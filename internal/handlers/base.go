package handlers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/responses"
)

type BaseHandler struct{}

// Success - Untuk response standard (Create, Update, Delete, Single Detail)
func (h *BaseHandler) Success(ctx *fiber.Ctx, data interface{}, message string) error {
	return ctx.Status(fiber.StatusOK).JSON(responses.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error - Untuk custom error response
func (h *BaseHandler) Error(ctx *fiber.Ctx, status int, message string, errors interface{}) error {
	return ctx.Status(status).JSON(responses.BaseResponse{
		Success: false,
		Message: message,
		Errors:  errors, // Pastikan di struct BaseResponse kamu ada field Errors
	})
}

// Pagination - Khusus untuk list data (History transaksi, daftar user, dll)
func (h *BaseHandler) Pagination(ctx *fiber.Ctx, response responses.PageResponse) error {
	return ctx.Status(fiber.StatusOK).JSON(responses.PageResponse{
		Success:    true,
		Message:    response.Message,
		Data:       response.Data,
		Pagination: response.Pagination,
	})
}
