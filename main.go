package main

import (
	"github.com/rudirahardian/go_env/app/routes"
	"github.com/rudirahardian/go_env/config"
)

func main() {
	dotenv := config.DotEnvVariable("PORT")
	routes.RouteInit(dotenv)
}