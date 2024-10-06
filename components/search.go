package components

import (
	"net/http"

	"github.com/infraregistry/schema"
	"github.com/infraregistry/schema/models"
	"github.com/labstack/echo/v4"
)

func Search(c echo.Context) error {
	res, err := schema.Select[models.Component]("component")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
