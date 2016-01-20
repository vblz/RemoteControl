package interfaces

import "net/http"

// BaseEndPointInfo base interface EndPointInfo implemintation
type BaseEndPointInfo struct {
	CurrentPath    string
	CurrentHandler http.HandlerFunc
	CurrentType    EndPointType
}

// Path return setted path of current instance
func (ep BaseEndPointInfo) Path() string {
	return ep.CurrentPath
}

// Handler return setted handler of current instance
func (ep BaseEndPointInfo) Handler() http.HandlerFunc {
	return ep.CurrentHandler
}

// Type return setted type of EndPoin of current instance
func (ep BaseEndPointInfo) Type() EndPointType {
	return ep.CurrentType
}
