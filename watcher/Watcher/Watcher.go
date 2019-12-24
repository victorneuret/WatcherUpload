package Watcher

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"

	"github.com/victorneuret/WatcherUpload/Utils"
)

func Watcher() {
	w := watcher.New()

	go watcherGoRoutine(w)

	Utils.CreateDirIfNotExist("/tmp/watching")
	if err := w.AddRecursive("/tmp/watching"); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
