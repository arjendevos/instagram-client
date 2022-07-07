package instagram

import (
	"errors"

	"github.com/arjendevos/instagram-client/helpers"
	"github.com/arjendevos/instagram-client/models/dm"
	"github.com/arjendevos/instagram-client/models/responses"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i *Instagram) Insert(a *responses.Profile) error {
	p := helpers.ConvertProfile(a)
	count, err := dm.InstagramAccounts(dm.InstagramAccountWhere.Username.EQ(p.Username)).Count(i.ctx, i.db)
	if err != nil {
		return err
	}

	if count > 0 {
		log.Info().Str("username", p.Username).Msg("user already exists")
		return errors.New("user already exists")
	}

	return p.Insert(i.ctx, i.db, boil.Infer())
}
