package app

import (
	"chat-app-golang/internal/endpoint"
	"chat-app-golang/internal/service"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	echo *echo.Echo

	testEndpoint *endpoint.Test
	testSrv      *service.Test
}

func New(mongoConnection string) (*App, error) {
	a := &App{}

	a.testSrv = service.NewTest()
	a.testEndpoint = endpoint.NewTest(a.testSrv)

	a.echo = echo.New()

	a.echo.GET("/test", a.testEndpoint.GetHello)

	return a, nil
}

func (a App) Run() error {

	if err := a.echo.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	return nil
}
