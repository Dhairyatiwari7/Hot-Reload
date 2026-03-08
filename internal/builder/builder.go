package builder

import (
	"log"
	"os/exec"
)

func Build(cmdStr string) error {

	log.Println("Building project...")

	cmd := exec.Command("cmd", "/C", cmdStr)

	cmd.Stdout = nil
	cmd.Stderr = nil

	return cmd.Run()
}