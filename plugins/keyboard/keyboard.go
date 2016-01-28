package keyboard

import (
	// #include <wtypes.h>
	// #include <Winuser.h>
	"C"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/vblazhnov/RemoteControl/interfaces"
)

var (
	pageContent []byte
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println("Can't load mouse static file: ", err)
		return
	}
	path := dir + "\\static\\plugins\\keyboard\\index.html"
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Can't load keyboard static file: ", err)
		return
	}
	pageContent = f
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
	if pageContent != nil {
		w.Write(pageContent)
	} else {
		http.NotFound(w, r)
	}
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
