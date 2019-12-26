package Routes

import (
	"github.com/victorneuret/WatcherUpload/server/Config"
	"github.com/victorneuret/WatcherUpload/server/Utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func upload(_ http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 << 20)
	if err != nil {
		log.Println(err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	config := Config.GetConfig()

	Utils.CreateDirIfNotExist(config.UploadDir + filepath.Dir(handler.Filename))

	f, err := os.OpenFile(config.UploadDir+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Upload of " + handler.Filename + " successful")
}
