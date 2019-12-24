package Watcher

import (
	"github.com/victorneuret/WatcherUpload/Config"
	"log"
	"time"

	"github.com/radovskyb/watcher"

	"github.com/victorneuret/WatcherUpload/Utils"
)

func Watcher() {
	w := watcher.New()

	go watcherGoRoutine(w)

	Utils.CreateDirIfNotExist(Config.GetConfig().WatchDir)
	if err := w.AddRecursive(Config.GetConfig().WatchDir); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
