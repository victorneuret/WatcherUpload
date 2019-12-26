package main

import (
	"github.com/victorneuret/WatcherUpload/server/Config"
	"github.com/victorneuret/WatcherUpload/server/Routes"
	"github.com/victorneuret/WatcherUpload/server/Utils"
)

func main() {
	Config.LoadConfiguration()
	Utils.CreateDirIfNotExist(Config.GetConfig().UploadDir)
	Routes.SetupRoutes()
}
