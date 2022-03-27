package routes

import (
	"fmt"
	"log"
	"net/http"
	"sigolang-sv/configs"

	"github.com/rs/cors"
)

func LaunchServer() {
	WebServer := fmt.Sprintf("Web Server is Running on Port %v", configs.WEB_PORT)
	log.Println(WebServer)
	handler := cors.AllowAll().Handler(Route())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", configs.WEB_PORT), handler))
}