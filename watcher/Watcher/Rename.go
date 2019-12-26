package Watcher

import (
	"github.com/radovskyb/watcher"
	"github.com/victorneuret/WatcherUpload/watcher/Config"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

func rename(event watcher.Event) {
	config := Config.GetConfig()

	oldRelPath, err := filepath.Rel(config.WatchDir, event.OldPath)
	if err != nil {
		log.Println(err)
		return
	}
	newRelPath, err := filepath.Rel(config.WatchDir, event.Path)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = http.PostForm(config.ServerURL+config.RenamePath, url.Values{
		"old": {oldRelPath},
		"new": {newRelPath},
	})
	if err != nil {
		log.Println(err)
		return
	}
}
