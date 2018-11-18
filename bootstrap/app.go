package bootstrap

import (
	"../config"
	"github.com/labstack/echo"
)


var App *echo.Echo

//type MyRouter struct {
//	*echo.Group
//}
//type CustomValidator struct {
//	validator *echo.Validator
//}

//func (cv *CustomValidator) Validate(i interface{}) error {
//	return cv.validator.Struct(i)
//}

func New() *echo.Echo {
	//App := route.Echo


	App = echo.New()
	//App.Validator  = &CustomValidator{validator: validator.New()}
	//App = echo.New()
	//
	//Route := MyRouter{}

	return App
}



func Start() {
	App.Start(":"+config.ExposePort)
}