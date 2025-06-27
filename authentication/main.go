package main

import (
	"log"
	"os"

	_ "embed"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type user struct {
	StreamKey string `form:"name" yaml:"username"`
	Password  string `form:"key" yaml:"password"`
}

type configFile struct {
	Listen string `yaml:"listen"`
	Users  []user `yaml:"users"`
}

func main() {
	file, err := os.ReadFile("vrstream.yml")
	if err != nil {
		log.Fatalf("Error reading vrstream.yml: %v", err)
	}

	var config configFile
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error parsing vrstream.yml: %v", err)
	}

	users := config.Users
	if len(users) == 0 {
		log.Fatal("No users found in vrstream.yml")
	}

	app := fiber.New()

	app.Post("/auth", func(c *fiber.Ctx) error {
		var input user
		if err := c.BodyParser(&input); err != nil {
			log.Printf("(/auth): BodyParser Error: %v", err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		for _, u := range users {
			if u.Password == input.Password && u.StreamKey == input.StreamKey {
				log.Printf("(/auth): Authenticated user: %s", u.StreamKey)
				c.Status(fiber.StatusNoContent)
				return nil
			}
		}

		log.Printf("(/auth): Authentication failed: %s", input.StreamKey)
		c.Status(fiber.StatusUnauthorized)
		return nil
	})

	app.Post("/publish", func(c *fiber.Ctx) error {
		var input user
		if err := c.BodyParser(&input); err != nil {
			log.Printf("(/publish): BodyParser Error: %v", err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		log.Printf("(/publish): Stream finished: %s", input.StreamKey)
		c.Status(fiber.StatusNoContent)
		return nil
	})

	if err := app.Listen(config.Listen); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
