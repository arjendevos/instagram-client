package database

import (
	"github.com/arjendevos/instagram-client/helpers"
	"github.com/arjendevos/instagram-client/models/dm"
	"github.com/arjendevos/instagram-client/models/responses"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (dc *DatabaseClient) InsertProfile(a *responses.Profile) error {
	p := helpers.ConvertProfile(a)

	return p.Insert(dc.Ctx, dc.DB, boil.Infer())
}

func (dc *DatabaseClient) DoesProfileExists(userID string) (bool, error) {
	count, err := dm.InstagramAccounts(dm.InstagramAccountWhere.InstagramID.EQ(userID)).Count(dc.Ctx, dc.DB)
	if err != nil {
		return true, err
	}

	return count > 0, nil
}
