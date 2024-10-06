package components

import (
	"net/http"

	"github.com/infraregistry/schema"
	"github.com/infraregistry/schema/models"
	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {
	node := c.Param("id")
	res, err := schema.SelectOnly[models.Component](node)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
