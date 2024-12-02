package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// #include "demo.h"
	"C"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)
	e.GET("/clang/:name", s.clangHandler)
	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) clangHandler(c echo.Context) error {
	name := c.Param("name")
	result := C.GoString(C.greet(C.CString(name)))
	resp := map[string]string{
		"response": result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World v1",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
