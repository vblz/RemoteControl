package shutdown

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
		log.Println("Can't load shutdown static file: ", err)
		return
	}
	path := dir + "\\static\\plugins\\shutdown\\index.html"
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Can't load shutdown static file: ", err)
		return
	}
	pageContent = f
}

// Shutdown shutdown the local machine
type Shutdown struct {
}

// GetHandlers implemintation of interfaces.EndPointInfo
func (sh Shutdown) GetHandlers() []interfaces.EndPointInfo {
	handlers := []interfaces.EndPointInfo{
		interfaces.BaseEndPointInfo{
			"/shutdown",
			apiServeRequest,
			interfaces.EndPointAPI},
		interfaces.BaseEndPointInfo{
			"/shutdown",
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

func apiServeRequest(
	w http.ResponseWriter,
	r *http.Request) {
	secStr := r.FormValue("sec")

	sec, err := strconv.Atoi(secStr)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	process(sec)
	fmt.Fprint(w, sec)
}

func process(sec int) {
	abort()
	if sec != 0 {
		start(sec)
	}
}
