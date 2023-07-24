package route

import (
	"net/http"

	"github.com/koskuma/go-worker-workshop/config"
	"github.com/labstack/echo/v4"

	"github.com/koskuma/go-worker-workshop/cmd/api/handler"
	"github.com/koskuma/go-worker-workshop/cmd/api/middleware"

	_ "github.com/koskuma/go-worker-workshop/cmd/api/docs" // docs is generated by Swag CLI, you have to import it.
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRoute(config config.Config, app *echo.Echo, userHandler handler.IUser, probeHandler handler.IProbe) {
	app.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello world")
	})
	app.GET("/readyz", probeHandler.DBReadyCheck)

	u := app.Group("/user")

	u.Use(middleware.NewVerifyJWTAuth([]byte(config.JWTSecret), config.JWTSigningMethod))
	u.Use(middleware.ExtractJWTClaims)

	u.GET("", userHandler.GetAll)
	u.GET("/", userHandler.GetUser)
	u.POST("/", userHandler.Create)

	app.GET("/swagger/*", echoSwagger.WrapHandler)
}
