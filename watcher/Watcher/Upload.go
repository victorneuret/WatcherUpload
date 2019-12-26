package Watcher

import (
	"bytes"
	"github.com/radovskyb/watcher"
	"github.com/victorneuret/WatcherUpload/watcher/Config"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func upload(event watcher.Event) {
	file, err := os.Open(event.Path)
	if err != nil {
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

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	relPath, err := filepath.Rel(config.WatchDir, event.Path)
	if err != nil {
		log.Println(err)
		return
	}

	part, err := writer.CreateFormFile("file", relPath)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Println(err)
		return
	}

	err = writer.Close()
	if err != nil {
		log.Println(err)
		return
	}

	r, _ := http.NewRequest("POST", config.ServerURL+config.UploadPath, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return
	}

	message, _ := ioutil.ReadAll(res.Body)
	log.Printf(string(message))
}
