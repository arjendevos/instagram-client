package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

func GetEmailProviders() []string {
	return []string{
		"gmail",
		"yahoo",
		"icloud",
		"hotmail",
		"outlook",
		"mail",
	}
}

func GetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Error().Str("Could not get key: ", key).Send()

	}

	return env
}

func GetEnvArray(key string) []string {
	env := GetEnv(key)
	return strings.Split(env, ",")
}

func HandleError(err error, cb func()) {
	if err != nil {
		log.Err(err).Send()
		Wait(60)
		cb()
	}
}

func ExtractEmails(text string) []string {
	r, _ := regexp.Compile("[a-zA-Z0-9-_.]+@[a-zA-Z0-9-_.]+")

	return r.FindAllString(text, -1)
}

func FindPrimaryEmail(emails []string) (email string, provider string) {
	if len(emails) > 0 {
		providers := GetEmailProviders()

		for _, e := range emails {
			for _, p := range providers {
				if strings.Contains(e, p) {
					return e, p
				}
			}
		}

		return emails[0], ""
	}

	return "", ""
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

func Wait(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}
