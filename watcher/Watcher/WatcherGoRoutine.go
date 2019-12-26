package Watcher

import (
	"log"
	"os"
	"path/filepath"

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
	filename := filepath.Base(event.Path)
	if filename == ".DS_Store" {
		return
	}

	fi, err := os.Stat(event.Path)
	if err == nil && fi.Mode().IsDir() {
		return
	}

	switch event.Op {
	case watcher.Create:
		log.Println("File Created:", event.Path)
		upload(event)
	case watcher.Write:
		log.Println("File Written:", event.Path)
		upload(event)
	case watcher.Remove:
		log.Println("File Removed:", event.Path)
		remove(event)
	case watcher.Rename:
		log.Println("File Renamed from:", event.OldPath, "to:", event.Path)
		rename(event)
	case watcher.Move:
		log.Println("File Moved from:", event.OldPath, "to:", event.Path)
		rename(event)
	default:
		log.Println("Action not handled:", event.Op, "on", event.Path)
	}
}
