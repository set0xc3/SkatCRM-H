package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"SkatCRM-Tiny/internal/frontend"
	"SkatCRM-Tiny/internal/frontend/templates"
	"SkatCRM-Tiny/internal/frontend/templates/views"
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
	e.GET("/clients", func(c echo.Context) error {
		return templates.Render(c, templates.LayoutTempl(views.ClientsTempl()))
	})
	e.GET("/view/clients", func(c echo.Context) error {
		return templates.Render(c, views.ClientsTempl())
	})
	// e.GET("/api/v1/clients", nil)
	e.GET("/api/v1/clients/:count/:offset", func(c echo.Context) error {
		count, _ := strconv.Atoi(c.Param("count"))
		offset, _ := strconv.Atoi(c.Param("offset"))

		clients, _ := s.db.GetClientInstance().GetClients(count, offset)
		return c.JSON(http.StatusOK, clients)
	})
	// e.GET("/api/v1/client/:id", nil)

	// e.POST("/api/v1/client", nil)
	// e.DELETE("/api/v1/client/:id", nil)

	return e
}
