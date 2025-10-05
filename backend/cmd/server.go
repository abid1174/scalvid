package cmd

import (
	"scalvid/config"
	"scalvid/rest"
)

func Serve() {
	config := config.GetConfig()

	rest.StartServer(config)
}
