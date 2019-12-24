package Config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	WatchDir string `json:"watch-dir"`
}

var Config Configuration

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
	err = jsonParser.Decode(&Config)
	if err != nil {
		panic("Config file loading failed")
	}
}
