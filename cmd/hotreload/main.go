package main

import (
	"flag"
	"log"

	"github.com/dhairya/hotreload/internal/builder"
	"github.com/dhairya/hotreload/internal/runner"
	filewatcher "github.com/dhairya/hotreload/internal/watcher"
)

func main() {

	root := flag.String("root", ".", "project root")
	buildCmd := flag.String("build", "", "build command")
	runCmd := flag.String("exec", "", "run command")

	flag.Parse()

	log.Println("HotReload starting...")

	changes := make(chan bool)

	go filewatcher.Start(*root, changes)

	err := builder.Build(*buildCmd)
	if err != nil {
		log.Println("Initial build failed:", err)
		return
	}

	runner.Restart(*runCmd)

	for range changes {

		log.Println("Change detected")

		err := builder.Build(*buildCmd)
		if err != nil {
			log.Println("Build failed")
			continue
		}

		runner.Restart(*runCmd)
	}
}