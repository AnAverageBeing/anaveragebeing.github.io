package main

import (
	"flag"
	"fmt"
	"log"

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
	//start server
	fmt.Printf("Server listening on %s...\n", *addr)
	log.Fatal(app.Listen(*addr))
}
