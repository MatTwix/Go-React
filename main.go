package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type Request struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	API_PATH := os.Getenv("API_PATH")

	requests := []Request{}

	app.Get(API_PATH, func(c fiber.Ctx) error {
		return c.Status(200).JSON(requests)
	})

	// Create request
	app.Post(API_PATH, func(c fiber.Ctx) error {
		request := &Request{}

		if err := c.Bind().Body(request); err != nil {
			return err
		}

		if request.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "request body is required"})
		}

		request.ID = len(requests) + 1
		requests = append(requests, *request)

		return c.Status(201).JSON(request)
	})

	// Update request
	app.Patch(API_PATH+"/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for idx, request := range requests {
			if strconv.Itoa(request.ID) == id {
				requests[idx].Completed = true
				return c.Status(200).JSON(requests[idx])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Request not found"})
	})

	// Delete request
	app.Delete(API_PATH+"/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for idx, request := range requests {
			if strconv.Itoa(request.ID) == id {
				requests = append(requests[:idx], requests[idx+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Request not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
