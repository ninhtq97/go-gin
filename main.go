package main

import (
	"fmt"
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/restful"
	"os"
)

// @Title 			Bvote API
// @Version         1.0
// @Description 	This is a server Bvote
// @Schemes			http https

// @BasePath		/api

// @SecurityDefinitions.apikey Bearer
// @In header
// @Name authorization
// @Description "Type 'Bearer TOKEN' to correctly set the API Key"
func main() {
	conf, err := config.LoadConfig(".env")

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load env conf: %s\n", err)
		os.Exit(1)
	}

	server := &restful.Server{Config: conf}
	server.SetupRouter()

	server.Router.Run(fmt.Sprintf(":%v", conf.ServerPort))
}
