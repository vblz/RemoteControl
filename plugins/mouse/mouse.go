package mouse

import (
	"fmt"
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
	path := dir + "\\static\\plugins\\mouse\\index.html"
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Can't load mouse static file: ", err)
		return
	}
	pageContent = f
}

const (
	clickTypeLeft int = iota
	clickTypeRight
)

// Control allow to remote use mouse
type Control struct {
}

type position struct {
	x, y int32
}

// GetHandlers implemintation of interfaces.EndPointInfo
func (c Control) GetHandlers() []interfaces.EndPointInfo {
	handlers := []interfaces.EndPointInfo{
		interfaces.BaseEndPointInfo{
			"/mouse/move",
			moveServerRequest,
			interfaces.EndPointAPI},
		interfaces.BaseEndPointInfo{
			"/mouse/click",
			clickServerRequest,
			interfaces.EndPointAPI},
		interfaces.BaseEndPointInfo{
			"/mouse",
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

func moveServerRequest(
	w http.ResponseWriter,
	r *http.Request) {
	xStr, yStr := r.FormValue("x"), r.FormValue("y")

	x, errX := strconv.ParseInt(xStr, 0, 32)
	y, errY := strconv.ParseInt(yStr, 0, 32)
	if errX != nil || errY != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	newX, newY := move(int32(x), int32(y))

	fmt.Fprint(w, newX, newY)
}

func clickServerRequest(
	w http.ResponseWriter,
	r *http.Request) {
	typeStr := r.FormValue("type")
	clickType, err := strconv.Atoi(typeStr)
	if err != nil || clickType > clickTypeRight {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	click(clickType)
}
