package main

import (
	"ninhtq/go-gin/core/config"
	"ninhtq/go-gin/db"
	"ninhtq/go-gin/restful"
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
	config.Init(".env")
	db.Init()
	restful.Init()
}
