package Routes

import (
	"github.com/victorneuret/WatcherUpload/server/Config"
	"log"
	"net/http"
	"os"
)

func remove(_ http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	filePath := r.Form.Get("file")
	if filePath == "" {
		log.Println("Missing 'file' parameter")
		return
	}

	err = os.RemoveAll(Config.GetConfig().UploadDir + filePath)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(filePath, "removed successfully")
}
