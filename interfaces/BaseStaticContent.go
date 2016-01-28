package interfaces

import "html/template"

// BaseStaticContent base interface StaticContent implemintation
type BaseStaticContent struct {
	CurrentPath  template.URL
	CurrentData  template.HTML
	CurrentTitle string
}

// Path return setted path of current instance
func (sc BaseStaticContent) Path() template.URL {
	return sc.CurrentPath
}

// Data return setted HTML data of current instance
func (sc BaseStaticContent) Data() template.HTML {
	return sc.CurrentData
}

// Title return setted title of current instance
func (sc BaseStaticContent) Title() string {
	return sc.CurrentTitle
}
