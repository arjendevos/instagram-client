package helpers

import (
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

func GetEmailProviders() []string {
	return []string{
		"google",
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

func Handle(err error) {
	if err != nil {
		log.Err(err).Send()
		panic(err)
	}
}

func ExtractEmails(text string) []string {
	r, _ := regexp.Compile("[a-zA-Z0-9-_.]+@[a-zA-Z0-9-_.]+")

	return r.FindAllString(text, -1)
}

func FindPrimaryEmail(emails []string) (email string, provider string) {
	providers := GetEmailProviders()

	for _, e := range emails {
		for _, p := range providers {
			if strings.Contains(e, p) {
				return e, p
			}
		}
	}

	return "", ""
}
