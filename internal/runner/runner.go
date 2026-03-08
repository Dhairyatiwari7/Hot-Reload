package runner

import (
	"log"
	"os"
	"os/exec"
)

var serverCmd *exec.Cmd

func Restart(runCmd string) {

	if serverCmd != nil && serverCmd.Process != nil {
		log.Println("Stopping previous server...")
		serverCmd.Process.Kill()
	}

	log.Println("Starting server...")

	serverCmd = exec.Command("cmd", "/C", runCmd)

	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = os.Stderr

	err := serverCmd.Start()

	if err != nil {
		log.Println("Server start failed:", err)
	}
}