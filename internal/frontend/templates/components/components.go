package components

import (
	"github.com/a-h/templ"
)

type Checkbox struct {
	Label   string
	Name    string
	Checked bool
	Class   string
	Attrs   templ.Attributes
}

type SidebarItem struct {
	Name string
	URL  string
	Icon string
}

type Sidebar struct {
	Items []SidebarItem
}
