package templates

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"SkatCRM-Tiny/internal/frontend/templates/components"
	"SkatCRM-Tiny/internal/frontend/templates/icons"
)

var sidebarItems = []components.SidebarItem{
	{
		Name: "Главная",
		URL:  "",
		Icon: icons.HomeIcon,
	},
	{
		Name: "Клиенты",
		URL:  "clients",
		Icon: icons.ClientsIcon,
	},
	{
		Name: "Звонки",
		URL:  "calls",
		Icon: icons.CallsIcon,
	},
	{
		Name: "Заказы",
		URL:  "orders",
		Icon: icons.OrdersIcon,
	},
	{
		Name: "Отчёты",
		URL:  "reports",
		Icon: icons.ReportsIcon,
	},
	{
		Name: "Товары",
		URL:  "products",
		Icon: icons.WarehouseIcon,
	},
}

func Render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
