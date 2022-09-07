package routes

import (
	"echo-rest/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/user", controllers.FetchAllPegawai)
	e.POST("/user", controllers.StorePegawai)
	e.PUT("/user", controllers.UpdateUser)
	e.DELETE("/user", controllers.DeleteUser)

	return e
}
