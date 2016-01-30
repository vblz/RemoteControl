package keyboard

import (
	// #include <wtypes.h>
	// #include <Winuser.h>
	"C"
	"html/template"
	"net/http"
	"strconv"

	"github.com/vblazhnov/RemoteControl/interfaces"

	"github.com/vblazhnov/RemoteControl/utils"
)

var (
	pageContent template.HTML
)

func init() {
	pageContent = utils.ReadHTML("\\static\\plugins\\keyboard\\index.html")
}

// Control allow to remote use mouse
type Control struct {
}

// GetHandlers implemintation of interfaces.EndPointInfo
func (c Control) GetHandlers() []interfaces.EndPointInfo {
	handlers := []interfaces.EndPointInfo{
		interfaces.BaseEndPointInfo{
			"/key",
			keyServerRequest,
			interfaces.EndPointAPI},
		interfaces.BaseEndPointInfo{
			"/keyboard",
			contentServeRequest,
			interfaces.EndPointContent},
	}

	return handlers
}

// GetMainContent implemintation of interfaces.Plugin
func (c Control) GetMainContent() []interfaces.StaticContent {
	statics := []interfaces.StaticContent{}

	return statics
}

func contentServeRequest(
	w http.ResponseWriter,
	r *http.Request) {
	w.Write([]byte(pageContent))
}

// key type
const (
	keyUp uint64 = iota
	keyDown
	keyLeft
	keyRight
	keyVolUp
	keyVolDown
	keySpace
)

func keyServerRequest(
	w http.ResponseWriter,
	r *http.Request) {
	keyStr := r.FormValue("key")

	key, err := strconv.ParseUint(keyStr, 0, 32)
	if err != nil || key > keySpace {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	go press(key)
}
