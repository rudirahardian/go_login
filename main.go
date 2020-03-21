package main

import (
	"github.com/rudirahardian/go_env/app/routes"
	"github.com/rudirahardian/go_env/config"
)

func main() {
	PORT := config.DotEnvVariable("PORT")
	MODE := config.DotEnvVariable("APP_ENV")
	routes.RouteInit(PORT,MODE)
}