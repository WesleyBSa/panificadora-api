package main

import (
	"panificadora-api/config"
	"panificadora-api/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
