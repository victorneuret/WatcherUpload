package Watcher

import (
	"github.com/victorneuret/WatcherUpload/Config"
	"log"

	"github.com/radovskyb/watcher"
)

func watcherGoRoutine(w *watcher.Watcher) {
	for {
		select {
		case event := <-w.Event:
			if event.Path != Config.GetConfig().WatchDir {
				modificationType(event)
			}
		case err := <-w.Error:
			log.Fatalln(err)
		case <-w.Closed:
			return
		}
	}
}

func modificationType(event watcher.Event) {
	switch event.Op {
	case watcher.Create:
	case watcher.Write:
	case watcher.Remove:
	case watcher.Rename:
	case watcher.Move:
	default:
	}
}
