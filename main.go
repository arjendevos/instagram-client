package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"os"

	c "github.com/arjendevos/instagram-client/client"
	h "github.com/arjendevos/instagram-client/helpers"
	"github.com/joho/godotenv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	godotenv.Load()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	usernamePtr := flag.String("username", "", "An username of the account you want to start from")
	flag.Parse()
	if usernamePtr == nil || *usernamePtr == "" {
		h.Handle(errors.New("username is empty"))
	}

	client := c.NewClient()

	profile, err := client.GetProfile(*usernamePtr)
	h.Handle(err)

	// // resp, err := client.GetRecommendedAccounts(profile.Data.User.ID)
	// // h.Handle(err)

	// // err = WriteToFile(resp.Users, "users.json")
	// // h.Handle(err)

	err = WriteToFile(h.ConvertProfile(profile), "profile.json")
	h.Handle(err)

	// api.Start()

}

func WriteToFile(data any, file string) error {
	f, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, f, 0644)
	if err != nil {
		return err
	}

	return nil
}
