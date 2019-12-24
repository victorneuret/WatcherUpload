package main

import (
	"github.com/victorneuret/WatcherUpload/watcher/Config"
	"github.com/victorneuret/WatcherUpload/watcher/Watcher"
)

func main() {
	Config.LoadConfiguration()
	Watcher.Watcher()
}
