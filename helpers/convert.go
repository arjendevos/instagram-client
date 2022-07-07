package helpers

import (
	"github.com/arjendevos/instagram-client/models/dm"
	"github.com/arjendevos/instagram-client/models/responses"
)

func ConvertProfile(pr *responses.Profile) *dm.InstagramAccount {
	emails := ExtractEmails(pr.Data.User.Biography)
	pe, pv := FindPrimaryEmail(emails)
	hasManagement := false
	if len(emails) > 0 && pv == "" {
		hasManagement = true
	}

	return &dm.InstagramAccount{
		InstagramID:           pr.Data.User.ID,
		IsBusinessAccount:     pr.Data.User.IsBusinessAccount,
		IsProfessionalAccount: pr.Data.User.IsProfessionalAccount,
		IsVerified:            pr.Data.User.IsVerified,
		Username:              pr.Data.User.Username,
		FullName:              ns(pr.Data.User.FullName),
		ProfilePicture:        pr.Data.User.ProfilePicURLHd,
		PostCount:             pr.Data.User.EdgeOwnerToTimelineMedia.Count,
		FollowersCount:        int64(pr.Data.User.EdgeFollowedBy.Count),
		FollowingCount:        pr.Data.User.EdgeFollow.Count,
		ExternalURL:           ns(pr.Data.User.ExternalURL),
		BusinessEmail:         ns(pr.Data.User.BusinessEmail),
		BusinessPhone:         ns(pr.Data.User.BusinessPhoneNumber),
		BusinessCategory:      ns(pr.Data.User.BusinessCategoryName),
		HasManagement:         hasManagement,
		CategoryName:          ns(pr.Data.User.CategoryName),
		Biography:             ns(pr.Data.User.Biography),
		Email:                 ns(pe),
		EmailProvider:         ns(pv),
	}
}
