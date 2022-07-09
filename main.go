package main

import (
	"context"
	"errors"
	"flag"

	c "github.com/arjendevos/instagram-client/client"
	"github.com/arjendevos/instagram-client/database"
	"github.com/arjendevos/instagram-client/helpers"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	helpers.ConfigureLogger()
}

func main() {
	usernamePtr := flag.String("username", "", "An username of the account you want to start from")
	flag.Parse()
	if usernamePtr == nil || *usernamePtr == "" {
		panic(errors.New("username is empty"))
	}

	// Initalize project
	client := c.NewClient()
	ctx := context.Background()
	dc, err := database.NewDatabaseClient(ctx)
	helpers.HandleError(err, main)

	defer dc.DB.Close()

	startProfile, err := client.GetProfile(*usernamePtr)
	helpers.HandleError(err, main)

	err = client.Start(dc, startProfile.Data.User.ID)
	helpers.HandleError(err, main)

}
