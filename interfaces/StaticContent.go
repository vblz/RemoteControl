package interfaces

import "html/template"

// StaticContent interface for use for main pages
type StaticContent interface {
	Path() template.URL
	Data() template.HTML
	Title() string
}
