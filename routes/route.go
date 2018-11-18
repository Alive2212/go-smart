package routes

import (
	"../bootstrap"
	"github.com/labstack/echo"
)

var RouteApi *echo.Group
var RouteWeb *echo.Group

//var RouteApi1 *route.Illuminate

func New() {

	RouteApi = bootstrap.App.Group("/api")
	addApiRoute()

	//RouteApi1 = route.Illuminate{RouteApi}


	RouteWeb = bootstrap.App.Group("/web")
	addWebRoute()
}
