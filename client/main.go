package client

import (
	"fmt"

	"github.com/arjendevos/instagram-client/database"
	"github.com/arjendevos/instagram-client/helpers"
	"github.com/rs/zerolog/log"
)

func (c *Client) Start(dc *database.DatabaseClient, userID string) error {
	log.Info().Str("userID", userID).Msg("getting recommended accounts")
	var nextUserID string
	recommendedAccounts, err := c.GetRecommendedAccounts(userID)
	if err != nil {
		return err
	}

	for _, account := range recommendedAccounts.Users {
		exists, err := dc.DoesProfileExists(fmt.Sprint(account.Pk))
		if exists {
			if err != nil {
				log.Err(err)
			} else {
				log.Warn().Str("username", account.Username).Msg("user already exists")
			}
			continue
		}

		helpers.Wait(15)

		p, err := c.GetProfile(account.Username)
		if err != nil {
			log.Err(err)
			continue
		}

		err = dc.InsertProfile(p)
		if err != nil {
			log.Err(err)
			continue
		}

		if p.Data.User.EdgeFollowedBy.Count > 20000 && nextUserID == "" {
			nextUserID = p.Data.User.ID
		}
	}

	helpers.Wait(60)
	c.Start(dc, nextUserID)

	return nil
}
