package Watcher

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"

	"github.com/victorneuret/WatcherUpload/watcher/Config"
	"github.com/victorneuret/WatcherUpload/watcher/Utils"
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
