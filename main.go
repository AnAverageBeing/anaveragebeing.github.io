package main

import (
	"flag"
	"fmt"
	"log"
	"personal-site/pkg/api"

	"github.com/gofiber/fiber/v2"
)

var (
	addr = flag.String("addr", ":8080", "webserver addr")
)

func main() {
	flag.Parse()
	app := fiber.New()

	// Serve static files from the "web" directory
	app.Static("/", "./web")

	// POST endpoint for saving message data to JSON file
	app.Post("/api/messages", api.MessagesAPI)

	//start server
	fmt.Printf("Server listening on %s...\n", *addr)
	log.Fatal(app.Listen(*addr))
}
