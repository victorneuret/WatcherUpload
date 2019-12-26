package Watcher

import (
	"github.com/radovskyb/watcher"
	"github.com/victorneuret/WatcherUpload/watcher/Config"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

func remove(event watcher.Event) {
	config := Config.GetConfig()

	relPath, err := filepath.Rel(config.WatchDir, event.Path)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = http.PostForm(config.ServerURL+config.RemovePath, url.Values{"file": {relPath}})
	if err != nil {
		log.Println(err)
		return
	}
}
