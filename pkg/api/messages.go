package api

import (
	"encoding/json"
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func MessagesAPI(c *fiber.Ctx) error {
	// Parse the request body
	var msg Message
	if err := c.BodyParser(&msg); err != nil {
		return err
	}

	// Check if any of the fields are empty
	if msg.Name == "" || msg.Email == "" || msg.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Please provide all fields",
		})
	}

	// Read existing messages from file
	file, err := os.OpenFile("messages.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the messages from file into a slice of Message structs
	var messages []Message
	if err := json.NewDecoder(file).Decode(&messages); err != nil && err != io.EOF {
		return err
	}

	// Append new message to slice of messages
	messages = append(messages, msg)

	// Write updated messages to file
	file.Seek(0, 0)
	file.Truncate(0)
	if err := json.NewEncoder(file).Encode(messages); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
