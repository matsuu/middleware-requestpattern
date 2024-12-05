// Package echo is middleware for echo framework
package echo

import (
	"fmt"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

// RequestPatternHeader is for Response header
const RequestPatternHeader = "X-Request-Pattern"

// Middleware returns middleware to set X-Request-Pattern response header.
func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			var b strings.Builder
			path := c.Path()
			if path == "" {
				path = req.URL.Path
			}
			b.WriteString(path)
			query := c.QueryParams()
			if len(query) > 0 {
				b.WriteRune('?')
				keys := make([]string, 0, len(query))
				for k := range query {
					chunk := fmt.Sprintf("%s=*", k)
					keys = append(keys, chunk)
				}
				sort.Strings(keys)
				b.WriteString(strings.Join(keys, "&"))
			}
			c.Response().Header().Set(RequestPatternHeader, b.String())
			return next(c)
		}
	}
}
