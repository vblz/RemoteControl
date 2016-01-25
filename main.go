package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/vblazhnov/RemoteControl/auth"
	"github.com/vblazhnov/RemoteControl/interfaces"
	"github.com/vblazhnov/RemoteControl/plugins/mouse"
	"github.com/vblazhnov/RemoteControl/plugins/shutdown"
)

var (
	serverAddress   string
	serverPort      uint
	user            string
	password        string
	wrapper         auth.Wrapper
)

func init() {
	flag.StringVar(&serverAddress, "host", "0.0.0.0", "the host that server binds to")
	flag.UintVar(&serverPort, "port", 1234, "the port that server binds to")
	flag.StringVar(&user, "user", "user", "username for auth while using the service")
	flag.StringVar(&password, "password", "пароль", "password for auth while using the service")
}

func main() {
	flag.Parse()
	initPlugins()
	startServer()
}

func startServer() {
	err := http.ListenAndServe(serverAddress+":"+strconv.FormatUint(uint64(serverPort), 10), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initPlugins() {
	plugins := []interfaces.Plugin{
		shutdown.Shutdown{},
		mouse.Control{},
	}
	registerPlugins(plugins)
}

func registerPlugins(plugins []interfaces.Plugin) {
	wrapper = auth.Wrapper{user, password, serverAddress}
	for _, p := range plugins {
		for _, ep := range p.GetHandlers() {
			switch ep.Type() {
			case interfaces.EndPointAPI:
				handleFunc("/api/v1"+ep.Path(), ep.Handler())
			case interfaces.EndPointContent:
				handleFunc(ep.Path(), ep.Handler())
			default:
				log.Println("Incorrect plugin endpoint type: ", ep)
			}
		}
	}
}

func handleFunc(path string, fun http.HandlerFunc) {
	http.HandleFunc(path, wrapper.Wrap(fun))
}
