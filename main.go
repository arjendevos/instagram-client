package main

import (
	"context"
	"errors"
	"flag"

	c "github.com/arjendevos/instagram-client/client"
	"github.com/arjendevos/instagram-client/database"
	"github.com/arjendevos/instagram-client/helpers"
	h "github.com/arjendevos/instagram-client/helpers"
	"github.com/arjendevos/instagram-client/instagram"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	h.ConfigureLogger()
}

func main() {
	usernamePtr := flag.String("username", "", "An username of the account you want to start from")
	flag.Parse()
	if usernamePtr == nil || *usernamePtr == "" {
		panic(errors.New("username is empty"))
	}

	client := c.NewClient()
	ctx := context.Background()
	db, err := database.StartDB()
	helpers.HandleError(err, main)

	instagram := instagram.NewInstagram(ctx, db)

	startProfile, err := client.GetProfile(*usernamePtr)
	helpers.HandleError(err, main)

	recommendedAccounts, err := client.GetRecommendedAccounts(startProfile.Data.User.ID)
	helpers.HandleError(err, main)

	for _, account := range recommendedAccounts.Users {
		p, err := client.GetProfile(account.Username)
		helpers.HandleError(err, main)
		instagram.Insert(p)
		helpers.Wait(15)
	}

	// api.Start(db)
}
