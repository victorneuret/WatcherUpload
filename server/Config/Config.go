package Config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Configuration struct {
	UploadDir string `json:"upload-dir"`
}

var config Configuration
var mutex sync.Mutex

func GetConfig() *Configuration {
	mutex.Lock()
	defer mutex.Unlock()
	return &config
}

func LoadConfiguration() {
	configFile, err := os.Open("Config/config.json")
	defer func() {
		if err := configFile.Close(); err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		panic("Can't open config file")
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		panic("Config file loading failed")
	}
}
