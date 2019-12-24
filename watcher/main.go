package main

import (
	"github.com/victorneuret/WatcherUpload/Config"
	"github.com/victorneuret/WatcherUpload/Watcher"
)

func main() {
	Config.LoadConfiguration()
	Watcher.Watcher()
}
