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
	"syscall"

	"github.com/vblazhnov/RemoteControl/interfaces"
)

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	keybdEventInfo = user32.NewProc("SendInput")
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

func contentServeRequest(
	w http.ResponseWriter,
	r *http.Request) {
	if pageContent != nil {
		w.Write(pageContent)
	} else {
		http.NotFound(w, r)
	}
}

func keyServerRequest(
	w http.ResponseWriter,
	r *http.Request) {
	keyStr := r.FormValue("key")

	key, err := strconv.ParseUint(keyStr, 0, 32)
	if err != nil || key > keyVolDown {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	go press(key)
}

const (
	keyUp uint64 = iota
	keyDown
	keyLeft
	keyRight
	keyVolUp
	keyVolDown
)

func press(key uint64) {
	switch key {
	case keyUp:
		keybdEvent(123)
	}
}

func keybdEvent(keyCode int32) {
	// down := C.INPUT{_type: 1}
	// kis := []C.INPUT{
	// 	down,
	// 	//{_type: C.INPUT_KEYBOARD, ki: C.KEYBDINPUT{wVk: uint16(C.VK_VOLUME_DOWN)}},
	// 	//{_type: C.INPUT_KEYBOARD, ki: C.KEYBDINPUT{wVk: uint16(C.VK_VOLUME_DOWN), dwFlags: C.KEYEVENTF_KEYUP}},
	// }
	//
	// keybdEventInfo.Call(
	// 	uintptr(2),
	// 	uintptr(unsafe.Pointer(&kis[0])),
	// 	uintptr(unsafe.Sizeof(C.INPUT)),
	// )
}
