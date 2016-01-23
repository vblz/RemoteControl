package shutdown

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/vblazhnov/RemoteControl/interfaces"
)

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
	f, e := ioutil.ReadFile(".\\static\\plugins\\shutdown\\index.html")
	if e != nil {
		fmt.Fprint(w, e)
	}
	w.Write(f)
}

func apiServeRequest(
	w http.ResponseWriter,
	r *http.Request) {
	secStr := r.FormValue("sec")

	sec, err := strconv.Atoi(secStr)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	abort()
	if sec != 0 {
		start(sec)
		fmt.Fprint(w, sec)
	}
}

func abort() {
	run("/a")
}

func start(sec int) {
	run("/s", "/t", strconv.Itoa(sec))
}

func run(args ...string) {
	cmd := exec.Command("shutdown", args...)
	err := cmd.Run()
	if err != nil {
		log.Println("error: " + err.Error())
	}
}
