package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/vblazhnov/RemoteControl/interfaces"
	"github.com/vblazhnov/RemoteControl/plugins/shutdown"
)

var (
	serverAddress string
	serverPort    uint
	user          string
	password      string
)

func init() {
	flag.StringVar(&serverAddress, "host", "127.0.0.1", "the host that server binds to")
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
	plugins := []interfaces.Plugin{}
	plugins = append(plugins, shutdown.Shutdown{})
	registerPlugins(plugins)
}

func registerPlugins(plugins []interfaces.Plugin) {
	for _, p := range plugins {
		for _, ep := range p.GetHandlers() {
			http.HandleFunc(ep.Path(), ep.Handler())
		}
	}
}
