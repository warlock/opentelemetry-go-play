package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	apmecho "github.com/opentracing-contrib/echo"
	"github.com/opentracing/opentracing-go"
	"github.com/warlock/opentelemetry-go-play/tracer"
)

const (
	DefaultComponentName = "echo-demo"
)

func main() {
	// 1. init tracer
	tracer, closer := tracer.Init(DefaultComponentName)
	if closer != nil {
		defer closer.Close()
	}
	// 2. ste the global tracer
	if tracer != nil {
		opentracing.SetGlobalTracer(tracer)
	}

	e := echo.New()

	e.Use(apmecho.Middleware(DefaultComponentName))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
