package main

import (
	"fmt"
	"os"

	"github.com/jmmesquitacardoso/apitest/internal/handlers"
	"github.com/jmmesquitacardoso/apitest/spec"
	"github.com/labstack/echo/v4"

	middleware "github.com/oapi-codegen/echo-middleware"
)

func main() {
	fmt.Println("main")

	swagger, err := spec.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	e := echo.New()

	// Hide echo banner and port, so we only output valid logs
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.OapiRequestValidator(swagger))

	h := handlers.NewStrict()

	myStrictApiHandler := spec.NewStrictHandler(h, nil)
	spec.RegisterHandlers(e, myStrictApiHandler)

	err = e.Start(fmt.Sprintf("%s:%d", "localhost", 8080))

	fmt.Println("err", err)
}
