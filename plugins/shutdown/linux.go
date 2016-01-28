// +build linux
// +build darwin

package shutdown

import (
	"log"
	"os/exec"
	"strconv"
)

func abort() {
	run("-c")
}

func start(sec int) {
	run("-h", "-t", "sec", strconv.Itoa(sec))
}

// possible crossplatform
func run(args ...string) {
	cmd := exec.Command("shutdown", args...)
	err := cmd.Run()
	if err != nil {
		log.Println("error: " + err.Error())
	}
}
