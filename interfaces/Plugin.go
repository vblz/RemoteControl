package interfaces

type Plugin interface {
	GetHandlers() []EndPointInfo
	GetMainContent() []StaticContent
}
