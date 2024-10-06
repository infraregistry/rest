package components

import (
	"github.com/infraregistry/rest/middleware"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) *echo.Group {
	g := e.Group("/components", middleware.SessionMiddleware())

	g.GET("", Search)
	g.GET("/:id", Get)
	g.GET("/graph", Graph)

	return g
}
