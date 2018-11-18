package routes

import (
	"../app/Http/Controllers/User"
	"github.com/labstack/echo"
)

func addApiRoute() {
	RouteApiV1 := RouteApi.Group("/v1")
	RouteApiV1.Any("/", func(context echo.Context) error {
		return User.Handler(context)
	})
	RouteApiV1.Any("/:id/", func(context echo.Context) error {
		return User.Handler(context)
	})
	RouteApiV1.Any("/:id/:get_method/", func(context echo.Context) error {
		return User.Handler(context)
	})
}
