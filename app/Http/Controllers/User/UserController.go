package User

import (
	"../../../../app"
	"../../../../database"
	"github.com/alive2212/go-illuminate/response"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"strconv"
)

// must be created by programmer
type model *app.User

// must be created by programmer
var sampleModel = app.User{Name: "string", PhoneNumber: 9120000000, CountryCode: "+98"}

var models []model

/**
Handler Controller
 */
func Handler(ctx echo.Context) error {
	requestMethod := ctx.Request().Method
	switch requestMethod {
	case "GET":
		if getRequestGetMethod(ctx) == "edit" {
			return Edit(ctx)
		}
		if hasRequestId(ctx) {
			if getRequestId(ctx) == "create" {
				return Create(ctx)
			}
			return Show(ctx)
			break
		}
		return Index(ctx)
		break
	case "POST":
		// store something into table
		return Store(ctx)
		break
	case "PUT":
		// update something into table
		return Update(ctx)
		break
	case "PATCH":
		return Update(ctx)
		break
	case "DELETE":
		return Destroy(ctx)
		break
	case "OPTIONS":
		break
	default:
		break
	}
	return ctx.JSON(200, map[string]string{
		"message": "unknown type request",
	})
}

func hasRequestGetMethod(ctx echo.Context) bool {
	return ctx.Param("get_method") != ""
}

func hasRequestId(ctx echo.Context) bool {
	return ctx.Param("id") != ""

}

func getRequestId(ctx echo.Context) string {
	return ctx.Param("id")

}

func getRequestGetMethod(ctx echo.Context) string {
	return ctx.Param("get_method")
}

// store record
func Store(ctx echo.Context) error {
	database.DB.AutoMigrate(&app.User{})

	params := getFormParams(ctx)
	phoneNumber,_ := strconv.ParseInt(params["phone_number"],10,64)
	//var phoneNumberInt int64 = phoneNumber
	database.DB.Create(app.User{
		Name:        params["name"],
		PhoneNumber: phoneNumber ,
		CountryCode: params["country_code"],
	})
	return response.SuccessfulResponse(ctx, response.SuccessfulResponseModel{
		Data:       models,
		MessageKey: "Successful",
	})
}

// get all records with dependencies
func Index(ctx echo.Context) error {
	database.DB.Find(&models)
	return ctx.JSON(200, models)
}

// update record
func Update(ctx echo.Context) (err error) {
	updateParams := getFormParams(ctx)
	database.DB.Find(&models, getRequestId(ctx))
	database.DB.Model(&models).Update(updateParams)
	return response.SuccessfulResponse(ctx, response.SuccessfulResponseModel{
		Data:       updateParams,
		MessageKey: "successful",
	})
}

func getFormParams(ctx echo.Context) map[string]string {
	params, _ := ctx.FormParams()
	updateParams := make(map[string]string)
	for key, values := range params {
		updateParams[key] = values[0]
	}
	return updateParams
}

// remove record
func Destroy(ctx echo.Context) error {
	database.DB.Find(&models, getRequestId(ctx))
	database.DB.Delete(models)
	return response.SuccessfulResponse(ctx, response.SuccessfulResponseModel{
		Data:       models,
		MessageKey: "successful",
	})
}

// response record with all dependencies
func Edit(ctx echo.Context) error {
	database.DB.Find(&models, getRequestId(ctx))
	return response.SuccessfulResponse(ctx, response.SuccessfulResponseModel{
		Data:       models,
		MessageKey: "successful",
	})
}

// response record in abstract mode
func Show(ctx echo.Context) error {
	database.DB.Find(&models, getRequestId(ctx))
	return response.SuccessfulResponse(ctx, response.SuccessfulResponseModel{
		Data:       models,
		MessageKey: "successful",
	})
}

// get all parameter of models to create it
func Create(ctx echo.Context) error {
	return response.SuccessfulResponse(ctx, response.SuccessfulResponseModel{
		Data:       sampleModel,
		MessageKey: "successful",
	})
}
