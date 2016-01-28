// +build windows

package shutdown

import (
	"log"
	"os/exec"
	"strconv"
)

func abort() {
	run("/a")
}

func start(sec int) {
	run("/s", "/t", strconv.Itoa(sec))
}

// possible crossplatform
func run(args ...string) {
	cmd := exec.Command("shutdown", args...)
	err := cmd.Run()
	if err != nil {
		log.Println("error: " + err.Error())
	}
}
