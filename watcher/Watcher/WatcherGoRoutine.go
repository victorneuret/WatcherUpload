package Watcher

import (
	"fmt"
	"github.com/victorneuret/WatcherUpload/Config"
	"log"

	"github.com/radovskyb/watcher"
)

func watcherGoRoutine(w *watcher.Watcher) {
	for {
		select {
		case event := <-w.Event:
			if event.Path != Config.GetConfig().WatchDir {
				fmt.Println(event)
			}
		case err := <-w.Error:
			log.Fatalln(err)
		case <-w.Closed:
			return
		}
	}
}
