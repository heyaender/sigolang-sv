package main

import (
	"sigolang-sv/databases"
	"sigolang-sv/routes"
)

func main() {
	databases.MySQLConnect()
	routes.LaunchServer()
}