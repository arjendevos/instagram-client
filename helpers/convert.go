package helpers

import (
	"github.com/arjendevos/instagram-client/models"
	"github.com/arjendevos/instagram-client/models/responses"
)

func ConvertProfile(pr *responses.Profile) models.User {
	emails := ExtractEmails(pr.Data.User.Biography)
	pe, pv := FindPrimaryEmail(emails)

	return models.User{
		ID:                    pr.Data.User.ID,
		IsBusinessAccount:     pr.Data.User.IsBusinessAccount,
		IsProfessionalAccount: pr.Data.User.IsProfessionalAccount,
		IsVerfied:             pr.Data.User.IsVerified,
		Username:              pr.Data.User.Username,
		FullName:              pr.Data.User.FullName,
		ProfilePicture:        pr.Data.User.ProfilePicURLHd,
		Count: models.Count{
			Post:      pr.Data.User.EdgeOwnerToTimelineMedia.Count,
			Followers: pr.Data.User.EdgeFollowedBy.Count,
			Followed:  pr.Data.User.EdgeFollow.Count,
		},
		ExternalUrl: pr.Data.User.ExternalURL,
		Business: models.Business{
			Email:    pr.Data.User.BusinessEmail,
			Phone:    pr.Data.User.BusinessPhoneNumber,
			Category: pr.Data.User.BusinessCategoryName,
		},
		Category:      pr.Data.User.CategoryName,
		Biography:     pr.Data.User.Biography,
		Email:         pe,
		EmailProvider: pv,
	}
}
