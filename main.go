package main

import (
	"os"

	"github.com/bluebeel/orm/app"
)

func main() {
	port := ":" + os.Getenv("PORT")
	server := &app.App{}
	server.Initialize()
	server.Run(port)
}
