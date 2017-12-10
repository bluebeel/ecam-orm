package main

import (
	"github.com/bluebeel/orm/app"
)

func main() {
	server := &app.App{}
	server.Initialize()
	server.Run(":3000")
}
