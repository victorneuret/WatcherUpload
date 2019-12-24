package Watcher

import (
	"WatcherModule/Utils"
	"github.com/radovskyb/watcher"
	"log"
	"time"
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
