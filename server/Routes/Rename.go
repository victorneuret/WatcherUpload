package Routes

import (
	"github.com/victorneuret/WatcherUpload/server/Config"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func rename(_ http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	oldPath := r.Form.Get("old")
	newPath := r.Form.Get("new")
	if oldPath == "" || newPath == "" {
		log.Println("Missing 'old' or 'new' parameter")
		return
	}

	config := Config.GetConfig()

	err = os.Rename(config.UploadDir+oldPath, config.UploadDir+newPath)
	if err != nil {
		log.Println(err)
		return
	}

	if filepath.Dir(oldPath) != filepath.Dir(newPath) {
		log.Println(oldPath, "moved to", newPath, "successfully")
	} else {
		log.Println(oldPath, "renamed to", newPath, "successfully")
	}
}
