package main

import (
	"github.com/teguhbudhi13/user-service/app"
	"github.com/teguhbudhi13/user-service/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":" + config.Server.Port)
}
