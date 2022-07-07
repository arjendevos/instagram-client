package api

import (
	"database/sql"
	"fmt"

	"github.com/arjendevos/instagram-client/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func Start(db *sql.DB) {
	port := helpers.GetEnv("PORT")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	// app.Delete("/delete", deleteHandler) // Add this

	log.Fatal().Interface("app failed", (app.Listen(fmt.Sprintf(":%v", port))))
}
