package main

import (
	"github.com/victorneuret/watcher/Config"
	"github.com/victorneuret/watcher/Watcher"
)

func main() {
	Config.LoadConfiguration()
	Watcher.Watcher()
}
