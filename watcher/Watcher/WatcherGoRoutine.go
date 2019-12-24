package Watcher

import (
	"fmt"
	"log"

	"github.com/radovskyb/watcher"
)

func watcherGoRoutine(w *watcher.Watcher) {
	for {
		select {
		case event := <-w.Event:
			fmt.Println(event) // Print the event's info.
		case err := <-w.Error:
			log.Fatalln(err)
		case <-w.Closed:
			return
		}
	}
}
