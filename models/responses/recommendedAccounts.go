package responses

type RecommendedAccountsResponse struct {
	IsRecommendAccount bool   `json:"is_recommend_account"`
	Status             string `json:"status"`
	Users              []*struct {
		Pk                 int64  `json:"pk"`
		Username           string `json:"username"`
		FullName           string `json:"full_name"`
		IsPrivate          bool   `json:"is_private"`
		ProfilePicURL      string `json:"profile_pic_url"`
		ProfilePicID       string `json:"profile_pic_id"`
		IsVerified         bool   `json:"is_verified"`
		FollowFrictionType int    `json:"follow_friction_type"`
		GrowthFrictionInfo struct {
			HasActiveInterventions bool `json:"has_active_interventions"`
			Interventions          struct {
			} `json:"interventions"`
		} `json:"growth_friction_info"`
		ChainingInfo struct {
			Sources   string `json:"sources"`
			Algorithm string `json:"algorithm"`
		} `json:"chaining_info"`
		ProfileChainingSecondaryLabel string `json:"profile_chaining_secondary_label"`
		SocialContext                 string `json:"social_context"`
		FriendshipStatus              struct {
			Following       bool `json:"following"`
			FollowedBy      bool `json:"followed_by"`
			Blocking        bool `json:"blocking"`
			Muting          bool `json:"muting"`
			IsPrivate       bool `json:"is_private"`
			IncomingRequest bool `json:"incoming_request"`
			OutgoingRequest bool `json:"outgoing_request"`
			IsBestie        bool `json:"is_bestie"`
			IsRestricted    bool `json:"is_restricted"`
			IsFeedFavorite  bool `json:"is_feed_favorite"`
		} `json:"friendship_status"`
	} `json:"users"`
}
