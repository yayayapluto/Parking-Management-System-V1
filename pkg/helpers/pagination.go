package helpers

import (
	"github.com/gofiber/fiber/v2"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/pkg/logger"
)

func ParsePaginationRequest(ctx *fiber.Ctx) requests.PaginationRequest {
	req := requests.PaginationRequest{
		Page:  1,
		Limit: 10,
	}

	if err := ctx.QueryParser(&req); err != nil {
		logger.Error(err.Error(), "helpers-ParsePaginationRequest")
	}

	return req
}
