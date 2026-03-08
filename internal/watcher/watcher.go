package filewatcher

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func Start(root string, changes chan bool) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for {

		select {

		case event := <-watcher.Events:

			if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {

				log.Println("File changed:", event.Name)

				changes <- true
			}

		case err := <-watcher.Errors:

			log.Println("Watcher error:", err)
		}
	}
}