package shutdown

import (
	"fmt"
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
	pageContent = utils.ReadHTML("\\static\\plugins\\shutdown\\index.html")
}

// Control shutdown the local machine
type Control struct {
}

// GetHandlers implemintation of interfaces.Plugin
func (c Control) GetHandlers() []interfaces.EndPointInfo {
	handlers := []interfaces.EndPointInfo{
		interfaces.BaseEndPointInfo{
			"/shutdown",
			apiServeRequest,
			interfaces.EndPointAPI},
	}

	return handlers
}

// GetMainContent implemintation of interfaces.Plugin
func (c Control) GetMainContent() []interfaces.StaticContent {
	statics := []interfaces.StaticContent{
		interfaces.BaseStaticContent{
			"/shutdown",
			pageContent,
			"Shutdown",
		},
	}

	return statics
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
