package Routes

import (
	"log"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
