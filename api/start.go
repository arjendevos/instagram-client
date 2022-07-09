package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	v1 "github.com/arjendevos/instagram-client/api/v1"
	"github.com/arjendevos/instagram-client/database"
	"github.com/arjendevos/instagram-client/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func init() {
	godotenv.Load("../.env")
	helpers.ConfigureLogger()
}

func main() {
	ctx := context.Background()
	dc, err := database.NewDatabaseClient(ctx)
	helpers.HandleError(err, main)

	Route(dc.Ctx, dc.DB)
}

func Route(ctx context.Context, db *sql.DB) {
	port := helpers.GetEnv("PORT")
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        3,
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowMethods: "GET,POST,PATCH",
	// }))

	endpoints := v1.InitEndpoints(ctx, db)

	app.Get("/api/v1/profiles", endpoints.GetProfiles)

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	return postHandler(c, db)
	// })

	// app.Put("/update", func(c *fiber.Ctx) error {
	// 	return putHandler(c, db)
	// })

	// app.Delete("/delete", func(c *fiber.Ctx) error {
	// 	return deleteHandler(c, db)
	// })
	// app.Delete("/delete", deleteHandler) // Add this
	log.Fatal().Err(app.Listen(fmt.Sprintf(":%v", port)))
	log.Info().Str("port", port).Msg("Api route succesfully started")

}
