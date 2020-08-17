package main

import (
	"github.com/juanbelieni/go-simple-api/app"
	"github.com/juanbelieni/go-simple-api/database"
)

func main() {
	app := app.GetApp()
	app.Listen(3333)
	defer database.Connection.Close()
}
