package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"SkatCRM-Tiny/internal/frontend"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fileServer := http.FileServer(http.FS(frontend.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/", echo.HandlerFunc(s.HelloWorldHandler))
	e.GET("/websocket", echo.HandlerFunc(s.websocketHandler))

	// TODO: Client
	e.GET("/view/clients", echo.HandlerFunc(s.GetClientsHandler))
	// e.GET("/api/v1/clients", nil)
	// e.GET("/api/v1/clients/:count/:offset", nil)
	// e.GET("/api/v1/client/:id", nil)

	// e.POST("/api/v1/client", nil)
	// e.DELETE("/api/v1/client/:id", nil)

	return e
}
