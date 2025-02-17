package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"SkatCRM-Tiny/internal/backend/database"
	"SkatCRM-Tiny/internal/backend/database/entities"
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

	e.GET("/", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, templates.LayoutTempl(c.Request().URL.Path, views.SandboxTempl()))
	})
	e.GET("/clients", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, templates.LayoutTempl(c.Request().URL.Path, views.ClientsTempl()))
	})
	e.GET("/calls", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, templates.LayoutTempl(c.Request().URL.Path, views.EmptyTempl()))
	})
	e.GET("/orders", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, templates.LayoutTempl(c.Request().URL.Path, views.EmptyTempl()))
	})
	e.GET("/reports", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, templates.LayoutTempl(c.Request().URL.Path, views.EmptyTempl()))
	})
	e.GET("/products", func(c echo.Context) error {
		return templates.Render(c, templates.LayoutTempl(c.Request().URL.Path, views.EmptyTempl()))
	})
	e.GET("/views/", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, views.SandboxTempl())
	})
	e.GET("/views/clients", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, views.ClientsTempl())
	})
	e.GET("/views/calls", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, views.EmptyTempl())
	})
	e.GET("/views/orders", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, views.EmptyTempl())
	})
	e.GET("/views/reports", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, views.EmptyTempl())
	})
	e.GET("/views/products", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return templates.Render(c, views.EmptyTempl())
	})
	e.GET("/lazy-load", func(c echo.Context) error {
		// time.Sleep(1 * time.Second)
		return c.JSON(http.StatusOK, "")
	})
	e.GET("/sandbox", func(c echo.Context) error {
		return templates.Render(c, templates.LayoutTempl(c.Request().URL.Path, views.SandboxTempl()))
	})

	e.GET("/api/v1/clients/:count/:offset", func(c echo.Context) error {
		count, _ := strconv.Atoi(c.Param("count"))
		offset, _ := strconv.Atoi(c.Param("offset"))

		clients, _ := database.GetInstance().GetClients().FetchClients(count, offset)
		return c.JSON(http.StatusOK, clients)
	})
	// e.GET("/api/v1/client/:id", nil)

	e.POST("/api/v1/client", func(c echo.Context) error {
		client := new(entities.ClientInfo)
		if err := c.Bind(client); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		client.Id = uuid.NewString()
		err := database.GetInstance().GetClients().InsertClient(*client)
		if err != nil {
			return fmt.Errorf("Failed insert client: %s", err)
		}

		return templates.Render(c, views.ClientsTempl())
	})
	e.DELETE("/api/v1/client/:id", func(c echo.Context) error {
		id := c.Param("id")

		err := database.GetInstance().GetClients().DeleteClient(id)
		if err != nil {
			return fmt.Errorf("Failed delete client: %s", err)
		}

		return c.NoContent(http.StatusOK)
	})

	return e
}
