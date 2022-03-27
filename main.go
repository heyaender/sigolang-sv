package main

import (
	"sigolang-sv/database"
	"sigolang-sv/routes"
)

func main() {
	database.MySQLConnect()
	routes.LaunchServer()
}