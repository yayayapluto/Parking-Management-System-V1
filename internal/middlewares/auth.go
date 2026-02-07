package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"parking-management-system-v1/pkg/config"
	"strings"
)

// JWTMiddleware validates JWT tokens from Authorization header
func JWTMiddleware(cfg *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get Authorization header
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "missing authorization header",
			})
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "invalid authorization header format",
			})
		}

		tokenString := parts[1]

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(cfg.JWTConfig.Secret), nil
		})

		if err != nil || !token.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "invalid or expired token",
			})
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "invalid token claims",
			})
		}

		// Store claims in context for later use
		ctx.Locals("user_id", claims["user_id"])
		ctx.Locals("username", claims["username"])
		ctx.Locals("email", claims["email"])
		ctx.Locals("role_id", claims["role_id"])
		ctx.Locals("role_name", claims["role_name"])
		ctx.Locals("permissions", claims["permissions"])

		return ctx.Next()
	}
}

// OptionalJWTMiddleware validates JWT tokens but doesn't fail if missing
func OptionalJWTMiddleware(cfg *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.Next()
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.Next()
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(cfg.JWTConfig.Secret), nil
		})

		if err != nil || !token.Valid {
			return ctx.Next()
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx.Next()
		}

		// Store claims in context
		ctx.Locals("user_id", claims["user_id"])
		ctx.Locals("username", claims["username"])
		ctx.Locals("email", claims["email"])
		ctx.Locals("role_id", claims["role_id"])
		ctx.Locals("role_name", claims["role_name"])
		ctx.Locals("permissions", claims["permissions"])

		return ctx.Next()
	}
}

// RoleMiddleware checks if user has required role
func RoleMiddleware(requiredRoles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		roleName := ctx.Locals("role_name")
		if roleName == nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "user role not found",
			})
		}

		userRole := roleName.(string)
		for _, role := range requiredRoles {
			if userRole == role {
				return ctx.Next()
			}
		}

		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "insufficient permissions",
		})
	}
}

// PermissionMiddleware checks if user has required permission
func PermissionMiddleware(requiredPermissions ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		permissions := ctx.Locals("permissions")
		if permissions == nil {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "user permissions not found",
			})
		}

		userPermissions, ok := permissions.([]interface{})
		if !ok {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "invalid permissions format",
			})
		}

		for _, required := range requiredPermissions {
			found := false
			for _, perm := range userPermissions {
				if perm.(string) == required {
					found = true
					break
				}
			}
			if !found {
				return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"success": false,
					"message": "insufficient permissions",
				})
			}
		}

		return ctx.Next()
	}
}
