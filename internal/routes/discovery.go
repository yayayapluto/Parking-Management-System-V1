package routes

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func DiscoveryHandler(fb *fiber.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		p := c.Path()
		if p != "/" && p[len(p)-1] == '/' {
			p = p[:len(p)-1]
		}

		// Hitung Level
		level := 0
		if p != "/" {
			level = len(strings.Split(strings.Trim(p, "/"), "/"))
		}

		routes := fb.Stack()
		type RouteInfo struct {
			Path   string `json:"path"`
			Method string `json:"method"`
		}

		var children []RouteInfo
		for _, routeStack := range routes {
			for _, r := range routeStack {
				if r.Method == "HEAD" || strings.Contains(r.Path, "metrics") {
					continue
				}

				if strings.HasPrefix(r.Path, p) && r.Path != p {
					childPath := strings.TrimPrefix(r.Path, p)
					childParts := strings.Split(strings.Trim(childPath, "/"), "/")

					if len(childParts) == 1 {
						children = append(children, RouteInfo{
							Path:   r.Path,
							Method: r.Method,
						})
					}
				}
			}
		}

		return c.JSON(fiber.Map{
			"current_level": level,
			"current_path":  p,
			"sub_routes":    children,
		})
	}
}
