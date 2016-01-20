package interfaces

import "net/http"

// EndPointType - type of EndPointInfo
// Means that current EndPoint use for api or content
type EndPointType int

const (
	// EndPointContent means use current EndPoint for contenet
	EndPointContent EndPointType = iota
	// EndPointAPI means use current EndPoint for API
	EndPointAPI
)

// EndPointInfo interface for use during plugins registering
type EndPointInfo interface {
	Path() string
	Handler() http.HandlerFunc
	Type() EndPointType
}
