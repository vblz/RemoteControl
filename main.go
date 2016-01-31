package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/vblazhnov/go-http-digest-auth"

	"github.com/vblazhnov/RemoteControl/interfaces"
	"github.com/vblazhnov/RemoteControl/plugins/keyboard"
	"github.com/vblazhnov/RemoteControl/plugins/mouse"
	"github.com/vblazhnov/RemoteControl/plugins/shutdown"
	"github.com/vblazhnov/RemoteControl/utils"
)

var (
	serverAddress   string
	serverPort      uint
	user            string
	password        string
	certTLSFilePath string
	keyTLSFilePath  string
	wrapper         auth.Wrapper
)

func init() {
	flag.StringVar(&serverAddress, "host", "0.0.0.0", "the host that server binds to")
	flag.UintVar(&serverPort, "port", 1234, "the port that server binds to")
	flag.StringVar(&user, "user", "user", "username for auth while using the service")
	flag.StringVar(&password, "password", "iospassword", "password for auth while using the service")
	flag.StringVar(&certTLSFilePath, "TLSCertPath", "", "path to cert file for use TLS")
	flag.StringVar(&keyTLSFilePath, "TLSKeyPath", "", "path to key file for use TLS")
}

func main() {
	flag.Parse()
	wrapper = auth.NewBaseAuth(auth.Info{user, password, serverAddress})

	initPlugins()
	initResources()
	startServer()
}

func startServer() {
	var err error
	addr := serverAddress + ":" + strconv.FormatUint(uint64(serverPort), 10)
	if certTLSFilePath == "" || keyTLSFilePath == "" {
		err = http.ListenAndServe(addr, nil)
	} else {
		err = http.ListenAndServeTLS(addr, certTLSFilePath, keyTLSFilePath, nil)
	}
	if err != nil {
		log.Fatal(err)
	}
}

var resources map[string][]byte

func initResources() {
	resources = make(map[string][]byte, 3)
	appendResources(utils.ReadFilesInDir("static\\css"))
	appendResources(utils.ReadFilesInDir("static\\js"))
	handleFilesMap(resources)
}

func appendResources(new map[string][]byte) {
	for k, v := range new {
		k = "/" + strings.Replace(k, "\\", "/", -1)
		resources[k] = v
	}
}

func handleFilesMap(files map[string][]byte) {
	for k := range files {
		handleRequest(k, func(w http.ResponseWriter, r *http.Request) {
			path := r.RequestURI
			res, ok := resources[path]
			if !ok {
				http.NotFound(w, r)
			} else {
				if strings.Contains(path, ".css") {
					w.Header().Add("Content-Type", "text/css")
				}
				w.Write(res)
			}
		})
	}
}

func initPlugins() {
	plugins := []interfaces.Plugin{
		shutdown.Control{},
		mouse.Control{},
		keyboard.Control{},
	}
	registerPlugins(plugins)
}

func registerPlugins(plugins []interfaces.Plugin) {
	var err error
	menus = make(map[template.URL]string, len(plugins))
	mainTemplate, err = template.New("mainTemplate").Parse(string(utils.ReadHTML("\\static\\template.html")))
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range plugins {
		for _, ep := range p.GetHandlers() {
			switch ep.Type() {
			case interfaces.EndPointAPI:
				handleRequest("/api/v1"+ep.Path(), ep.Handler())
			case interfaces.EndPointContent:
				handleRequest(ep.Path(), ep.Handler())
			default:
				log.Println("Incorrect plugin endpoint type: ", ep)
			}
		}
		for _, static := range p.GetMainContent() {
			menus[static.Path()] = static.Title()
			handleStatic(static.Path(), static.Data, static.Title)
		}
	}
}

var (
	mainTemplate *template.Template
	menus        map[template.URL]string
)

func handleStatic(path template.URL, data func() template.HTML, title func() string) {
	templateInner := struct {
		Title     string
		MenuItems map[template.URL]string
		Content   template.HTML
	}{
		title(),
		menus,
		data(),
	}
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		mainTemplate.Execute(w, templateInner)
	}
	http.HandleFunc(string(path), wrapper.Wrap(handleFunc))
}

func handleRequest(path string, fun http.HandlerFunc) {
	http.HandleFunc(path, wrapper.Wrap(fun))
}
