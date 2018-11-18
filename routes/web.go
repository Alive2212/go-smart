package routes

import "github.com/labstack/echo"

func addWebRoute() {
	RouteWeb.GET("/", func(context echo.Context) error {
		return context.String(200,"successful")
	})
}

