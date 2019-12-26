package Watcher

import (
	"log"

	"github.com/radovskyb/watcher"

	"github.com/victorneuret/WatcherUpload/watcher/Config"
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
		log.Println("File Created:", event.Path)
		upload(event)
	case watcher.Write:
		log.Println("File Written:", event.Path)
		upload(event)
	case watcher.Remove:
		log.Println("File Removed:", event.Path)
	case watcher.Rename:
		log.Println("File Renamed:", event.Path)
	case watcher.Move:
		log.Println("File Moved from:", event.OldPath, "to:", event.Path)
	default:
		log.Println("Action not handled:", event.Op, "on", event.Path)
	}
}
