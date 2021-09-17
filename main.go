package main

import (
	"todo/config"
	"todo/routes"
)

func main() {
	config.SetupDB()
	r := routes.SetupRouter()
	r.Run()
}
