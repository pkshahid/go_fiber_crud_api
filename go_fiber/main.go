package main

import (
        "github.com/gofiber/fiber/v2"
        "fmt"
        "go_fiber/database"
        "go_fiber/router"
)

func main() {
        // Start a new fiber app
        app := fiber.New()

        //Connect DB
        database.ConnectDB()
        // Send a string back for GET calls to the endpoint "/"
        app.Get("/", func(c *fiber.Ctx) error {
                err := c.SendString("And the API is UP!")
                return err
        })

        router.SetupRoutes(app)

        fmt.Println("Fiber API Server running on port 8000")

        // Listen on PORT 300
        app.Listen(":8000")
}



