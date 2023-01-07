package app

import (
	"chat-app-golang/internal/controller"
	"chat-app-golang/internal/db/mongoreposytory"
	"chat-app-golang/internal/service"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	echo *echo.Echo

	testController *controller.Test
	authController *controller.Auth

	testSrv *service.Test
}

func New() (*App, error) {
	a := &App{}

	a.testController =
		controller.NewTest(
			service.NewTest(
				mongoreposytory.User{}))

	a.authController =
		controller.NewAuth(
			service.NewAuth(
				mongoreposytory.User{}))

	a.echo = echo.New()

	a.echo.GET("api/test", a.testController.GetHello)

	a.echo.POST("api/registration", a.authController.Registration)
	//a.echo.POST("api/login", a.authController.Login)

	return a, nil
}

func (a App) Run() error {

	if err := a.echo.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	return nil
}
