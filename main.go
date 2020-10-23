package main

import "github.com/faruoqi/elemental/core"

const (
	ENV = "dev"
)

func main() {

	var app core.AppCore
	app = core.NewApp(ENV)
	app.Start()

}
